package user

import (
	"book_store/app/model/user/domain"
	helper2 "book_store/helper"
	"context"
	"database/sql"
)

type UserRepository interface {
	Profile(ctx context.Context, tx *sql.Tx, user domain.User, slug string) (domain.User, helper2.ResponseJson, error)
	ProfileEmotionalDay(ctx context.Context, tx *sql.Tx, emotion domain.UserEmotion) (domain.UserEmotion, helper2.ResponseJson, error)
}
