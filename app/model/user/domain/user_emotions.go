package domain

import (
	"book_store/app/model/helper"
	"time"
)

type UserEmotion struct {
	ID             int       `json:"id"`
	Emotion        string    `json:"emotion"`
	Session        string    `json:"session"`
	Temperature    string    `json:"temperature"`
	UserQuestionId int       `json:"user_question_id"`
	Time           time.Time `json:"time"`
	Date           time.Time `json:"date"`
	UserId         User      `json:"user_id"`
}

type UserEmotionResponse struct {
	ID             int       `json:"id"`
	Emotion        string    `json:"emotion"`
	Session        string    `json:"session"`
	Temperature    string    `json:"temperature"`
	UserQuestionId int       `json:"user_question_id"`
	Time           time.Time `json:"time"`
	Date           time.Time `json:"date"`
	UserId         User      `json:"user_id"`
}

func (data UserEmotion) ToResponse() UserEmotionResponse {
	return UserEmotionResponse{
		ID:             data.ID,
		Emotion:        data.Emotion,
		Session:        data.Session,
		Temperature:    data.Temperature,
		UserQuestionId: data.UserQuestionId,
		Time:           data.Time,
		Date:           data.Date,
		UserId:         data.UserId,
	}
}

func NewUserEmotionResponse(code int, status string, user UserEmotion) helper.ReturnResponse {
	userEmotionResponse := user.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userEmotionResponse,
	}
}
