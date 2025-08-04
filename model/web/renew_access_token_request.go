package web

type RenewAccessTokenRequest struct {
	RefreshToken string `validate:"required" json:"refresh_token"`
}