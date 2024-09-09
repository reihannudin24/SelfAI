package user

import (
	"book_store/app/model/user/domain"
	"book_store/app/repository/helper"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"fmt"
)

type UserRepositoryImpl struct {
	DB               *sql.DB
	helperRepository *helper.Repository
}

func NewUserRepository(db *sql.DB) UserRepository {
	if db == nil {
		panic("database connection is nil")
	}

	return &UserRepositoryImpl{
		DB: db,
		helperRepository: &helper.Repository{
			DB: db,
		},
	}
}

func (repository *UserRepositoryImpl) Profile(ctx context.Context, tx *sql.Tx, user domain.User, slug string) (domain.User, helper2.ResponseJson, error) {
	var SQL string
	var row *sql.Row
	//
	//if slug != "" {
	//	SQL = "SELECT email, username, firstname, lastname, phone_number, level, point, theme, bio, birthday, type FROM users WHERE  = ?"
	//	row = tx.QueryRowContext(ctx, SQL, slug)
	//} else {
	SQL = "SELECT email, username, firstname, lastname, phone_number, level, point, theme, bio,  type FROM users WHERE token = ?"
	row = tx.QueryRowContext(ctx, SQL, user.Token)
	//}

	err := row.Scan(
		&user.Email,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.PhoneNumber,
		&user.Level,
		&user.Point,
		&user.Theme,
		&user.Bio,
		&user.Type,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, helper2.ResponseJson{
				Status:  "ERROR",
				Code:    404,
				Message: "User not found",
				Error:   "No rows found for the given parameters",
			}, nil
		}
		return user, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to fetch user data",
			Error:   err.Error(),
		}, fmt.Errorf("failed to fetch user: %w", err)
	}

	return user, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully fetched user data",
		Error:   "",
	}, nil
}

func (repository *UserRepositoryImpl) ProfileEmotionalDay(ctx context.Context, tx *sql.Tx, emotion domain.UserEmotion) (domain.UserEmotion, helper2.ResponseJson, error) {
	//TODO implement me
	panic("implement me")
}
