package goals

import (
	"book_store/app/model/goals/domain"
	helper2 "book_store/helper"
	"context"
	"database/sql"
)

type GoalRepository interface {
	Create(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) (domain.Goals, helper2.ResponseJson, error)
	Update(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) (domain.Goals, helper2.ResponseJson, error)
	Delete(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) (domain.Goals, helper2.ResponseJson, error)
	Show(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) (domain.Goals, helper2.ResponseJson, error)
	All(ctx context.Context, tx *sql.Tx, goals domain.Goals, token string) ([]domain.Goals, helper2.ResponseJson, error)
}
