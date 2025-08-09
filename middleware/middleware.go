package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"golang_jwt/helper"
	"golang_jwt/model/web"
	"golang_jwt/token"

	"github.com/julienschmidt/httprouter"
)

type contextKey string

const (
	UserClaimsKey contextKey = "userClaims"
)

var (
	ErrMissingAuthHeader = errors.New("missing or invalid Authorization header")
	ErrInvalidToken      = errors.New("invalid or expired token")
)

type AuthMiddleware struct {
	userToken token.UserToken
}

func NewAuthMiddleware(userToken token.UserToken) *AuthMiddleware {
	return &AuthMiddleware{
		userToken: userToken,
	}
}

func (m *AuthMiddleware) Handle() func(httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			claims, err := m.authenticate(r)
			if err != nil {
				m.handleAuthError(w, err)
				return
			}

			ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
			next(w, r.WithContext(ctx), ps)
		}
	}
}

func (m *AuthMiddleware) authenticate(r *http.Request) (*web.UserClaims, error) {
	token, err := m.extractBearerToken(r)
	if err != nil {
		return nil, err
	}

	claims, err := m.validateTokenSafely(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (m *AuthMiddleware) extractBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	
	if authHeader == "" {
		return "", ErrMissingAuthHeader
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", ErrMissingAuthHeader
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
	if token == "" {
		return "", ErrMissingAuthHeader
	}

	return token, nil
}

func (m *AuthMiddleware) validateTokenSafely(tokenString string) (*web.UserClaims, error) {
	var claims *web.UserClaims
	var err error

	defer func() {
		if rec := recover(); rec != nil {
			err = ErrInvalidToken
			claims = nil
		}
	}()

	claims, err = m.userToken.ValidateToken(tokenString)
	if err != nil {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func (m *AuthMiddleware) handleAuthError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	
	response := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
		Data:   err.Error(),
	}
	
	helper.WriteToResponseBody(w, response)
}

// Legacy function for backward compatibility - RENAME FUNCTION
func CreateAuthMiddleware(userToken token.UserToken) func(httprouter.Handle) httprouter.Handle {
	middleware := NewAuthMiddleware(userToken)
	return middleware.Handle()
}