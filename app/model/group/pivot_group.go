package group

import "book_store/app/model/helper"

type PivotGroup struct {
	ID         int    `json:"id"`
	Status     string `json:"status"`
	UserId     int    `json:"user_id"`
	NewGroupId int    `json:"new_group_id"`
}

type PivotGroupResponse struct {
	ID         int    `json:"id"`
	Status     string `json:"status"`
	UserId     int    `json:"user_id"`
	NewGroupId int    `json:"new_group_id"`
}

func (data PivotGroup) ToResponse() PivotGroupResponse {
	return PivotGroupResponse{
		ID:         data.ID,
		Status:     data.Status,
		NewGroupId: data.NewGroupId,
		UserId:     data.UserId,
	}
}

func NewPivotGroupResponse(code int, status string, group PivotGroup) helper.ReturnResponse {
	userResponse := group.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userResponse,
	}
}
