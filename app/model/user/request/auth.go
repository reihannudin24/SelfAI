package request

import (
	"mime/multipart"
	"time"
)

type Register struct {
	Email string `validate:"required,min=5,max=150" json:"email"`
}

type SendVerifyCode struct {
	Token string `validate:"required" json:"token"`
}

type VerifyEmail struct {
	Token      string `validate:"required" json:"token"`
	VerifyCode string `validate:"required,min=6,max=6" json:"verify_code"`
}

type AddPassword struct {
	Token           string `validate:"required" json:"token"`
	Password        string `validate:"required" json:"password"`
	ConfirmPassword string `validate:"required" json:"confirm_password"`
}

type AddInformation struct {
	Token       string `validate:"required" json:"token"`
	Firstname   string `validate:"required" json:"firstname"`
	Lastname    string `validate:"required" json:"lastname"`
	Username    string `validate:"required" json:"username"`
	PhoneNumber string `validate:"required" json:"phone_number"`
}

type AddOptionalInformation struct {
	Token      string                `json:"token" validate:"required"`
	Bio        string                `json:"bio" validate:"required"`
	Theme      string                `json:"theme" validate:"required"`
	Birthday   time.Time             `json:"birthday" validate:"required"`
	FileHeader *multipart.FileHeader `json:"fileHeader"` // Removed 'required' tag
}

type Login struct {
	Email    string `validate:"required,min=5,max=150" json:"email"`
	Password string `validate:"required" json:"password"`
}

type Logout struct {
	Token string `validate:"required" json:"token"`
}
