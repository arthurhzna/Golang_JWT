package repository

import (
	"context"
	"database/sql"
	"errors"
	"rest_api_golang/helper"
	"rest_api_golang/model/domain"
)

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "SELECT ID, username, email FROM users WHERE email = $1 AND password = $2"
	rows, err := tx.QueryContext(ctx, SQL, user.Email, user.Password)
	helper.ErrorConditionCheck(err)
	defer rows.Close()

	if rows.Next() {
		return rows.Scan(&user.ID, &user.Username, &user.Email)
	} else {
		return user
	}
}

func (repository *userRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, user.Username, user.Email, user.Password).Scan(&id) 
	helper.ErrorConditionCheck(err)
	user.ID = id
	return user 
}

func (repository *userRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT username, email FROM users WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.ErrorConditionCheck(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Email)
		helper.ErrorConditionCheck(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *userRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, username, email FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.ErrorConditionCheck(err)
	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		helper.ErrorConditionCheck(err)
		users = append(users, user)
	}
	return users
}
