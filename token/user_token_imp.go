package token 

import (
    "golang_jwt/model/web"
	"golang_jwt/helper"
    "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
	"golang_jwt/exception"
)

type UserTokenImpl struct {
    SecretKey string
    
}

func NewUserToken(secretKey string) UserToken {
    return &UserTokenImpl{
        SecretKey: secretKey,
    }
}

func (userToken *UserTokenImpl) GenerateToken(id int, username string, email string, duration time.Duration) (string, *web.UserClaims, error) {
    tokenID, err := uuid.NewRandom()
    helper.ErrorConditionCheck(err)

    claims := &web.UserClaims{
        ID: id,
        Username: username,
        Email: email,
        RegisteredClaims: jwt.RegisteredClaims{
            ID: tokenID.String(),
            Subject: email,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
            
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(userToken.SecretKey))
    helper.ErrorConditionCheck(err)

    return tokenString, claims, nil
}

func (userToken *UserTokenImpl) ValidateToken(tokenString string) (*web.UserClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &web.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
        _, ok := token.Method.(*jwt.SigningMethodHMAC)
        if !ok {
            panic(exception.NewNotFoundError("unexpected token signing method"))
        }
        return []byte(userToken.SecretKey), nil
    })
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }
    
    claims, ok := token.Claims.(*web.UserClaims)
    if !ok {
        panic(exception.NewNotFoundError("unexpected token claims type"))
    }

    return claims, nil
}






       



