package user

import (
	"book_store/app/model/helper"
	"book_store/app/model/user/domain"
	"book_store/app/model/user/request"
	"book_store/app/repository/auth"
	"book_store/config"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Register(ctx context.Context, register request.Register) helper.ReturnResponse
	SendVerifyCode(ctx context.Context, register request.SendVerifyCode, token string) helper.ReturnResponse
	VerifyEmail(ctx context.Context, register request.VerifyEmail, token string) helper.ReturnResponse
	AddPassword(ctx context.Context, register request.AddPassword, token string) helper.ReturnResponse
	AddInformation(ctx context.Context, register request.AddInformation, token string) helper.ReturnResponse
	AddOptionalInformation(ctx context.Context, register request.AddOptionalInformation, token string) helper.ReturnResponse
}

// UserServiceImpl struct
type AuthServiceImpl struct {
	AuthRepository auth.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(db *sql.DB, validate *validator.Validate, authRepository auth.AuthRepository) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) performTransaction(ctx context.Context, task func(tx *sql.Tx) (helper.ReturnResponse, error)) helper.ReturnResponse {
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

// Register registers a new user
func (service *AuthServiceImpl) Register(ctx context.Context, request request.Register) helper.ReturnResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		if tx == nil {
			helper2.ErrorServiceResponse(500, "ERROR", "Internal Server Error", "Transaction is nil", err)
		}

		userReq := domain.User{
			Email: request.Email,
		}

		userRes, resJson, err := service.AuthRepository.Register(ctx, tx, userReq)
		if err != nil {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}

		if resJson.Status == "ERROR" {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}

		return helper.ReturnResponse{
			Code:    resJson.Code,
			Status:  resJson.Status,
			Message: resJson.Message,
			Data:    userRes,
		}, nil
	})
}

func (service *AuthServiceImpl) SendVerifyCode(ctx context.Context, request request.SendVerifyCode, token string) helper.ReturnResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		return helper.ReturnResponse{
			Code:    400,
			Status:  "ERROR",
			Message: "Bad Request",
			Data:    err.Error(),
		}
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		if tx == nil {
			return helper.ReturnResponse{
				Code:    500,
				Status:  "ERROR",
				Message: "Internal Server Error",
				Data:    "Transaction is nil",
			}, err
		}

		userReq := domain.User{
			Token: request.Token,
		}
		userRes, resJson, err := service.AuthRepository.SendVerifyCode(ctx, tx, userReq, token)
		if err != nil {
			return helper.ReturnResponse{
				Code:    resJson.Code,
				Status:  resJson.Status,
				Message: resJson.Message,
				Data:    resJson.Error,
			}, err
		}

		if resJson.Status == "ERROR" {
			return helper.ReturnResponse{
				Code:    resJson.Code,
				Status:  resJson.Status,
				Message: resJson.Message,
				Data:    resJson.Error,
			}, nil
		}

		return helper.ReturnResponse{
			Code:    resJson.Code,
			Status:  resJson.Status,
			Message: resJson.Message,
			Data:    userRes,
		}, nil

	})
}

func (service *AuthServiceImpl) VerifyEmail(ctx context.Context, request request.VerifyEmail, token string) helper.ReturnResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		if tx == nil {
			_, _ = helper2.ErrorServiceResponse(500, "ERROR", "Internal Server Error", "Transaction is nil", err)
		}

		userReq := domain.User{
			Token:      request.Token,
			VerifyCode: request.VerifyCode,
		}
		userRes, resJson, err := service.AuthRepository.VerifyEmail(ctx, tx, userReq, token)
		if err != nil {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}
		if resJson.Status == "ERROR" {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}

		return helper.ReturnResponse{
			Code:    resJson.Code,
			Status:  resJson.Status,
			Message: resJson.Message,
			Data:    userRes,
		}, nil

	})
}

func (service *AuthServiceImpl) AddPassword(ctx context.Context, request request.AddPassword, token string) helper.ReturnResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	// Validate password match
	if request.Password != request.ConfirmPassword {
		_, _ = helper2.ErrorServiceResponse(401, "ERROR", "Request Error", "Password doesn't match", err)
	}

	// Perform transaction
	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		// Prepare user data
		userReq := domain.User{
			Token:    request.Token,
			Password: request.Password,
		}

		// Call repository to add password
		userRes, resJson, err := service.AuthRepository.AddPassword(ctx, tx, userReq, token)
		if err != nil {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}
		if resJson.Status == "ERROR" {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}

		// Return success response
		return helper.ReturnResponse{
			Code:    resJson.Code,
			Status:  resJson.Status,
			Message: resJson.Message,
			Data:    userRes,
		}, nil
	})
}

func (service *AuthServiceImpl) AddInformation(ctx context.Context, request request.AddInformation, token string) helper.ReturnResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		userReq := domain.User{
			Token:       request.Token,
			Firstname:   request.Firstname,
			Lastname:    request.Lastname,
			Username:    request.Username,
			PhoneNumber: request.PhoneNumber,
		}

		userRes, resJson, err := service.AuthRepository.AddInformation(ctx, tx, userReq, token)
		if err != nil {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}
		if resJson.Status == "ERROR" {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}

		return helper.ReturnResponse{
			Code:    resJson.Code,
			Status:  resJson.Status,
			Message: resJson.Message,
			Data:    userRes,
		}, nil
	})
}

func (service *AuthServiceImpl) AddOptionalInformation(ctx context.Context, request request.AddOptionalInformation, token string) helper.ReturnResponse {
	// Validate the request structure
	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		userReq := domain.User{
			Token:    request.Token,
			Bio:      request.Bio,
			Birthday: request.Birthday,
			Theme:    request.Theme,
		}

		userRes, resJson, err := service.AuthRepository.AddOptionalInformation(ctx, tx, userReq, token, request.FileHeader)
		if err != nil {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}
		if resJson.Status == "ERROR" {
			_, _ = helper2.ErrorServiceResponse(resJson.Code, resJson.Status, resJson.Message, resJson.Error, err)
		}

		return helper.ReturnResponse{
			Code:    resJson.Code,
			Status:  resJson.Status,
			Message: resJson.Message,
			Data:    userRes,
		}, nil
	})
}
