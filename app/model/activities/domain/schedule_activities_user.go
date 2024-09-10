package domain

import "book_store/app/model/helper"

type ScheduleActivitiesUser struct {
	ID       int    `json:"id"`
	Schedule string `json:"schedule"`
	Message  string `json:"message"`
	UserId   int    `json:"user_id"`
}

type ScheduleActivitiesUserResponse struct {
	ID       int    `json:"id"`
	Schedule string `json:"schedule"`
	Message  string `json:"message"`
	UserId   int    `json:"user_id"`
}

func (data ScheduleActivitiesUser) ToResponse() ScheduleActivitiesUserResponse {
	return ScheduleActivitiesUserResponse{
		ID:       data.ID,
		Schedule: data.Schedule,
		Message:  data.Message,
		UserId:   data.UserId,
	}
}

func NewScheduleActivitiesUserResponse(code int, status string, user ScheduleActivitiesUser) helper.ReturnResponse {
	userResponse := user.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userResponse,
	}
}
