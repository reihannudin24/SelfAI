package path

import (
	file3 "book_store/app/controller/file"
	"book_store/app/repository/file"
	file2 "book_store/app/service/file"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func FileRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {

	fileRepository := file.NewUserRepository(db)
	fileService := file2.NewFileService(db, validate, fileRepository)
	fileController := file3.NewFileController(fileService)

	// Corrected route path with leading '/'
	router.POST("/api/file/upload", fileController.UploadController)
}
