package domain

import "book_store/app/model/helper"

type PositionLeaderboard struct {
	ID             int    `json:"id"`
	Position       string `json:"position"`
	Leaderboard_id int    `json:"leaderboard_id"`
	UserId         int    `json:"user_id"`
	NewGroupId     int    `json:"new_group_id"`
}

type PositionLeaderboardResponse struct {
	ID             int    `json:"id"`
	Position       string `json:"position"`
	Leaderboard_id int    `json:"leaderboard_id"`
	UserId         int    `json:"user_id"`
	NewGroupId     int    `json:"new_group_id"`
}

func (data PositionLeaderboard) ToResponse() PositionLeaderboardResponse {
	return PositionLeaderboardResponse{
		ID:             data.ID,
		Position:       data.Position,
		Leaderboard_id: data.Leaderboard_id,
		NewGroupId:     data.NewGroupId,
		UserId:         data.UserId,
	}
}

func NewPositionLeaderboardResponse(code int, status string, leaderboard PositionLeaderboard) helper.ReturnResponse {
	userResponse := leaderboard.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userResponse,
	}
}
