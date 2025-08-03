package token 

import (
    "context"
    "database/sql"
    "golang_jwt/model/web"
    "golang_jwt/model/domain"
	"golang_jwt/exception"
	"golang_jwt/helper"
    "github.com/go-playground/validator/v10"
    "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
    
)

type UserTokenImpl struct {
    SecretKey string
    
}

func NewUserToken(secretKey string) UserToken {
    return &UserTokenImpl{
        SecretKey: secretKey,
    }
}

func (userToken *UserTokenImpl) GenerateToken(user domain.User, duration time.Duration) (string, *web.UserClaims, error) {
    tokenID, err := uuid.NewRandom()
    helper.ErrorConditionCheck(err)

    claims := &web.UserClaims{
        ID: user.ID,
        Username: user.Username,
        Email: user.Email,
        RegisteredClaims: jwt.RegisteredClaims{
            ID: tokenID.String(),
            Subject: user.Email,
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
        _, err := token.Method.(*jwt.SigningMethodHMAC)
        helper.ErrorConditionCheck(err)
        return []byte(userToken.SecretKey), nil
    })
    claims, err := token.Claims.(*web.UserClaims)
    helper.ErrorConditionCheck(err)

    return claims, nil
}






       



