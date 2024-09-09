package helper

import (
	"book_store/app/model/user/domain"
	"book_store/helper"
	"database/sql"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

func (repo *Repository) ValidateExists(table string, column string, value interface{}) (bool, domain.User, error) {
	if repo.DB == nil {
		return false, domain.User{}, fmt.Errorf("database connection is nil")
	}

	query := fmt.Sprintf("SELECT id, email FROM %s WHERE %s = ? LIMIT 1", table, column)
	var user domain.User
	err := repo.DB.QueryRow(query, value).Scan(&user.ID, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, domain.User{}, nil // No rows found
		}
		return false, domain.User{}, fmt.Errorf("error validating record: %w", err)
	}

	return true, user, nil
}

func (repo *Repository) ValidateExistsGetPassword(table string, column string, value interface{}) (bool, domain.User, error) {
	if repo.DB == nil {
		return false, domain.User{}, fmt.Errorf("database connection is nil")
	}

	query := fmt.Sprintf("SELECT id, email , password FROM %s WHERE %s = ? LIMIT 1", table, column)
	var user domain.User
	err := repo.DB.QueryRow(query, value).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, domain.User{}, nil // No rows found
		}
		return false, domain.User{}, fmt.Errorf("error validating record: %w", err)
	}

	return true, user, nil
}

func (repo *Repository) ValidateUserExists(table string, column string, value interface{}) (bool, domain.User, error) {
	if repo.DB == nil {
		return false, domain.User{}, fmt.Errorf("database connection is nil")
	}

	query := fmt.Sprintf("SELECT id, email, verify_code, email_verify FROM %s WHERE %s = ? LIMIT 1", table, column)
	var user domain.User
	err := repo.DB.QueryRow(query, value).Scan(&user.ID, &user.Email, &user.VerifyCode, &user.EmailVerify)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, domain.User{}, nil // No rows found
		}
		return false, domain.User{}, fmt.Errorf("error validating record: %w", err)
	}

	return true, user, nil
}

func (repo *Repository) ValidateDontExists(table string, column string, value interface{}) (bool, error) {
	if repo.DB == nil {
		return false, fmt.Errorf("database connection is nil")
	}

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", table, column)
	var count int
	err := repo.DB.QueryRow(query, value).Scan(&count)
	if err != nil {
		helper.ErrorPrint(err, "Error validating record")
		return false, nil
	}

	return count == 0, nil
}

func (repo *Repository) ValidateDuplicate(table string, column string, value interface{}, input string, validate string) (bool, error) {
	if repo.DB == nil {
		return false, fmt.Errorf("database connection is nil")
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", validate, table, column)
	var dbValue string

	row := repo.DB.QueryRow(query, value)
	err := row.Scan(&dbValue)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if dbValue == input {
		return false, fmt.Errorf("%s duplicate", column)
	}

	return true, nil
}
