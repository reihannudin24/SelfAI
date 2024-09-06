package path

import (
	auth3 "book_store/app/controller/auth"
	"book_store/app/repository/auth"
	auth2 "book_store/app/service/auth"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func AuthRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	authRepository := auth.NewAuthRepository(db)
	authService := auth2.NewAuthService(db, validate, authRepository)
	authController := auth3.NewAuthController(authService)

	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/auth/send_verify_code", authController.SendVerifyCode)
	router.POST("/api/auth/verify_email", authController.VerifyEmail)
	router.POST("/api/auth/add_password", authController.AddPassword)
	router.POST("/api/auth/add_information", authController.AddInformation)
	router.POST("/api/auth/add_optional_information", authController.AddOptionalInformation)
}
