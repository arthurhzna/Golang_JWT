package service

import (
	"context"
	"rest_api_golang/model/web"	
)

type UserService interface {
	Login(ctx context.Context, request web.UserLoginRequest) web.UserTokenResponse
	Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse 
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
	ValidateToken(tokenString string) (*JwtClaims, error)
	// RefreshToken(ctx context.Context, oldToken string) web.UserTokenResponse
	// LogoutUser(ctx context.Context, tokenString string) web.WebResponse
}
