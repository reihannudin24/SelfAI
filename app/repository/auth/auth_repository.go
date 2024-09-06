package auth

import (
	"book_store/app/model/user/domain"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"mime/multipart"
)

type AuthRepository interface {
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, helper2.ResponseJson, error)

	Register(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, helper2.ResponseJson, error)
	SendVerifyCode(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error)
	VerifyEmail(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error)
	AddPassword(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error)
	AddInformation(ctx context.Context, tx *sql.Tx, user domain.User, token string) (domain.User, helper2.ResponseJson, error)
	AddOptionalInformation(ctx context.Context, tx *sql.Tx, user domain.User, token string, fileHeader *multipart.FileHeader) (domain.User, helper2.ResponseJson, error)
}
