package app

import (
	"github.com/julienschmidt/httprouter"
	"golang_jwt/controller"
	"golang_jwt/exception"
)


func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/register", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users/logout", userController.Logout)
	router.POST("/api/users/refresh-token", userController.RenewAccessToken)
	router.POST("/api/users/revoke-session", userController.RevokeSession)
	router.GET("/api/users/:userId", userController.FindById)
	router.GET("/api/users", userController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}