package group

import (
	"book_store/app/model/groups/domain"
	domain2 "book_store/app/model/user/domain"
	"book_store/app/repository/helper"
	helper2 "book_store/helper"
	"context"
	"database/sql"
	"fmt"
)

type GroupRepositoryImpl struct {
	DB               *sql.DB
	helperRepository *helper.Repository
}

func NewGroupRepository(db *sql.DB) GroupRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &GroupRepositoryImpl{
		DB: db,
		helperRepository: &helper.Repository{
			DB: db,
		},
	}
}

// Helper function to generate response JSON
func createResponse(status string, code int, message string, err error) helper2.ResponseJson {
	return helper2.ResponseJson{
		Status:  status,
		Code:    code,
		Message: message,
		Error: func() string {
			if err != nil {
				return err.Error()
			} else {
				return ""
			}
		}(),
	}
}

// Helper function to check if transaction is nil
func checkTransaction(tx *sql.Tx) error {
	if tx == nil {
		return fmt.Errorf("transaction is nil")
	}
	return nil
}

func (repository *GroupRepositoryImpl) validateUser(token string) (bool, domain2.User, error) {
	return repository.helperRepository.ValidateExists("users", "token", token)
}

func (repository *GroupRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper2.ResponseJson, error) {
	if err := checkTransaction(tx); err != nil {
		return group, createResponse("ERROR", 500, "Transaction is nil", err), err
	}

	exists, foundUser, err := repository.validateUser(token)
	if err != nil {
		return group, createResponse("ERROR", 500, "Internal Server Error", err), err
	}
	if !exists {
		return group, createResponse("ERROR", 404, "User not found", fmt.Errorf("user not found")), nil
	}

	SQL := "INSERT INTO group(name, link, type, user_id) VALUES (?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, group.Name, group.Link, group.Type, foundUser.ID)
	if err != nil {
		return group, createResponse("ERROR", 500, fmt.Sprintf("Failed to create group: %s", err.Error()), err), err
	}

	return group, createResponse("OK", 201, "Successfully created group", nil), nil
}

func (repository *GroupRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper2.ResponseJson, error) {
	if err := checkTransaction(tx); err != nil {
		return group, createResponse("ERROR", 500, "Transaction is nil", err), err
	}

	exists, foundUser, err := repository.validateUser(token)
	if err != nil {
		return group, createResponse("ERROR", 500, "Internal Server Error", err), err
	}
	if !exists {
		return group, createResponse("ERROR", 404, "User not found", fmt.Errorf("user not found")), nil
	}

	SQL := "UPDATE group SET name = ?, link = ?, type = ?, user_id = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, group.Name, group.Link, group.Type, foundUser.ID, group.ID)
	if err != nil {
		return group, createResponse("ERROR", 500, fmt.Sprintf("Failed to update group: %s", err.Error()), err), err
	}

	return group, createResponse("OK", 201, "Successfully updated group", nil), nil
}

func (repository *GroupRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper2.ResponseJson, error) {
	if err := checkTransaction(tx); err != nil {
		return group, createResponse("ERROR", 500, "Transaction is nil", err), err
	}

	exists, _, err := repository.validateUser(token)
	if err != nil {
		return group, createResponse("ERROR", 500, "Internal Server Error", err), err
	}
	if !exists {
		return group, createResponse("ERROR", 404, "User not found", fmt.Errorf("user not found")), nil
	}

	SQL := "DELETE FROM group WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, group.ID)
	if err != nil {
		return group, createResponse("ERROR", 500, fmt.Sprintf("Failed to delete group: %s", err.Error()), err), err
	}

	return group, createResponse("OK", 201, "Successfully deleted group", nil), nil
}

func (repository *GroupRepositoryImpl) Join(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper2.ResponseJson, error) {
	if err := checkTransaction(tx); err != nil {
		return group, createResponse("ERROR", 500, "Transaction is nil", err), err
	}

	exists, foundUser, err := repository.validateUser(token)
	if err != nil {
		return group, createResponse("ERROR", 500, "Internal Server Error", err), err
	}
	if !exists {
		return group, createResponse("ERROR", 404, "User not found", fmt.Errorf("user not found")), nil
	}

	exists, foundGroup, err := repository.helperRepository.ValidateExists("group", "group", group.ID)
	if err != nil {
		return group, createResponse("ERROR", 500, "Internal Server Error", err), err
	}
	if !exists {
		return group, createResponse("ERROR", 404, "Group not found", fmt.Errorf("group not found")), nil
	}

	SQL := "INSERT INTO pivot_groups(status, user_id, group_id) VALUES (?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, "member", foundUser.ID, foundGroup.ID)
	if err != nil {
		return group, createResponse("ERROR", 500, fmt.Sprintf("Failed to join group: %s", err.Error()), err), err
	}

	return group, createResponse("OK", 201, "Successfully joined group", nil), nil
}

func (repository *GroupRepositoryImpl) Kick(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper2.ResponseJson, error) {
	if err := checkTransaction(tx); err != nil {
		return group, createResponse("ERROR", 500, "Transaction is nil", err), err
	}

	exists, foundUser, err := repository.validateUser(token)
	if err != nil {
		return group, createResponse("ERROR", 500, "Internal Server Error", err), err
	}
	if !exists {
		return group, createResponse("ERROR", 404, "User not found", fmt.Errorf("user not found")), nil
	}

	exists, foundGroup, err := repository.helperRepository.ValidateExists("group", "group", group.ID)
	if err != nil {
		return group, createResponse("ERROR", 500, "Internal Server Error", err), err
	}
	if !exists {
		return group, createResponse("ERROR", 404, "Group not found", fmt.Errorf("group not found")), nil
	}

	SQL := "DELETE FROM pivot_groups WHERE user_id = ? AND group_id = ?"
	_, err = tx.ExecContext(ctx, SQL, foundUser.ID, foundGroup.ID)
	if err != nil {
		return group, createResponse("ERROR", 500, fmt.Sprintf("Failed to kick user from group: %s", err.Error()), err), err
	}

	return group, createResponse("OK", 201, "Successfully kicked user from group", nil), nil
}

func (repository *GroupRepositoryImpl) Show(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper2.ResponseJson, error) {
	exists, foundUser, err := repository.validateUser(token)
	if err != nil {
		return group, createResponse("ERROR", 500, "Internal Server Error", err), err
	}
	if !exists {
		return group, createResponse("ERROR", 404, "User not found", fmt.Errorf("user not found")), nil
	}

	SQL := "SELECT id, name, link, type FROM group WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, foundUser.ID)
	if err != nil {
		return group, createResponse("ERROR", 500, "Failed to retrieve group", err), err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&group.ID, &group.Name, &group.Link, &group.Type)
		if err == sql.ErrNoRows {
			return group, createResponse("ERROR", 404, "Group not found", fmt.Errorf("group not found")), nil
		} else if err != nil {
			return group, createResponse("ERROR", 500, "Failed to parse group", err), err
		}
	}

	return group, createResponse("OK", 200, "Successfully retrieved group", nil), nil
}

func (repository *GroupRepositoryImpl) All(ctx context.Context, tx *sql.Tx, group domain.Group, token string) ([]domain.Group, helper2.ResponseJson, error) {
	var groups []domain.Group

	exists, foundUser, err := repository.validateUser(token)
	if err != nil {
		return nil, createResponse("ERROR", 500, "Internal Server Error", err), err
	}
	if !exists {
		return nil, createResponse("ERROR", 404, "User not found", fmt.Errorf("user not found")), nil
	}

	SQL := "SELECT id, name, link, type FROM group WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, foundUser)
	if err != nil {
		return nil, createResponse("ERROR", 500, "Failed to retrieve groups", err), err
	}
	defer rows.Close()

	for rows.Next() {
		var g domain.Group
		err = rows.Scan(&g.ID, &g.Name, &g.Link, &g.Type)
		if err != nil {
			return nil, createResponse("ERROR", 500, "Failed to parse group", err), err
		}
		groups = append(groups, g)
	}

	return groups, createResponse("OK", 200, "Successfully retrieved groups", nil), nil
}

func (repository *GroupRepositoryImpl) Invite(ctx context.Context, tx *sql.Tx, group domain.Group, token string) (domain.Group, helper2.ResponseJson, error) {
	//TODO implement me
	panic("implement me")
}
