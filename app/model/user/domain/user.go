package domain

import (
	"book_store/app/model/helper"
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	EmailVerify bool      `json:"email_verify"`
	VerifyCode  string    `json:"verify_code"`
	PhoneNumber string    `json:"phone_number"`
	Level       int       `json:"level"`
	Point       int       `json:"point"`
	Theme       string    `json:"theme"`
	Bio         string    `json:"bio"`
	Birthday    time.Time `json:"birthday"`
	Password    string    `json:"password"`
	Type        string    `json:"type"`
	Token       string    `json:"token"`
}

type UserResponse struct {
	ID          int       `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	EmailVerify bool      `json:"email_verify"`
	PhoneNumber string    `json:"phone_number"`
	Level       int       `json:"level"`
	Theme       string    `json:"theme"`
	Point       int       `json:"point"`
	Bio         string    `json:"bio"`
	Birthday    time.Time `json:"birthday"`
	Type        string    `json:"type"`
}

func (u User) ToResponse() UserResponse {
	return UserResponse{
		ID:          u.ID,
		Firstname:   u.Firstname,
		Lastname:    u.Lastname,
		Username:    u.Username,
		Email:       u.Email,
		EmailVerify: u.EmailVerify,
		PhoneNumber: u.PhoneNumber,
		Level:       u.Level,
		Theme:       u.Theme,
		Point:       u.Point,
		Bio:         u.Bio,
		Birthday:    u.Birthday,
		Type:        u.Type,
	}
}

func NewUserResponse(code int, status string, user User) helper.ReturnResponse {
	userResponse := user.ToResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   userResponse,
	}
}
