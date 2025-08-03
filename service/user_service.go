package service

import (
	"context"
	"golang_jwt/model/web"	
)

type UserService interface {
	Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse 
	Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
	Logout(ctx context.Context, sessionId string) 
	RenewAccessToken(ctx context.Context, request web.RenewAccessTokenRequest) web.RenewAccessTokenResponse
	RevokeSession(ctx context.Context, sessionId string)
}
