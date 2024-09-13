package activities

import (
	request3 "book_store/app/model/activities/request"
	helper2 "book_store/app/model/helper"
	request2 "book_store/app/model/user/request"
	"book_store/app/service/activities"
	"book_store/helper"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ActivitiesController interface {
	Create(w http.ResponseWriter, r *http.Request, router httprouter.Param)
	Update(w http.ResponseWriter, r *http.Request, router httprouter.Param)
	Delete(w http.ResponseWriter, r *http.Request, router httprouter.Param)
	Show(w http.ResponseWriter, r *http.Request, router httprouter.Param)
	All(w http.ResponseWriter, r *http.Request, router httprouter.Param)
}

type ActivitiesControllerImpl struct {
	Service activities.ActivitiesService
}

func NewActivitiesController(service activities.ActivitiesService) *ActivitiesControllerImpl {
	return &ActivitiesControllerImpl{
		Service: service,
	}
}

func (controller ActivitiesControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request3.CreateActivities{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.Create(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}

func (controller ActivitiesControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request3.UpdateActivities{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.Update(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}

func (controller ActivitiesControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request3.DeleteActivities{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.Delete(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}

func (controller ActivitiesControllerImpl) Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.AccSessionAuth{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	slug := r.URL.Query().Get("slug")
	if slug == "" {
		helper.ErrorController(w, "Slug is required", nil)
		return
	}

	response := controller.Service.Show(r.Context(), request, slug, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}

func (controller ActivitiesControllerImpl) All(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request2.AccSessionAuth{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.All(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}
