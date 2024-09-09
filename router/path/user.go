package path

import (
	user3 "book_store/app/controller/user"
	"book_store/app/repository/user"
	user2 "book_store/app/service/user"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func UserRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {

	userRepository := user.NewUserRepository(db)
	userService := user2.NewUserService(db, validate, userRepository)
	userController := user3.NewUserController(userService)

	router.GET("/api/user/profile/:slug", userController.Profile)

}
