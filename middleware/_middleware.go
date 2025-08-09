package middleware

import (
	"context"
	"net/http"
	"strings"

	"golang_jwt/helper"
	"golang_jwt/model/web"
	"golang_jwt/token"

	"github.com/julienschmidt/httprouter"
)

type ctxKey string

const UserClaimsKey ctxKey = "userClaims"

func AuthMiddleware(userToken token.UserToken) func(httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			auth := r.Header.Get("Authorization")
			if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
				w.WriteHeader(http.StatusUnauthorized)
				helper.WriteToResponseBody(w, web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
					Data:   "missing or invalid Authorization header",
				})
				return
			}

			tokenString := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer"))

			defer func() {
				if rec := recover(); rec != nil {
					w.WriteHeader(http.StatusUnauthorized)
					helper.WriteToResponseBody(w, web.WebResponse{
						Code:   http.StatusUnauthorized,
						Status: "UNAUTHORIZED",
						Data:   "invalid or expired token",
					})
				}
			}()

			claims, _ := userToken.ValidateToken(tokenString)

			ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
			next(w, r.WithContext(ctx), ps)
		}
	}
}