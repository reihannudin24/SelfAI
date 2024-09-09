package domain

import (
	"book_store/app/model/helper"
	"book_store/app/model/user/domain"
	"time"
)

type Activities struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	Address   string    `json:"address"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Date      time.Time `json:"date"`
	Remainder string    `json:"remainder"`
	Type      string    `json:"type"`
	UserId    int       `json:"user_id"`
}

type ActivitiesResponse struct {
	ID        int         `json:"id"`
	Content   string      `json:"content"`
	Address   string      `json:"address"`
	StartTime time.Time   `json:"start_time"`
	EndTime   time.Time   `json:"end_time"`
	Date      time.Time   `json:"date"`
	Remainder string      `json:"remainder"`
	Type      string      `json:"type"`
	User      domain.User `json:"user"`
	UserId    int         `json:"user_id"`
}

func (data Activities) ToResponse() ActivitiesResponse {
	return ActivitiesResponse{
		ID:        data.ID,
		Content:   data.Content,
		Address:   data.Address,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		Date:      data.Date,
		Remainder: data.Remainder,
		Type:      data.Type,
		User:      domain.User{},
		UserId:    data.UserId,
	}
}

func NewActivitiesResponse(code int, status string, activities Activities) helper.ReturnResponse {
	dataResponse := activities.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   dataResponse,
	}

}
