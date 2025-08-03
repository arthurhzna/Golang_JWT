package token

import (
	"context"
	"database/sql"
	"golang_jwt/model/domain"
)


type UserToken interface {
	GenerateToken(user domain.User) (string, *web.UserClaims, error)
	ValidateToken(tokenString string) (*web.UserClaims, error)
}