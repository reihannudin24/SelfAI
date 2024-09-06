package helper

import (
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type UploadRepository struct {
	DB *sql.DB
}

// UploadFile saves the uploaded file to the public directory and returns the file path.
func (repo *UploadRepository) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Generate a unique file name
	ext := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	imagePath := filepath.Join("public/assets/img/profil", fileName) // Adjust the directory as needed

	// Create the directory if it doesn't exist
	err = os.MkdirAll(filepath.Dir(imagePath), os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// Save the file
	out, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return imagePath, nil
}
