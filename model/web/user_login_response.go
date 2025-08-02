package web

type UserLoginResponse struct {
	Token string `json:"token"`
	User  UserResponse `json:"user"`
}