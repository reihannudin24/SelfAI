package group

import "book_store/app/model/helper"

type Group struct {
	ID     int    `json:"id"`
	Name   string `json:"goal"`
	Bio    string `json:"bio"`
	Link   string `json:"link"`
	Type   string `json:"type"`
	UserId int    `json:"user_id"`
}

type GroupResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"goal"`
	Bio    string `json:"bio"`
	Link   string `json:"link"`
	Type   string `json:"type"`
	UserId int    `json:"user_id"`
}

func (data Group) ToResponse() GroupResponse {
	return GroupResponse{
		ID:     data.ID,
		Name:   data.Type,
		Bio:    data.Type,
		Link:   data.Type,
		Type:   data.Type,
		UserId: data.UserId,
	}
}

func NewGroupResponse(code int, status string, group Group) helper.ReturnResponse {
	userResponse := group.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userResponse,
	}
}
