package domain

import "book_store/app/model/helper"

type Notification struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	RId     int    `json:"r_id"`
	UserId  int    `json:"user_id"`
}

type NotificationResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	RId     int    `json:"r_id"`
	UserId  int    `json:"user_id"`
}

func (n Notification) ToResponse() NotificationResponse {
	return NotificationResponse{
		ID:      n.ID,
		Title:   n.Title,
		Content: n.Content,
		RId:     n.RId,
		UserId:  n.UserId,
	}
}

func NewNotificationResponse(code int, status string, notification Notification) helper.ReturnResponse {
	userResponse := notification.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userResponse,
	}
}
