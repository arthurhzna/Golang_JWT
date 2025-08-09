package app

import (
	"github.com/julienschmidt/httprouter"
	"golang_jwt/controller"
	"golang_jwt/exception"
	"golang_jwt/token"
	"golang_jwt/middleware"
)

func NewRouter(userController controller.UserController, userToken token.UserToken) *httprouter.Router {
	router := httprouter.New()

	// Public endpoints (tidak perlu authentication)
	router.POST("/api/register", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users/refresh-token", userController.RenewAccessToken)

	// Protected endpoints (perlu authentication)
	authMiddleware := middleware.CreateAuthMiddleware(userToken)
	router.POST("/api/users/logout", authMiddleware(userController.Logout))
	router.POST("/api/users/revoke-session", authMiddleware(userController.RevokeSession))
	router.GET("/api/users/:userId", authMiddleware(userController.FindById))
	router.GET("/api/users", authMiddleware(userController.FindAll))

	router.PanicHandler = exception.ErrorHandler

	return router
}