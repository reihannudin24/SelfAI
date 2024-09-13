package router

import (
	"book_store/router/path"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"log"
)

func ApiRouter(db *sql.DB, validate *validator.Validate) *httprouter.Router {

	router := httprouter.New()
	path.AuthRouter(router, db, validate)
	path.UserRouter(router, db, validate)
	path.ActivitiesRouter(router, db, validate)

	log.Printf("Running Localhost server: http://localhost:3000/")

	return router
}
