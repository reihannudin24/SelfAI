package activities

import (
	"book_store/app/model/activities/domain"
	helper2 "book_store/helper"
	"context"
	"database/sql"
)

type ActivitiesRepository interface {
	Create(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error)
	Update(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error)
	Delete(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error)
	Show(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) (domain.Activities, helper2.ResponseJson, error)
	All(ctx context.Context, tx *sql.Tx, activities domain.Activities, token string) ([]domain.Activities, helper2.ResponseJson, error)
}
