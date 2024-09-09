package file

import (
	"book_store/app/model/user/domain"
	"book_store/app/repository/helper"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

type FileRepositoryImpl struct {
	DB               *sql.DB
	helperRepository *helper.Repository
}

func NewUserRepository(db *sql.DB) FileRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &FileRepositoryImpl{
		DB:               db,
		helperRepository: &helper.Repository{DB: db},
	}
}

func (repository *FileRepositoryImpl) UploadFile(ctx context.Context, tx *sql.Tx, file domain.File, fileUser *multipart.FileHeader, token string) (domain.File, helper2.ResponseJson, error) {
	if tx == nil {
		return file, helper2.ResponseJson{"ERROR", 500, "Internal Server Error", "Transaction is nil", nil}, fmt.Errorf("transaction is nil")
	}

	// Validate if the user exists using the provided token
	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return file, helper2.ResponseJson{"ERROR", 500, "Internal Server Error", err.Error(), nil}, err
	}
	if !exists {
		return file, helper2.ResponseJson{"ERROR", 409, "Conflict", "Email has not been registered", nil}, fmt.Errorf("email has not registered")
	}

	// open the uploaded file
	src, err := fileUser.Open()
	if err != nil {
		return file, helper2.ResponseJson{"ERROR", 400, "Bad Request", "Failed to open file", nil}, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	destPath := filepath.Join("public/assets/files/uploaded", file.Files)
	if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
		return file, helper2.ResponseJson{"ERROR", 500, "Internal Server Error", "Failed to create directory", nil}, fmt.Errorf("failed to create directory: %w", err)
	}

	// create the destination file
	dst, err := os.Create(destPath)
	if err != nil {
		return file, helper2.ResponseJson{"ERROR", 500, "Internal Server Error", "Failed to create file", nil}, fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// copy the uploaded file to the destination
	if _, err := io.Copy(dst, src); err != nil {
		return file, helper2.ResponseJson{"ERROR", 500, "Internal Server Error", "Failed to copy file", nil}, fmt.Errorf("failed to copy file: %w", err)
	}

	// Set additional fields in the domain.File struct
	file.FilesPath = "/public/" + file.Files
	file.UserId = strconv.Itoa(foundUser.ID)

	query := `INSERT INTO files (name, slug, typ, files, files_path, user_id) VALUES (? ,?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, query, file.Name, file.Slug, file.Type, fileUser.Filename, file.FilesPath, file.UserId)
	if err != nil {
		return file, helper2.ResponseJson{"ERROR", 500, "Internal Server Error", "Failed to save file metadata", nil}, fmt.Errorf("failed to save file metadata: %w", err)
	}

	// Return success response
	return file, helper2.ResponseJson{"OK", 201, "File uploaded successfully", nil, nil}, nil

}
