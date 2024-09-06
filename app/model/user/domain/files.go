package domain

import "book_store/app/model/helper"

type File struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Type      string `json:"type"`
	Files     string `json:"files"`
	FilesPath string `json:"files_path"`
	UserId    string `json:"user_id"`
}

type Summarization struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Type    string `json:"type"`
	Try     int    `json:"try"`
	AI      string `json:"ai"`
	Temp    string `json:"temp"`
	Token   string `json:"token"`
	UserId  int    `json:"userId_id"`
	FileId  int    `json:"file_id"`
}

type FileResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Type          string `json:"type"`
	Files         string `json:"files"`
	FilesPath     string `json:"files_path"`
	Summarization []Summarization
	UserId        string `json:"user_id"`
}

func (f File) toResponse() FileResponse {
	return FileResponse{
		ID:            f.ID,
		Name:          f.Name,
		Slug:          f.Slug,
		Type:          f.Type,
		Files:         f.Files,
		FilesPath:     f.FilesPath,
		UserId:        f.UserId,
		Summarization: []Summarization{},
	}
}

func NewFilesResponse(code int, status string, file File) helper.ReturnResponse {
	fileResponse := file.toResponse()
	return helper.ReturnResponse{
		Code:   code,
		Status: status,
		Data:   fileResponse,
	}
}
