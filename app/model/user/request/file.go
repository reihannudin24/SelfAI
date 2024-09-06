package request

import "mime/multipart"

// UploadFile represents the request payload for uploading a file.
type UploadFile struct {
	FileName   string                `validate:"required,min=5,max=150" json:"file_name"`
	FileHeader *multipart.FileHeader `validate:"required" json:"file_header"`
}
