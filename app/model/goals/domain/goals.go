package domain

import (
	"book_store/app/model/helper"
	"time"
)

type Goals struct {
	ID         int       `json:"id"`
	Goal       string    `json:"goal"`
	Status     string    `json:"status"`
	Type       string    `json:"type"`
	Sentiment  string    `json:"sentiment"`
	Time       time.Time `json:"time"`
	Date       time.Time `json:"date"`
	NewGroupId int       `json:"new_group_id"`
	UserId     int       `json:"user_id"`
	Token      string    `json:"token"`
}

type GoalsResponse struct {
	ID         int       `json:"id"`
	Goal       string    `json:"goal"`
	Status     string    `json:"status"`
	Type       string    `json:"type"`
	Sentiment  string    `json:"sentiment"`
	Time       time.Time `json:"time"`
	Date       time.Time `json:"date"`
	NewGroupId int       `json:"new_group_id"`
	UserId     int       `json:"user_id"`
}

func (g Goals) ToResponse() GoalsResponse {
	return GoalsResponse{
		ID:         g.ID,
		Goal:       g.Goal,
		Status:     g.Status,
		Type:       g.Type,
		Sentiment:  g.Sentiment,
		Time:       g.Time,
		Date:       g.Date,
		NewGroupId: g.NewGroupId,
		UserId:     g.UserId,
	}
}

func NewGoalResponse(code int, status string, goals Goals) helper.ReturnResponse {
	userResponse := goals.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userResponse,
	}
}
