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
	"errors"
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
	
	accessToken, accessClaims, err := service.UserToken.GenerateToken(user.ID, user.Username, user.Email, 15*time.Minute)
	helper.ErrorConditionCheck(err)

    refreshToken, refreshClaims, err := service.UserToken.GenerateToken(user.ID, user.Username, user.Email, 24*time.Hour)
    helper.ErrorConditionCheck(err)

	session := domain.Session{
		ID: refreshClaims.RegisteredClaims.ID,
		User_Email: user.Email,
		Refresh_Token: refreshToken,
		Is_Revoked: false,
		Expires_At: refreshClaims.RegisteredClaims.ExpiresAt.Time,
	}
	session = service.UserRepository.CreateSession(ctx, tx, session)
	helper.ErrorConditionCheck(err)

	return helper.ToUserLoginResponse(accessToken, accessClaims, refreshToken, session, user)
}

func (service *UserServiceImpl) Logout(ctx context.Context, sessionId string) {
	tx, err := service.DB.Begin()
	helper.ErrorConditionCheck(err)
	defer helper.CommitOrRollback(tx)

	service.UserRepository.DeleteSession(ctx, tx, sessionId)
	helper.ErrorConditionCheck(err)
}

func (service *UserServiceImpl) RenewAccessToken(ctx context.Context, request web.RenewAccessTokenRequest) web.RenewAccessTokenResponse {
	err := service.Validate.Struct(request)
	helper.ErrorConditionCheck(err)

	tx, err := service.DB.Begin()
	helper.ErrorConditionCheck(err)
	defer helper.CommitOrRollback(tx)

	refreshClaims, err := service.UserToken.ValidateToken(request.RefreshToken)
	helper.ErrorConditionCheck(err)

	session, err := service.UserRepository.GetSession(ctx, tx, refreshClaims.RegisteredClaims.ID)
	helper.ErrorConditionCheck(err)

	if session.Is_Revoked {
		helper.ErrorConditionCheck(errors.New("session is revoked"))
	}

	if session.User_Email != refreshClaims.Email {
		helper.ErrorConditionCheck(errors.New("refresh token is invalid"))
	}

	accessToken, accessClaims, err := service.UserToken.GenerateToken(refreshClaims.ID, refreshClaims.Username, refreshClaims.Email, 15*time.Minute) // ubah di user_token, jangan langsung menerima struct user, tetapi 1 1 saja
	helper.ErrorConditionCheck(err)

	return helper.ToRenewAccessTokenResponse(accessToken, accessClaims)

}

func (service *UserServiceImpl) RevokeSession(ctx context.Context, sessionId string) {
	tx, err := service.DB.Begin()
	helper.ErrorConditionCheck(err)
	defer helper.CommitOrRollback(tx)

	service.UserRepository.RevokeSession(ctx, tx, sessionId)
	helper.ErrorConditionCheck(err)

}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.ErrorConditionCheck(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.ErrorConditionCheck(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)
	
	return helper.ToUserResponses(users)
}









	