package repository

import (
	"context"
	"database/sql"
	"rest_api_golang/model/domain"
)

type CategoryRepository interface {
	Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}


