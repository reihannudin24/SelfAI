package user

import (
	helper2 "book_store/app/model/helper"
	request2 "book_store/app/model/user/request"
	"book_store/app/service/user"
	"book_store/helper"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Profile(w http.ResponseWriter, r *http.Request, param httprouter.Param)
}

type UserControllerImpl struct {
	Service user.UserService
}

func NewUserController(service user.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		Service: service,
	}
}

func (controller UserControllerImpl) Profile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.AccSessionAuth{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	// Handling path parameter
	slug := params.ByName("slug")
	
	fmt.Printf("slug: %s\n", slug)

	response := controller.Service.Profile(r.Context(), request, slug)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}
