package group

import (
	"book_store/app/model/groups/domain"
	"book_store/helper"
	"context"
	"database/sql"
)

// GroupRepository defines the contract for group-related data operations.
type GroupRepository interface {
	Create(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper.ResponseJson, error)
	Update(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper.ResponseJson, error)
	Delete(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper.ResponseJson, error)
	Join(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper.ResponseJson, error)
	Kick(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper.ResponseJson, error)
	Invite(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper.ResponseJson, error)
	Show(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper.ResponseJson, error)
	All(ctx context.Context, tx *sql.Tx, group domain.Group, token string) ([]domain.Group, helper.ResponseJson, error)
}
