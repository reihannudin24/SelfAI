package auth

import (
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

func (repository *AuthRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, helper2.ResponseJson, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	exists, err := repository.helperRepository.ValidateDontExists("users", "email", user.Email)
	if err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}
	if !exists {
		_, _, _ = helper2.ErrorRequest(user, "Email telah terdaftar", nil)
	}

	bcryptToken, err := bcrypt.GenerateFromPassword([]byte(user.Token), bcrypt.DefaultCost)
	if err != nil {
		return user, helper2.ResponseJson{"ERROR", 500, "Internal Server Error", "Gagal membuat token"}, err
	}

	SQL := "INSERT INTO users(email, level, point, type, token) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, user.Email, 0, 0, "basic", bcryptToken)
	if err != nil {
		_, _, _ = helper2.ErrorRequest(user, "Gagal mendaftarkan pengguna\"", nil)
	}

	return user, helper2.ResponseJson{
		"OK",
		201,
		"Berhasi membuat akun",
		nil}, nil
}

func (repository *AuthRepositoryImpl) SendVerifyCode(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}
	if !exists {
		_, _, _ = helper2.ErrorRequest(user, "User not found", nil)
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(900000) + 100000

	SQL := "UPDATE users SET verify_code = ? WHERE token = ?"
	_, err = tx.ExecContext(ctx, SQL, randomNumber, token)
	if err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	err = sendVerificationCodeEmail(foundUser.Email, randomNumber)
	if err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "Verification code sent successfully",
		Error:   "",
	}, nil
}

func (repository *AuthRepositoryImpl) VerifyEmail(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error) {
	var err error
	if tx == nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	exists, foundUser, err := repository.helperRepository.ValidateUserExists("users", "token", token)
	if err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}
	if !exists {
		_, _, _ = helper2.ErrorRequest(user, "User not found", nil)
	}
	if user.VerifyCode != foundUser.VerifyCode {
		_, _, _ = helper2.ErrorRequest(user, "Verification code does not match", nil)
	}

	if foundUser.EmailVerify == true {
		_, _, _ = helper2.ErrorRequest(user, "User already verify", nil)
	}

	updateSQL := "UPDATE users SET email_verify = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, updateSQL, true, foundUser.ID)
	if err != nil {
		_, _, _ = helper2.ErrorRequest(user, "Failed to update email verification status", nil)
	}

	if err = tx.Commit(); err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
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
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}
	if !exists {
		_, _, _ = helper2.ErrorRequest(user, "User not found", nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		_, _, _ = helper2.ErrorRequest(user, "Failed to generate password hash", nil)
	}

	SQL := "UPDATE users SET password = ? WHERE email = ?"
	_, err = tx.ExecContext(ctx, SQL, hashedPassword, foundUser.Email)
	if err != nil {
		_, _, _ = helper2.ErrorRequest(user, "Failed to update password", nil)
	}
	if err = tx.Commit(); err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "User password added successfully",
		Error:   "",
	}, nil
}

func (repository *AuthRepositoryImpl) AddInformation(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}
	if !exists {
		_, _, _ = helper2.ErrorRequest(user, "User not found", nil)
	}

	SQL := "UPDATE users SET username = ?, firstname = ?, lastname = ?, phone_number = ? WHERE email = ?"
	_, err = tx.ExecContext(ctx, SQL, user.Username, user.Firstname, user.Lastname, user.PhoneNumber, foundUser.Email)
	if err != nil {
		_, _, _ = helper2.ErrorRequest(user, "Failed to update user information", nil)
	}

	if err = tx.Commit(); err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "User information successfully updated",
		Error:   "",
	}, nil
}

func (repository *AuthRepositoryImpl) AddOptionalInformation(ctx context.Context, tx *sql.Tx, user domain.User, token string, fileHeader *multipart.FileHeader) (domain.User, helper2.ResponseJson, error) {
	if tx == nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}
	if !exists {
		_, _, _ = helper2.ErrorRequest(user, "User not found", nil)
	}

	var imagePath string
	if fileHeader != nil {
		imagePath, err = repository.UploadRepo.UploadFile(fileHeader)
		if err != nil {
			_, _, _ = helper2.ErrorRequest(user, "Failed to upload file", nil)
		}
	}

	SQL := "UPDATE users SET bio = ?, image_path = ?, birthday = ?,  theme = ? WHERE email = ?"
	_, err = tx.ExecContext(ctx, SQL, user.Bio, imagePath, user.Birthday, user.Theme, foundUser.Email)
	if err != nil {
		_, _, _ = helper2.ErrorRequest(user, "Failed to update email verification status", nil)
	}

	if err = tx.Commit(); err != nil {
		_, _, _ = helper2.ErrorInter(user, nil)
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully updated information",
		Error:   "",
	}, nil
}
