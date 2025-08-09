package repository

import (
	"context"
	"database/sql"
	"golang_jwt/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	CreateSession(ctx context.Context, tx *sql.Tx, session domain.Session) domain.Session
	GetSession(ctx context.Context, tx *sql.Tx, id string) (domain.Session, error)
	RevokeSession(ctx context.Context, tx *sql.Tx, id string) error
	DeleteSession(ctx context.Context, tx *sql.Tx, id string) error

	DeleteExpiredSessions(ctx context.Context, tx *sql.Tx) error
}


