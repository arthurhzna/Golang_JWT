package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_jwt/helper"
	"golang_jwt/model/domain"
)

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
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
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		helper.ErrorConditionCheck(err)
		users = append(users, user)
	}
	return users
}

func (repository *userRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id, username, email, password FROM users WHERE email = $1"
	row := tx.QueryRowContext(ctx, SQL, email)
	
	user := domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) CreateSession(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session {
	SQL := "INSERT INTO sessions (id, user_email, refresh_token, is_revoked, expires_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := tx.ExecContext(ctx, SQL, session.ID, session.User_Email, session.Refresh_Token, session.Is_Revoked, session.Expires_At)
	helper.ErrorConditionCheck(err)
	return session
}

func (repository *userRepositoryImpl) GetSession(ctx context.Context, tx *sql.Tx, id string) (domain.Session, error) {
	SQL := "SELECT * FROM sessions WHERE id = $1"
	row := tx.QueryRowContext(ctx, SQL, id)
	
	session := domain.Session{}
	err := row.Scan(&session.ID, &session.User_Email, &session.Refresh_Token, &session.Is_Revoked, &session.Created_At, &session.Expires_At)
	helper.ErrorConditionCheck(err)
	return session, nil
}

func (repository *userRepositoryImpl) RevokeSession(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "UPDATE sessions SET is_revoked = true WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.ErrorConditionCheck(err)
	return nil
}

func (repository *userRepositoryImpl) DeleteSession(ctx context.Context, tx *sql.Tx, id string) error {
	SQL := "DELETE FROM sessions WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.ErrorConditionCheck(err)
	return nil
}


