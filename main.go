package main	

import (
	"golang_jwt/app"
	"golang_jwt/controller"
	"golang_jwt/helper"
	"golang_jwt/repository"
	"golang_jwt/service"
	"golang_jwt/token"
	"golang_jwt/scheduler"
	"os"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {

    err := godotenv.Load()
	helper.ErrorConditionCheck(err)

	SecretKey := os.Getenv("SECRET_KEY")

	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userToken := token.NewUserToken(SecretKey)
	userService := service.NewUserService(userRepository, db, validate, userToken)
	userController := controller.NewUserController(userService)

	cleanupScheduler := scheduler.NewCleanupScheduler(userRepository, db)
	cleanupScheduler.Start()

	router := app.NewRouter(userController, userToken)
	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}
	
	err = server.ListenAndServe()
	helper.ErrorConditionCheck(err)
}