package goals

import (
	"book_store/app/model/goals/domain"
	"book_store/app/repository/helper"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"fmt"
)

type GoalRepositoryImpl struct {
	DB               *sql.DB
	helperRepository *helper.Repository
}

func NewGoalRepository(db *sql.DB) GoalRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &GoalRepositoryImpl{
		DB: db,
		helperRepository: &helper.Repository{
			DB: db,
		},
	}
}

func (repository *GoalRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) (domain.Goals, helper2.ResponseJson, error) {
	if tx == nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "INSERT INTO goals(goal, process, status, type, sentiment, time, date, new_group_id, user_id) VALUES (?, ?, ?, ? ,?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, goals.Goal, 0, goals.Status, goals.Type, goals.Sentiment, goals.Time, goals.Date, goals.NewGroupId, foundUser.ID)
	if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: fmt.Sprintf("Failed to add goal: %s", err.Error()),
			Error:   err.Error(),
		}, err
	}

	return goals, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "Successfully added goal",
		Error:   "",
	}, nil
}

func (repository *GoalRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) (domain.Goals, helper2.ResponseJson, error) {
	if tx == nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)

	if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	data_exist, _, err := repository.helperRepository.ValidateExists("goals", "id", goals.ID)
	if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !data_exist {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "Activity does not exist",
			Error:   "Activity not found",
		}, fmt.Errorf("activity not found")
	}

	SQL := "UPDATE goals SET goal  = ?,  status  = ?, type  = ?, sentiment  = ?, time  = ?, date  = ?, new_group_id  = ?, WHERE id = ? AND user_id  = ?"
	_, err = tx.ExecContext(ctx, SQL, goals.Goal, goals.Status, goals.Type, goals.Sentiment, goals.Time, goals.Date, goals.NewGroupId, goals.ID, foundUser.ID)
	if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: fmt.Sprintf("Failed to update activity: %s", err.Error()),
			Error:   err.Error(),
		}, err
	}

	return goals, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully updated goals",
		Error:   "",
	}, nil
}

func (repository *GoalRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) (domain.Goals, helper2.ResponseJson, error) {
	if tx == nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, _, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	data_exist, _, err := repository.helperRepository.ValidateExists("goals", "id", goals.ID)
	if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !data_exist {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "Activity does not exist",
			Error:   "Activity not found",
		}, fmt.Errorf("activity not found")
	}

	SQL := "DELETE FROM goals WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, goals.ID)
	if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: fmt.Sprintf("Failed to delete activity: %s", err.Error()),
			Error:   err.Error(),
		}, err
	}

	return goals, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully deleted goal",
		Error:   "",
	}, nil
}

func (repository *GoalRepositoryImpl) Show(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) (domain.Goals, helper2.ResponseJson, error) {
	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "SELECT id, goal, process, status, type, sentiment, time, date, new_group_id, FROM goals WHERE id = ? AND user_id = ?"
	row := tx.QueryRowContext(ctx, SQL, goals.ID, foundUser.ID)

	err = row.Scan(&goals.ID, &goals.Goal, &goals.Goal, &goals.Status, &goals.Type, &goals.Sentiment, &goals.Time, &goals.Date, &goals.NewGroupId, &goals.ID, &foundUser.ID)
	if err == sql.ErrNoRows {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "Activity not found",
			Error:   "Activity not found",
		}, fmt.Errorf("activity not found")
	} else if err != nil {
		return goals, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to retrieve activity",
			Error:   err.Error(),
		}, err
	}

	return goals, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully retrieved activity",
		Error:   "",
	}, nil
}

func (repository *GoalRepositoryImpl) All(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) ([]domain.Goals, helper2.ResponseJson, error) {

	var goalsList []domain.Goals

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return goalsList, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return goalsList, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "SELECT id, goal, process, status, type, sentiment, time, date, new_group_id, FROM goals WHERE id = ? "
	row := tx.QueryRowContext(ctx, SQL, goals.ID, foundUser.ID)

	err = row.Scan(&goals.ID, &goals.Goal, &goals.Goal, &goals.Status, &goals.Type, &goals.Sentiment, &goals.Time, &goals.Date, &goals.NewGroupId, &goals.ID, &foundUser.ID)
	if err == sql.ErrNoRows {
		return goalsList, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "Activity not found",
			Error:   "Activity not found",
		}, fmt.Errorf("activity not found")
	} else if err != nil {
		return goalsList, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to retrieve activity",
			Error:   err.Error(),
		}, err
	}

	return goalsList, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully retrieved goal",
		Error:   "",
	}, nil
}
