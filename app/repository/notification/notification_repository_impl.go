package notification

import (
	"book_store/app/model/notification/domain"
	"book_store/app/repository/helper"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"fmt"
)

type NotificationRepositoryImpl struct {
	DB               *sql.DB
	helperRepository *helper.Repository
}

func NewNotificationRepository(db *sql.DB) NotificationRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &NotificationRepositoryImpl{
		DB: db,
		helperRepository: &helper.Repository{
			DB: db,
		},
	}
}

func (repository *NotificationRepositoryImpl) Push(ctx context.Context, tx *sql.Tx, notification domain.Notification, token string) (domain.Notification, helper2.ResponseJson, error) {
	if tx == nil {
		return notification, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Transaction is nil",
			Error:   "Internal Server Error",
		}, fmt.Errorf("transaction is nil")
	}

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if err != nil {
		return notification, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Internal Server Error",
			Error:   err.Error(),
		}, err
	}
	if !exists {
		return notification, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "INSERT INTO notification(title, slug , content, category, r_id , user_id) VALUES (?, ?, ?, ? ,?, ?)"
	_, err = tx.ExecContext(ctx, SQL, notification.Title, notification.Content, notification.Content, notification.Content, notification.Content, foundUser.ID)
	if err != nil {
		return notification, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: fmt.Sprintf("Failed to add activity: %s", err.Error()),
			Error:   err.Error(),
		}, err
	}

	return notification, helper2.ResponseJson{
		Status:  "OK",
		Code:    201,
		Message: "Successfully added notification",
		Error:   "",
	}, nil

}

func (repository *NotificationRepositoryImpl) Show(ctx context.Context, tx *sql.Tx, notification domain.Notification, token string) (domain.Notification, helper2.ResponseJson, error) {

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if !exists {
		return notification, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "SELECT id, title, slug, content, category, r_id FROM notifications WHERE id = ? AND user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, notification.ID, foundUser.ID)
	err = rows.Scan(&notification.ID, &notification.Title, &notification.Content, &notification.Content, &notification.Content, &foundUser.ID)
	if err == sql.ErrNoRows {
		return notification, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "Activity not found",
			Error:   "Activity not found",
		}, fmt.Errorf("activity not found")
	} else if err != nil {
		return notification, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to retrieve activity",
			Error:   err.Error(),
		}, err
	}

	return notification, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully retrieved notification",
		Error:   "",
	}, nil
}

func (repository *NotificationRepositoryImpl) All(ctx context.Context, tx *sql.Tx, notification domain.Notification, token string) ([]domain.Notification, helper2.ResponseJson, error) {
	var notificationsList []domain.Notification

	exists, foundUser, err := repository.helperRepository.ValidateExists("users", "token", token)
	if !exists {
		return notificationsList, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    404,
			Message: "User not found",
			Error:   "User not found",
		}, fmt.Errorf("user not found")
	}

	SQL := "SELECT id, title, slug, content, category, r_id FROM notifications WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, foundUser.ID)
	if err != nil {
		return notificationsList, helper2.ResponseJson{
			Status:  "ERROR",
			Code:    500,
			Message: "Failed to retrieve activities",
			Error:   err.Error(),
		}, err
	}
	defer rows.Close()

	for rows.Next() {
		var activity domain.Notification
		err = rows.Scan(&notification.ID, &notification.Title, &notification.Content, &notification.Content, &notification.Content, &foundUser.ID)
		if err != nil {
			return notificationsList, helper2.ResponseJson{
				Status:  "ERROR",
				Code:    500,
				Message: "Failed to parse activities",
				Error:   err.Error(),
			}, err
		}
		notificationsList = append(notificationsList, activity)
	}

	return notificationsList, helper2.ResponseJson{
		Status:  "OK",
		Code:    200,
		Message: "Successfully retrieved all notification",
		Error:   "",
	}, nil
}
