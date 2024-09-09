package mutuals

import "book_store/app/model/helper"

type Mutual struct {
	ID           int `json:"id"`
	MyUserId     int `json:"my_user_id"`
	FriendUserId int `json:"friend_user_id"`
}

type MutualResponse struct {
	ID           int `json:"id"`
	MyUserId     int `json:"my_user_id"`
	FriendUserId int `json:"friend_user_id"`
}

func (data Mutual) ToResponse() MutualResponse {
	return MutualResponse{
		ID:           data.ID,
		MyUserId:     data.MyUserId,
		FriendUserId: data.FriendUserId,
	}
}

func NewNMutualsResponse(code int, status string, mutual Mutual) helper.ReturnResponse {
	userResponse := mutual.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userResponse,
	}
}
