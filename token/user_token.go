package token

import (
	"golang_jwt/model/web"
	"time"
)


type UserToken interface {
	GenerateToken(id int, username string, email string, duration time.Duration) (string, *web.UserClaims, error)
	ValidateToken(tokenString string) (*web.UserClaims, error)
}