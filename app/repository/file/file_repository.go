package file

import (
	"book_store/app/model/user/domain"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"mime/multipart"
)

type FileRepository interface {
	UploadFile(ctx context.Context, tx *sql.Tx, file domain.File, fileUser *multipart.FileHeader, token string) (domain.File, helper2.ResponseJson, error)
}
