package user

import (
	"book_store/app/model/helper"
	"book_store/app/model/user/domain"
	"book_store/app/model/user/request"
	"book_store/app/repository/user"
	"book_store/config"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	Profile(ctx context.Context, request request.AccSessionAuth, slug string) helper.ReturnResponse
}

type UserServiceImpl struct {
	UserRepository user.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(db *sql.DB, validate *validator.Validate, userRepository user.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) performTransaction(ctx context.Context, task func(tx *sql.Tx) (helper.ReturnResponse, error)) helper.ReturnResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		helper2.ErrorServiceInternal(err)
	}
	defer config.CommitOrRollback(tx)

	response, err := task(tx)
	if err != nil {
		helper2.ErrorServiceInternal(err)
	}
	return response
}

func (service *UserServiceImpl) Profile(ctx context.Context, request request.AccSessionAuth, slug string) helper.ReturnResponse {

	tx, err := service.DB.Begin()
	userResponse := domain.User{
		Token: request.Token,
	}

	userRes, resJson, err := service.UserRepository.Profile(ctx, tx, userResponse, slug)
	if resJson.Status == "ERROR" {
		_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
	}

	return helper.ReturnResponse{
		Code:    resJson.Code,
		Status:  resJson.Status,
		Message: resJson.Message,
		Error:   resJson.Error,
		Data:    userRes,
	}

}
