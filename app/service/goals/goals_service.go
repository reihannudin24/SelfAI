package goals

import (
	"book_store/app/model/goals/domain"
	"book_store/app/model/goals/request"
	"book_store/app/model/helper"
	request2 "book_store/app/model/user/request"
	"book_store/app/repository/goals"
	"book_store/config"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type GoalsService interface {
	Add(ctx context.Context, request request.AddGoals, token string) helper.ReturnResponse
	Update(ctx context.Context, request request.UpdateGoals, token string) helper.ReturnResponse
	Delete(ctx context.Context, request request.DeleteGoals, token string) helper.ReturnResponse
	Show(ctx context.Context, request request2.AccSessionAuth, slug string, token string) helper.ReturnResponse
	All(ctx context.Context, request request2.AccSessionAuth, token string) helper.ReturnResponse
}

type GoalServiceImpl struct {
	GoalRepository goals.GoalRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewGoalService(db *sql.DB, validate *validator.Validate, goalRepository goals.GoalRepository) GoalsService {
	return &GoalServiceImpl{
		GoalRepository: goalRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *GoalServiceImpl) performTransaction(ctx context.Context, task func(tx *sql.Tx) (helper.ReturnResponse, error)) helper.ReturnResponse {
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

func (service *GoalServiceImpl) Add(ctx context.Context, request request.AddGoals, token string) helper.ReturnResponse {

	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		if tx == nil {
			helper2.ErrorServiceResponse(500, "ERROR", "Internal Server Error", "Transaction is nil", err)
		}

		userReq := domain.Goals{
			Goal:       request.Goal,
			Status:     request.Goal,
			Type:       request.Type,
			Sentiment:  request.Goal,
			Time:       request.Time,
			Date:       request.Date,
			NewGroupId: request.NewGroupId,
			UserId:     request.UserId,
		}

		userRes, resJson, err := service.GoalRepository.Create(ctx, tx, userReq, token)
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
			Error:   resJson.Error,
			Data:    userRes,
		}, nil
	})

}

func (service *GoalServiceImpl) Update(ctx context.Context, request request.UpdateGoals, token string) helper.ReturnResponse {

	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		if tx == nil {
			helper2.ErrorServiceResponse(500, "ERROR", "Internal Server Error", "Transaction is nil", err)
		}

		userReq := domain.Goals{
			ID:         request.ID,
			Goal:       request.Goal,
			Status:     request.Goal,
			Type:       request.Type,
			Sentiment:  request.Goal,
			Time:       request.Time,
			Date:       request.Date,
			NewGroupId: request.NewGroupId,
			UserId:     request.UserId,
		}

		userRes, resJson, err := service.GoalRepository.Update(ctx, tx, userReq, token)
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
			Error:   resJson.Error,
			Data:    userRes,
		}, nil
	})

}

func (service *GoalServiceImpl) Delete(ctx context.Context, request request.DeleteGoals, token string) helper.ReturnResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		if tx == nil {
			helper2.ErrorServiceResponse(500, "ERROR", "Internal Server Error", "Transaction is nil", err)
		}

		userReq := domain.Goals{
			ID: request.ID,
		}

		userRes, resJson, err := service.GoalRepository.Delete(ctx, tx, userReq, token)
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
			Error:   resJson.Error,
			Data:    userRes,
		}, nil
	})
}

func (service *GoalServiceImpl) Show(ctx context.Context, request request2.AccSessionAuth, slug string, token string) helper.ReturnResponse {
	tx, err := service.DB.Begin()
	userResponse := domain.Goals{
		Token: request.Token,
	}

	userRes, resJson, err := service.GoalRepository.Show(ctx, tx, userResponse, request.Token)
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

func (service *GoalServiceImpl) All(ctx context.Context, request request2.AccSessionAuth, token string) helper.ReturnResponse {
	tx, err := service.DB.Begin()
	userResponse := domain.Goals{
		Token: request.Token,
	}

	userRes, resJson, err := service.GoalRepository.All(ctx, tx, userResponse, request.Token)
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
