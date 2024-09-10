package notification

import (
	"book_store/app/model/notification/domain"
	"book_store/helper"
	"context"
	"database/sql"
)

type NotificationRepository interface {
	Push(ctx context.Context, tx *sql.Tx, notification domain.Notification, token string) (domain.Notification, helper.ResponseJson, error)
	Show(ctx context.Context, tx *sql.Tx, notification domain.Notification, token string) (domain.Notification, helper.ResponseJson, error)
	All(ctx context.Context, tx *sql.Tx, notification domain.Notification, token string) ([]domain.Notification, helper.ResponseJson, error)
}
