package activities

import (
	"book_store/app/model/activities/domain"
	"book_store/app/repository/helper"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"fmt"
)

type ActivitiesRepositoryImpl struct {
	DB               *sql.DB
	helperRepository *helper.Repository
}

func NewActivitiesRepository(db *sql.DB) ActivitiesRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &ActivitiesRepositoryImpl{
		DB: db,
		helperRepository: &helper.Repository{
			DB: db,
		},
	}
}

func helperSuccessResponse(data domain.Activities) (domain.Activities, helper2.ResponseJson, error) {
	return data, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "Activity successfully added",
		Error:   "",
	}, nil
}

func helperSuccessShowResponse(data []domain.Activities) ([]domain.Activities, helper2.ResponseJson, error) {
	return data, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Activities retrieved successfully",
		Error:   "",
	}, nil
}

func helperErrorResponse(data domain.Activities, err error, status string, code int, message string) (domain.Activities, helper2.ResponseJson, error) {
	var errMessage string
	if err != nil {
		errMessage = err.Error()
	}
	return data, helper2.ResponseJson{
		Status:  status,
		Code:    code,
		Message: message,
		Error:   errMessage,
	}, fmt.Errorf(message)
}

func helperErrorShowResponse(data []domain.Activities, err error, status string, code int, message string) ([]domain.Activities, helper2.ResponseJson, error) {
	var errMessage string
	if err != nil {
		errMessage = err.Error()
	}
	return data, helper2.ResponseJson{
		Status:  status,
		Code:    code,
		Message: message,
		Error:   errMessage,
	}, fmt.Errorf(message)
}

func (repository *ActivitiesRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error) {
	if tx == nil {
		return helperErrorResponse(activities, nil, "ERROR", 500, "Internal server error")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 500, "Internal server error")
	}
	if !exists {
		return helperErrorResponse(activities, nil, "ERROR", 404, "User not found")
	}

	SQL := "INSERT INTO activities(title, slug, content, address, start_time, end_time, date, remainder, type, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, activities.Title, activities.Slug, activities.Content, activities.Address, activities.StartTime, activities.EndTime, activities.Date, activities.Remainder, activities.Type, foundUser.ID)
	if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 400, "Failed to add activity")
	}

	return helperSuccessResponse(activities)
}

func (repository *ActivitiesRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error) {
	if tx == nil {
		return helperErrorResponse(activities, nil, "ERROR", 500, "Internal server error")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 500, "Internal server error")
	}

	if !exists {
		return helperErrorResponse(activities, nil, "ERROR", 404, "User not found")
	}

	activitiesExist, _, err := repository.helperRepository.ValidateExists("activities", "id", activities.ID)
	if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 500, "Internal server error")
	}
	if !activitiesExist {
		return helperErrorResponse(activities, nil, "ERROR", 404, "Activity not found")
	}

	SQL := "UPDATE activities SET title = ?, slug  = ?, content  = ?, address  = ?, start_time  = ?, end_time  = ?, date  = ?, remainder  = ?, type  = ? WHERE id = ? AND user_id = ?"
	_, err = tx.ExecContext(ctx, SQL, activities.Title, activities.Slug, activities.Content, activities.Address, activities.StartTime, activities.EndTime, activities.Date, activities.Remainder, activities.Type, activities.ID, foundUser.ID)
	if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 400, "Failed to update activity")
	}

	return helperSuccessResponse(activities)
}

func (repository *ActivitiesRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error) {
	if tx == nil {
		return helperErrorResponse(activities, nil, "ERROR", 500, "Internal server error")
	}

	exists, _, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 500, "Internal server error")
	}

	if !exists {
		return helperErrorResponse(activities, nil, "ERROR", 404, "User not found")
	}

	activitiesExist, _, err := repository.helperRepository.ValidateExists("activities", "id", activities.ID)
	if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 500, "Internal server error")
	}
	if !activitiesExist {
		return helperErrorResponse(activities, nil, "ERROR", 404, "Activity not found")
	}

	SQL := "DELETE FROM activities WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, activities.ID)
	if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 400, "Failed to delete activity")
	}

	return helperSuccessResponse(activities)
}

func (repository *ActivitiesRepositoryImpl) Show(ctx context.Context, tx *sql.Tx, activities domain.Activities, slug string, token string) (domain.Activities, helper2.ResponseJson, error) {
	if tx == nil {
		return helperErrorResponse(activities, nil, "ERROR", 500, "Internal server error")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 500, "Internal server error")
	}
	if !exists {
		return helperErrorResponse(activities, nil, "ERROR", 404, "User not found")
	}

	SQL := "SELECT id, title, slug, content, address, start_time, end_time, date, remainder, type, user_id FROM activities WHERE id = ? AND user_id = ?"
	row := tx.QueryRowContext(ctx, SQL, activities.ID, foundUser.ID)
	err = row.Scan(&activities.ID, &activities.Title, &activities.Slug, &activities.Content, &activities.Address, &activities.StartTime, &activities.EndTime, &activities.Date, &activities.Remainder, &activities.Type, &foundUser.ID)
	if err == sql.ErrNoRows {
		return helperErrorResponse(activities, nil, "ERROR", 404, "Activity not found")
	} else if err != nil {
		return helperErrorResponse(activities, err, "ERROR", 500, "Failed to retrieve activity")
	}

	return helperSuccessResponse(activities)
}

func (repository *ActivitiesRepositoryImpl) All(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) ([]domain.Activities, helper2.ResponseJson, error) {
	var activitiesList []domain.Activities

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return helperErrorShowResponse(activitiesList, err, "ERROR", 500, "Internal server error")
	}
	if !exists {
		return helperErrorShowResponse(activitiesList, nil, "ERROR", 404, "User not found")
	}

	SQL := "SELECT id, title, slug, content, address, start_time, end_time, date, remainder, type, user_id FROM activities WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, foundUser.ID)
	if err != nil {
		return helperErrorShowResponse(activitiesList, err, "ERROR", 500, "Failed to retrieve activities")
	}
	defer rows.Close()

	for rows.Next() {
		var activity domain.Activities
		err = rows.Scan(&activity.ID, &activity.Title, &activity.Slug, &activity.Content, &activity.Address, &activity.StartTime, &activity.EndTime, &activity.Date, &activity.Remainder, &activity.Type, &foundUser.ID)
		if err != nil {
			return helperErrorShowResponse(activitiesList, err, "ERROR", 500, "Failed to retrieve activities")
		}
		activitiesList = append(activitiesList, activity)
	}

	return helperSuccessShowResponse(activitiesList)
}
