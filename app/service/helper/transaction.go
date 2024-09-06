package helper

import (
	"book_store/app/model/helper"
	"book_store/config"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

// TransactionServiceImpl is the implementation of transaction service
type TransactionServiceImpl struct {
	DB       *sql.DB
	Validate *validator.Validate
}

// NewTransactionServiceImpl creates a new instance of TransactionServiceImpl
func NewTransactionServiceImpl(db *sql.DB, validate *validator.Validate) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		DB:       db,
		Validate: validate,
	}
}

// PerformTransaction executes a database transaction and manages commit/rollback
func (service *TransactionServiceImpl) PerformTransaction(ctx context.Context, task func(tx *sql.Tx) (helper.ReturnResponse, error)) helper.ReturnResponse {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return helper2.ErrorService(err, 500, "Internal Server Error", "Failed to start transaction")
	}
	defer config.CommitOrRollback(tx)

	response, err := task(tx)
	if err != nil {
		return helper2.ErrorService(err, 500, "Internal Server Error", err.Error())
	}

	return response
}
