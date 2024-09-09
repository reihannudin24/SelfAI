package auth

import (
	helper2 "book_store/app/model/helper"
	request2 "book_store/app/model/user/request"
	auth "book_store/app/service/auth"
	"book_store/helper"
	"github.com/julienschmidt/httprouter"
	"mime/multipart"
	"net/http"
	"time"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request, param httprouter.Param)
	Logout(w http.ResponseWriter, r *http.Request, param httprouter.Param)
	Register(w http.ResponseWriter, r *http.Request, param httprouter.Param)
	SendVerifyCode(w http.ResponseWriter, r *http.Request, param httprouter.Param)
	VerifyEmail(w http.ResponseWriter, r *http.Request, param httprouter.Param)
	AddPassword(w http.ResponseWriter, r *http.Request, param httprouter.Param)
	AddInformation(w http.ResponseWriter, r *http.Request, param httprouter.Param)
	AddOptionalInformation(w http.ResponseWriter, r *http.Request, param httprouter.Param)
}

type AuthControllerImpl struct {
	Service auth.AuthService
}

func NewAuthController(service auth.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		Service: service,
	}
}

func (controller AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.Login{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	response := controller.Service.Login(r.Context(), request)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}

func (controller AuthControllerImpl) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.Logout{}
	err := helper.ReadFromRequestBody(r, &request)

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.Logout(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}

func (controller AuthControllerImpl) SendVerifyCode(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.SendVerifyCode{}
	err := helper.ReadFromRequestBody(r, &request)
	helper.ErrorController(w, "Invalid request body", err)

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.SendVerifyCode(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	helper.ErrorController(w, "Failed to write response", err)
}

func (controller AuthControllerImpl) VerifyEmail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.VerifyEmail{}
	err := helper.ReadFromRequestBody(r, &request)
	helper.ErrorController(w, "Invalid request body", err)

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.VerifyEmail(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	helper.ErrorController(w, "Failed to write response", err)
}

func (controller AuthControllerImpl) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.Register{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	response := controller.Service.Register(r.Context(), request)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}

func (controller AuthControllerImpl) AddPassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.AddPassword{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.AddPassword(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}

func (controller AuthControllerImpl) AddInformation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.AddInformation{}
	err := helper.ReadFromRequestBody(r, &request)
	helper.ErrorController(w, "Invalid request body", err)

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.AddInformation(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	helper.ErrorController(w, "Failed to write response", err)
}

func (controller AuthControllerImpl) AddOptionalInformation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseMultipartForm(10 << 20) // Set a limit for the size of the multipart form data
	if err != nil {
		helper.ErrorController(w, "Invalid multipart form", err)
		return
	}

	token, err := helper.GetHeaderAuth(r)
	if err != nil {
		helper.ErrorController(w, "Failed to get authorization header", err)
		return
	}

	bio := r.FormValue("bio")
	theme := r.FormValue("theme")
	birthdayStr := r.FormValue("birthday")

	dateFormats := []string{
		time.RFC3339,
		"2006-01-02",
		"01/02/2006",
	}

	var birthday time.Time
	var parseErr error
	for _, format := range dateFormats {
		birthday, parseErr = time.Parse(format, birthdayStr)
		if parseErr == nil {
			break
		}
	}
	if parseErr != nil {
		helper.ErrorController(w, "Invalid birthday format", parseErr)
		return
	}

	// Retrieve the file from the request
	_, fileHeader, err := r.FormFile("fileHeader")
	if err != nil && err != http.ErrMissingFile {
		helper.ErrorController(w, "Failed to get file", err)
		return
	}

	var fileHeaderPtr *multipart.FileHeader
	if err == nil {
		fileHeaderPtr = fileHeader
	}

	request := request2.AddOptionalInformation{
		Token:      token,
		Bio:        bio,
		Theme:      theme,
		Birthday:   birthday,
		FileHeader: fileHeaderPtr,
	}

	response := controller.Service.AddOptionalInformation(r.Context(), request, token)

	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}
