package web

type UserLoginRequest struct {
	Email    string `validate:"required,min=1,max=100,email,lowercase" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"` 
}