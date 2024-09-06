package file

import (
	"book_store/app/model/helper"
	"book_store/app/model/user/domain"
	"book_store/app/model/user/request"
	"book_store/app/repository/file"
	"book_store/config"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type FileService interface {
	UploadFile(ctx context.Context, request request.UploadFile, token string) helper.ReturnResponse
}

type FileServiceImpl struct {
	FileRepository file.FileRepository
	DB             *sql.DB
	validate       *validator.Validate
}

func NewFileService(db *sql.DB, validate *validator.Validate, fileRepository file.FileRepository) FileService {
	return &FileServiceImpl{
		FileRepository: fileRepository,
		DB:             db,
		validate:       validate,
	}
}

func (service *FileServiceImpl) performTransaction(ctx context.Context, task func(tx *sql.Tx) (helper.ReturnResponse, error)) helper.ReturnResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		return helper2.ErrorService(err, 500, "Internal Server Error", "Failed to start transaction")
	}
	defer config.CommitOrRollback(tx)

	response, err := task(tx)
	if err != nil {
		return helper2.ErrorService(err, 500, "Internal Server Error", err.Error())
	}

	return response
}

func (service *FileServiceImpl) UploadFile(ctx context.Context, request request.UploadFile, token string) helper.ReturnResponse {

	err := service.validate.Struct(request)
	if err != nil {
		return helper2.ErrorService(err, 400, "Bad Request", err.Error())
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		file := domain.File{
			Name: request.FileName,
		}

		fileRes, _, err := service.FileRepository.UploadFile(ctx, tx, file, request.FileHeader, token)
		if err != nil {
			return helper.ReturnResponse{}, err
		}

		return helper.ReturnResponse{
			Code:    201,
			Message: "File uploaded successfully",
			Data:    fileRes,
		}, nil
	})
}
