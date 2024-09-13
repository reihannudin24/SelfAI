package activities

import (
	domain2 "book_store/app/model/activities/domain"
	"book_store/app/model/activities/request"
	"book_store/app/model/helper"
	request2 "book_store/app/model/user/request"
	"book_store/app/repository/activities"
	"book_store/config"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type ActivitiesService interface {
	Create(ctx context.Context, request request.CreateActivities, token string) helper.ReturnResponse
	Update(ctx context.Context, request request.UpdateActivities, token string) helper.ReturnResponse
	Delete(ctx context.Context, request request.DeleteActivities, token string) helper.ReturnResponse
	Show(ctx context.Context, request request2.AccSessionAuth, slug string, token string) helper.ReturnResponse
	All(ctx context.Context, request request2.AccSessionAuth, token string) helper.ReturnResponse
}

type ActivitiesServiceImpl struct {
	ActivitiesRepository activities.ActivitiesRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewActivitiesService(db *sql.DB, validate *validator.Validate, activitiesRepository activities.ActivitiesRepository) ActivitiesService {
	return &ActivitiesServiceImpl{
		ActivitiesRepository: activitiesRepository,
		DB:                   db,
		Validate:             validate,
	}
}

func (service *ActivitiesServiceImpl) performTransaction(ctx context.Context, task func(tx *sql.Tx) (helper.ReturnResponse, error)) helper.ReturnResponse {
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

func (service *ActivitiesServiceImpl) Create(ctx context.Context, request request.CreateActivities, token string) helper.ReturnResponse {

	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		if tx == nil {
			helper2.ErrorServiceResponse(500, "ERROR", "Internal Server Error", "Transaction is nil", err)
		}

		slug := helper2.GenerateSlug(request.Title)

		userReq := domain2.Activities{
			Title:     request.Title,
			Slug:      slug,
			Content:   request.Content,
			Address:   request.Address,
			StartTime: request.StartTime,
			EndTime:   request.EndTime,
			Date:      request.Date,
			Remainder: request.Remainder,
			Type:      request.Type,
		}

		userRes, resJson, err := service.ActivitiesRepository.Create(ctx, tx, userReq, request.Token)
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

func (service *ActivitiesServiceImpl) Update(ctx context.Context, request request.UpdateActivities, token string) helper.ReturnResponse {

	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		if tx == nil {
			helper2.ErrorServiceResponse(500, "ERROR", "Internal Server Error", "Transaction is nil", err)
		}

		slug := helper2.GenerateSlug(request.Title)

		userReq := domain2.Activities{
			ID:        request.ID,
			Title:     request.Title,
			Slug:      slug,
			Content:   request.Content,
			Address:   request.Address,
			StartTime: request.StartTime,
			EndTime:   request.EndTime,
			Date:      request.Date,
			Remainder: request.Remainder,
			Type:      request.Type,
		}

		userRes, resJson, err := service.ActivitiesRepository.Update(ctx, tx, userReq, request.Token)
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

func (service *ActivitiesServiceImpl) Delete(ctx context.Context, request request.DeleteActivities, token string) helper.ReturnResponse {

	err := service.Validate.Struct(request)
	if err != nil {
		helper2.ErrorServiceRequest(err)
	}

	return service.performTransaction(ctx, func(tx *sql.Tx) (helper.ReturnResponse, error) {
		if tx == nil {
			helper2.ErrorServiceResponse(500, "ERROR", "Internal Server Error", "Transaction is nil", err)
		}

		userReq := domain2.Activities{
			ID: request.Id,
		}

		userRes, resJson, err := service.ActivitiesRepository.Delete(ctx, tx, userReq, request.Token)
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

func (service *ActivitiesServiceImpl) Show(ctx context.Context, request request2.AccSessionAuth, slug string, token string) helper.ReturnResponse {

	tx, err := service.DB.Begin()
	userResponse := domain2.Activities{
		Token: request.Token,
	}

	userRes, resJson, err := service.ActivitiesRepository.Show(ctx, tx, userResponse, request.Token, slug)
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

func (service *ActivitiesServiceImpl) All(ctx context.Context, request request2.AccSessionAuth, token string) helper.ReturnResponse {

	tx, err := service.DB.Begin()
	userResponse := domain2.Activities{
		Token: request.Token,
	}

	userRes, resJson, err := service.ActivitiesRepository.All(ctx, tx, userResponse, request.Token)
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
