package helper

import (
	"book_store/app/model/user/domain"
	"context"
	"database/sql"
)

type RelationRepository interface {
	GetUserById(ctx context.Context, tx *sql.Tx, userId int, token string) ([]domain.User, error)
}

type RelationRepositoryImpl struct {
	DB               *sql.DB
	helperRepository *Repository
}

func NewRelationRepository(db *sql.DB) RelationRepository {
	return &RelationRepositoryImpl{
		DB: db,
		helperRepository: &Repository{
			DB: db,
		},
	}
}

func (repository *RelationRepositoryImpl) GetUserById(ctx context.Context, tx *sql.Tx, userId int, token string) ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}
