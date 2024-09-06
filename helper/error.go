package helper

import (
	"book_store/app/model/helper"
	"fmt"
	"log"
	"net/http"
)

//
//type ResponseJson struct {
//	Status  string      `json:"status"`
//	Code    int         `json:"code"`
//	Message string      `json:"message"`
//	Error   interface{} `json:"error"`
//}
//
//func (e ResponseJson) ResponseMessageJson(status string, code int, message string, error interface{}) ResponseJson {
//	fmt.Printf("status : ", status)
//	fmt.Printf("codes : ", code)
//	fmt.Printf("message : ", message)
//	return ResponseJson{
//		Status:  status,
//		Code:    code,
//		Message: message,
//		Error:   error,
//	}
//}

type ResponseJson struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// Optional: If you need to use this method, make sure it's utilized correctly
func (e ResponseJson) ResponseMessageJson(status string, code int, message string, error interface{}) ResponseJson {
	return ResponseJson{
		Status:  status,
		Code:    code,
		Message: message,
		Error:   error,
	}
}

func ErrorController(w http.ResponseWriter, message string, err error) {
	if err != nil {
		http.Error(w, message, http.StatusInternalServerError)
	}
	return
}

func ErrorService(err error, code int, status string, message string) helper.ReturnResponse {
	if err != nil {
		return helper.ReturnResponse{
			Code:   code,
			Status: status,
			Data:   message,
		}
	}
	return helper.ReturnResponse{}
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorPrint(err error, message string) error {
	if err != nil {
		log.Printf("%s : %v", message, err)
	}
	return nil
}

func ErrorPrintObject(err error, object interface{}, message string) (interface{}, error) {
	if err != nil {
		log.Printf("%s : %v", message, err)
		return object, fmt.Errorf("%s: %w", message, err)
	}
	return object, err
}

func ErrorInter(data interface{}, err error) (interface{}, ResponseJson, error) {
	if err != nil {
		return data, ResponseJson{
			Status:  "ERROR",
			Code:    401,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	return data, ResponseJson{}, nil
}

func ErrorRequest(data interface{}, message string, err error) (interface{}, ResponseJson, error) {
	if err != nil {
		return data, ResponseJson{
			Status:  "ERROR",
			Code:    401,
			Message: "Request Error",
			Error:   message,
		}, err
	}
	return data, ResponseJson{}, nil
}

func ErrorServiceInternal(err error) interface{} {
	return helper.ReturnResponse{
		Code:    500,
		Status:  "ERROR",
		Message: "Internal Server Error",
		Data:    err.Error(),
	}
}

func ErrorServiceRequest(err error) interface{} {
	return helper.ReturnResponse{
		Code:    400,
		Status:  "ERROR",
		Message: "Bad Request",
		Data:    err.Error(),
	}
}

func ErrorServiceResponse(code int, status string, message string, data interface{}, err error) (interface{}, error) {
	return helper.ReturnResponse{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}, err
}

func SuccessResponse(data interface{}, message string, err error) (interface{}, ResponseJson, error) {

	return data, ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: message,
		Error:   err.Error(),
	}, nil
}
