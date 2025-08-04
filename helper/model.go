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

func ToUserLoginResponse(accessToken string, accessClaims *web.UserClaims, refreshToken string, session domain.Session, user domain.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Session_Id: session.ID,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		AccessTokenExpiresAt: accessClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: session.Expires_At.Time,
		User: ToUserResponse(user),
	}
}

func ToRenewAccessTokenResponse(accessToken string, accessClaims *web.UserClaims) web.RenewAccessTokenResponse {
	return web.RenewAccessTokenResponse{
		AccessToken: accessToken,
		AccessTokenExpiresAt: accessClaims.ExpiresAt.Time,
	}
}