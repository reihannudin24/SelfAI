package auth

import (
	"book_store/config"
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"mime/multipart"
	"net/smtp"
	"time"

	"book_store/app/model/user/domain"
	"book_store/app/repository/helper"
	helper2 "book_store/helper"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct {
	DB               *sql.DB
	helperRepository *helper.Repository
	UploadRepo       *helper.UploadRepository
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &AuthRepositoryImpl{
		DB: db,
		helperRepository: &helper.Repository{
			DB: db,
		},
		UploadRepo: &helper.UploadRepository{
			DB: db,
		},
	}
}

func sendVerificationCodeEmail(to string, code int) error {
	from := "noreply@book_apps.com"
	password := "080421bdd277ed"
	smtpHost := "sandbox.smtp.mailtrap.io"
	smtpPort := "587"

	message := fmt.Sprintf("Subject: Verification Code\nContent-Type: text/plain; charset=UTF-8\n\nYour verification code is: %d", code)
	auth := smtp.PlainAuth("", "fcfc71c94472cb", password, smtpHost)

	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, from, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func (repository *AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, err := repository.helperRepository.ValidateDontExists("users", "email", user.Email)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}

	if !exists {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    401,
			Message: "Request Error",
			Error:   "Email telah terdaftar",
		}, fmt.Errorf("email telah terdaftar")
	}

	bcryptToken, err := bcrypt.GenerateFromPassword([]byte(user.Token), bcrypt.DefaultCost)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   "Failed to generate token",
		}, err
	}

	SQL := "INSERT INTO users(email, level, point, type, token) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, user.Email, 0, 0, "basic", bcryptToken)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: fmt.Sprintf("Gagal mendaftarkan pengguna: %s", err.Error()),
			Error:   err.Error(),
		}, err
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "Berhasi membuat akun",
		Error:   "",
	}, nil
}

func (repository *AuthRepositoryImpl) SendVerifyCode(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}

	if !exists {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(900000) + 100000

	SQL := "UPDATE users SET verify_code = ? WHERE token = ?"
	_, err = tx.ExecContext(ctx, SQL, randomNumber, token)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to update verification code",
			Error:   err.Error(),
		}, err
	}

	err = sendVerificationCodeEmail(foundUser.Email, randomNumber)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to send verification code email",
			Error:   err.Error(),
		}, err
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "Verification code sent successfully",
		Error:   "",
	}, nil
}

func (repository *AuthRepositoryImpl) VerifyEmail(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateUserExists("users", "token", token)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}

	if !exists {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	if user.VerifyCode != foundUser.VerifyCode {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    401,
			Message: "Verification code does not match",
			Error:   "Verification code does not match",
		}, fmt.Errorf("verification code does not match")
	}

	if foundUser.EmailVerify {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    400,
			Message: "User already verified",
			Error:   "User already verified",
		}, fmt.Errorf("user already verified")
	}

	updateSQL := "UPDATE users SET email_verify = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, updateSQL, 1, foundUser.ID)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to update email verification status",
			Error:   err.Error(),
		}, err
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "User verified successfully",
		Error:   "",
	}, nil
}

func (repository *AuthRepositoryImpl) AddPassword(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error) {

	if tx == nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}

	if !exists {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	if foundUser.Password != "" {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    400,
			Message: "User already has a password",
			Error:   "User already has a password",
		}, fmt.Errorf("user already has a password")
	}

	cost := 12
	generatedHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to generate password",
			Error:   err.Error(),
		}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(generatedHash), []byte(user.Password))
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    400,
			Message: "Password does not match",
			Error:   "Provided password does not match the hash",
		}, fmt.Errorf("password does not match")
	}

	fmt.Printf("user.Password: %s\n", user.Password)
	fmt.Printf("generatedHash: %s\n", generatedHash)

	// Update the password in the database
	updateSQL := "UPDATE users SET password = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, updateSQL, generatedHash, foundUser.ID)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to update password",
			Error:   err.Error(),
		}, err
	}

	// Return success response
	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "Password set successfully",
		Error:   "",
	}, nil
}

func (repository *AuthRepositoryImpl) AddInformation(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}

	if !exists {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	updateSQL := "UPDATE users SET username = ?, firstname = ?, lastname = ?, phone_number = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, updateSQL, user.Username, user.Firstname, user.Lastname, user.PhoneNumber, foundUser.ID)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to update user information",
			Error:   err.Error(),
		}, err
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "User information updated successfully",
		Error:   "",
	}, nil
}

func (repository *AuthRepositoryImpl) AddOptionalInformation(ctx context.Context, tx *sql.Tx, user domain.User, token string, fileHeader *multipart.FileHeader) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}

	print("exists : ", exists)

	if !exists {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	path, err := repository.UploadRepo.UploadFile(fileHeader)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to upload file",
			Error:   err.Error(),
		}, err
	}

	updateSQL := "UPDATE users SET bio = ?, birthday = ?, theme = ?, image_path = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, updateSQL, user.Bio, user.Birthday, user.Theme, path, foundUser.ID)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to update optional user information",
			Error:   err.Error(),
		}, err
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Optional information updated successfully",
		Error:   "",
	}, nil
}

func (repository *AuthRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExistsGetPassword("users", "email", user.Email)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}

	if !exists {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "Invalid email or password",
		}, fmt.Errorf("user not found")
	}

	print("Password db : %s", string(foundUser.Password))

	// Ensure foundUser.Password is the hashed password and user.Password is the plain text password
	fmt.Printf("Stored Hash: %s\n", foundUser.Password)

	// Generate a hash from the user password (for testing purposes, this should be removed in production)
	generatedHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to generate password",
			Error:   err.Error(),
		}, err
	}
	fmt.Println("Generated Hash (for reference):", string(generatedHash))

	// Compare the stored hash with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	fmt.Println("Err:", err)
	if err != nil {
		fmt.Println("Password does not match:", err)
	} else {
		fmt.Println("Password matches!")
	}

	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    401,
			Message: "Authentication failed",
			Error:   "Invalid email or password",
		}, fmt.Errorf("invalid email or password")
	}

	// Generate JWT token
	token, err := config.GenerateToken(user.Username)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to generate token",
			Error:   err.Error(),
		}, err
	}

	updateSQL := "UPDATE users SET token = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, updateSQL, token, foundUser.ID)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to store token",
			Error:   err.Error(),
		}, err
	}

	return foundUser, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Login successful",
		Error:   "",
		Data: map[string]interface{}{
			"user":  foundUser,
			"token": token,
		},
	}, nil
}

func (repository *AuthRepositoryImpl) Logout(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, helper2.ResponseJson, error) {

	if tx == nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	if user.Token == "" {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "User error",
			Error:   "User don't have any session",
		}, fmt.Errorf("User don't have any session")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", user.Token)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}

	if !exists {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "UPDATE users SET token = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, nil, foundUser.ID)
	if err != nil {
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to update verification code",
			Error:   err.Error(),
		}, err
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "Successfully logout",
		Error:   "",
	}, nil

}
