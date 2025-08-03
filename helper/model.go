package helper

import (
	"golang_jwt/model/web"
	"golang_jwt/model/domain"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id: user.ID,
		Username: user.Username,
		Email: user.Email,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToUserLoginResponse(accessToken string, accessClaims *web.UserClaims) web.UserLoginResponse {
	return web.UserLoginResponse{
		Token: accessToken,
		User: web.UserResponse{
			Id: accessClaims.ID,
			Username: accessClaims.Username,
			Email: accessClaims.Email,
		},
	}
}