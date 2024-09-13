package goal

import (
	"book_store/app/model/goals/request"
	helper2 "book_store/app/model/helper"
	request2 "book_store/app/model/user/request"
	"book_store/app/service/goals"
	"book_store/helper"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type GoalController interface {
	AddGoal(w http.ResponseWriter, r *http.Request, router httprouter.Param)
	Update(w http.ResponseWriter, r *http.Request, router httprouter.Param)
	Delete(w http.ResponseWriter, r *http.Request, router httprouter.Param)
	Show(w http.ResponseWriter, r *http.Request, router httprouter.Param)
	All(w http.ResponseWriter, r *http.Request, router httprouter.Param)
}

type GoalControllerImpl struct {
	Service goals.GoalsService
}

func NewGoalController(service goals.GoalsService) *GoalControllerImpl {
	return &GoalControllerImpl{
		Service: service,
	}
}

func (controller GoalControllerImpl) AddGoals(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request.AddGoals{}
	err := helper.ReadFromRequestBody(r, &request)
	if err != nil {
		helper.ErrorController(w, "Invalid request body", err)
		return
	}

	token, err := helper.GetHeaderAuth(r)
	helper.ErrorController(w, "Failed to write response", err)

	response := controller.Service.Add(r.Context(), request, token)
	webResponse := helper2.ModReturnData{
		Data: response,
	}

	err = helper.WriteToResponseBody(w, webResponse)
	if err != nil {
		helper.ErrorController(w, "Failed to write response", err)
	}
}

func (controller GoalControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request.UpdateGoals{}
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

func (controller GoalControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := request.DeleteGoals{}
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

func (controller GoalControllerImpl) Show(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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

func (controller GoalControllerImpl) All(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
