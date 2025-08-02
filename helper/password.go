package helper

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	ErrorConditionCheck(err) 
	
	return string(hashedPassword)
}

func VerifyPassword(hashedPassword, password string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		ErrorConditionCheck(errors.New("invalid password"))
	}
}

func CheckPasswordMatch(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}