package domain

import (
	"book_store/app/model/helper"
	"time"
)

type HowDays struct {
	ID            int       `json:"id"`
	UserSentiment string    `json:"user_sentiment"`
	Sentiment     string    `json:"sentiment"`
	Temperature   string    `json:"temperature"`
	Date          time.Time `json:"date"`
	UserId        int       `json:"user_id"`
	NewGroupId    int       `json:"new_group_id"`
}

type HowDaysResponse struct {
	ID            int       `json:"id"`
	UserSentiment string    `json:"user_sentiment"`
	Sentiment     string    `json:"sentiment"`
	Temperature   string    `json:"temperature"`
	Date          time.Time `json:"date"`
	UserId        int       `json:"user_id"`
	NewGroupId    int       `json:"new_group_id"`
}

func (data HowDays) ToResponse() HowDaysResponse {
	return HowDaysResponse{
		ID:            data.ID,
		UserSentiment: data.UserSentiment,
		Sentiment:     data.Sentiment,
		Temperature:   data.Temperature,
		Date:          data.Date,
		UserId:        data.UserId,
		NewGroupId:    data.NewGroupId,
	}
}

func NewHowDaysResponse(code int, status string, response HowDays) helper.ReturnResponse {
	dataResponse := response.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   dataResponse,
	}
}
