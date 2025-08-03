package service

import (
	"context"
	"golang_jwt/model/web"	
)

type UserService interface {
	Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
	Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse 
}
