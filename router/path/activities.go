package path

import (
	activities3 "book_store/app/controller/activities"
	"book_store/app/repository/activities"
	activities2 "book_store/app/service/activities"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func ActivitiesRouter(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {

	activityRepository := activities.NewActivitiesRepository(db)
	activitiesService := activities2.NewActivitiesService(db, validate, activityRepository)
	activitiesController := activities3.NewActivitiesController(activitiesService)

	router.POST("/api/activities/create", activitiesController.Create)
	router.PUT("/api/activities/update", activitiesController.Update)
	router.DELETE("/api/activities/delete", activitiesController.Delete)
	router.GET("/api/activities/show", activitiesController.Show)
	router.GET("/api/activities", activitiesController.All)

}
