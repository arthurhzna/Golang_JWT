package service

import (
    "context"
    "database/sql"
    "rest_api_golang/model/web"
    "rest_api_golang/model/domain"
	"rest_api_golang/exception"
	"rest_api_golang/helper"
    "rest_api_golang/repository"
    "github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
    UserRepository repository.UserRepository
    DB *sql.DB
    Validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, Validate *validator.Validate) UserService {
	return  &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: Validate,
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

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) web.UserTokenResponse {
	err := service.Validate.Struct(request)