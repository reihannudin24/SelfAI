package domain

import (
	"book_store/app/model/helper"
	"time"
)

type Leaderboard struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	NewGroupId int       `json:"new_group_id"`
}

type LeaderboardResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	NewGroupId int       `json:"new_group_id"`
}

func (data Leaderboard) ToResponse() LeaderboardResponse {
	return LeaderboardResponse{
		ID:         data.ID,
		Name:       data.Name,
		StartDate:  data.StartDate,
		EndDate:    data.EndDate,
		NewGroupId: data.NewGroupId,
	}
}

func NewLeaderboardResponse(code int, status string, leaderboard Leaderboard) helper.ReturnResponse {
	userResponse := leaderboard.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userResponse,
	}
}
