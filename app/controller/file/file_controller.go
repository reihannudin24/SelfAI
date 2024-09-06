package file

import (
	helper2 "book_store/app/model/helper"
	request2 "book_store/app/model/user/request"
	"book_store/app/service/file"
	"book_store/helper"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type FileController interface {
	UploadFile(w http.ResponseWriter, r *http.Request, param httprouter.Param)
}

type FileControllerImpl struct {
	Service file.FileService
}

func NewFileController(service file.FileService) *FileControllerImpl {
	return &FileControllerImpl{
		Service: service,
	}
}

func (controller FileControllerImpl) UploadController(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	request := request2.UploadFile{}
	err := helper.ReadFromRequestBody(r, &request)
	helper.ErrorController(w, "Invalid request body", err)

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.UploadFile(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	helper.ErrorController(w, "Failed to write response", err)
}
