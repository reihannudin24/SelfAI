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

func (repository *ActivitiesRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error) {
	if tx == nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "INSERT INTO activities(title, slug , content, address, start_time, end_time, date, remainder, type, user_id) VALUES (?, ?, ?, ? ,?, ?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, activities.Title, activities.Slug, activities.Content, activities.Address, activities.StartTime, activities.EndTime, activities.Date, activities.Remainder, activities.Type, foundUser.ID)
	if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: fmt.Sprintf("Failed to add activity: %s", err.Error()),
			Error:   err.Error(),
		}, err
	}

	return activities, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "Successfully added activity",
		Error:   "",
	}, nil
}

func (repository *ActivitiesRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error) {
	if tx == nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	activities_exist, _, err := repository.helperRepository.ValidateExists("activities", "id", activities.ID)
	if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !activities_exist {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "Activity does not exist",
			Error:   "Activity not found",
		}, fmt.Errorf("activity not found")
	}

	SQL := "UPDATE activities SET title = ?, slug  = ?, content  = ?, address  = ?, start_time  = ?, end_time  = ?, date  = ?, remainder  = ?, type  = ? WHERE id = ? AND user_id = ?"
	_, err = tx.ExecContext(ctx, SQL, activities.Title, activities.Slug, activities.Content, activities.Address, activities.StartTime, activities.EndTime, activities.Date, activities.Remainder, activities.Type, activities.ID, foundUser.ID)
	if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: fmt.Sprintf("Failed to update activity: %s", err.Error()),
			Error:   err.Error(),
		}, err
	}

	return activities, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully updated activity",
		Error:   "",
	}, nil
}

func (repository *ActivitiesRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error) {
	if tx == nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, _, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	activities_exist, _, err := repository.helperRepository.ValidateExists("activities", "id", activities.ID)
	if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !activities_exist {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "Activity does not exist",
			Error:   "Activity not found",
		}, fmt.Errorf("activity not found")
	}

	SQL := "DELETE FROM activities WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, activities.ID)
	if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: fmt.Sprintf("Failed to delete activity: %s", err.Error()),
			Error:   err.Error(),
		}, err
	}

	return activities, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully deleted activity",
		Error:   "",
	}, nil
}

func (repository *ActivitiesRepositoryImpl) Show(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error) {
	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "SELECT id, title, slug, content, address, start_time, end_time, date, remainder, type, user_id FROM activities WHERE id = ? AND user_id = ?"
	row := tx.QueryRowContext(ctx, SQL, activities.ID, foundUser.ID)
	err = row.Scan(&activities.ID, &activities.Title, &activities.Slug, &activities.Content, &activities.Address, &activities.StartTime, &activities.EndTime, &activities.Date, &activities.Remainder, &activities.Type, &foundUser.ID)
	if err == sql.ErrNoRows {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "Activity not found",
			Error:   "Activity not found",
		}, fmt.Errorf("activity not found")
	} else if err != nil {
		return activities, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to retrieve activity",
			Error:   err.Error(),
		}, err
	}

	return activities, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully retrieved activity",
		Error:   "",
	}, nil
}

func (repository *ActivitiesRepositoryImpl) All(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) ([]domain.Activities, helper2.ResponseJson, error) {
	var activitiesList []domain.Activities

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return activitiesList, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return activitiesList, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "SELECT id, title, slug, content, address, start_time, end_time, date, remainder, type, user_id FROM activities WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, foundUser.ID)
	if err != nil {
		return activitiesList, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to retrieve activities",
			Error:   err.Error(),
		}, err
	}
	defer rows.Close()

	for rows.Next() {
		var activity domain.Activities
		err = rows.Scan(&activity.ID, &activity.Title, &activity.Slug, &activity.Content, &activity.Address, &activity.StartTime, &activity.EndTime, &activity.Date, &activity.Remainder, &activity.Type, &foundUser.ID)
		if err != nil {
			return activitiesList, helper2.ResponseJson{
				Status:  "ERROR",
				Code:    500,
				Message: "Failed to parse activities",
				Error:   err.Error(),
			}, err
		}
		activitiesList = append(activitiesList, activity)
	}

	return activitiesList, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully retrieved all activities",
		Error:   "",
	}, nil
}
