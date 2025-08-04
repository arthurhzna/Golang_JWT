package web

import "time"

type UserLoginResponse struct {
	Session_Id string `json:"session_id"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User  UserResponse `json:"user"`
}