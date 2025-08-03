package service

import (
    "context"
    "database/sql"
    "golang_jwt/model/web"
    "golang_jwt/model/domain"
	"golang_jwt/exception"
	"golang_jwt/helper"
    "golang_jwt/repository"
	"golang_jwt/token"
    "github.com/go-playground/validator/v10"
	"time"
)

type UserServiceImpl struct {
    UserRepository repository.UserRepository
    DB *sql.DB
    Validate *validator.Validate
	UserToken token.UserToken
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, Validate *validator.Validate, userToken token.UserToken) UserService {
	return  &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: Validate,
		UserToken: userToken,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.ErrorConditionCheck(err)
	tx, err := service.DB.Begin()
	helper.ErrorConditionCheck(err)
	defer helper.CommitOrRollback(tx)

	hashedPassword := helper.HashPassword(request.Password)

	user := domain.User{
		Username: request.Username,
		Email: request.Email,
		Password: hashedPassword,
	}

	user = service.UserRepository.Register(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse {
	err := service.Validate.Struct(request)
	helper.ErrorConditionCheck(err)

	tx, err := service.DB.Begin()
	helper.ErrorConditionCheck(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	helper.ErrorConditionCheck(err)

	helper.VerifyPassword(user.Password, request.Password)
	
	accessToken, accessClaims, err := service.UserToken.GenerateToken(user, 15*time.Minute)
	helper.ErrorConditionCheck(err)

    refreshToken, refreshClaims, err := service.UserToken.GenerateToken(user, 24*time.Hour)
    helper.ErrorConditionCheck(err)

	

	return helper.ToUserLoginResponse(accessToken, accessClaims)







	