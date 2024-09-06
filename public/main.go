package main

import (
	"book_store/config"
	"book_store/helper"
	"book_store/router"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {

	db := config.Connection()
	validate := validator.New()
	apiRouter := router.ApiRouter(db, validate)

	server := &http.Server{
		Addr:    ":3000",
		Handler: apiRouter,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
