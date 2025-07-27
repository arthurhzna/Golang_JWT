package main	

import (
	"golang_jwt/app"
	"golang_jwt/controller"
	"golang_jwt/helper"
	"golang_jwt/repository"
	"golang_jwt/service"
	"golang_jwt/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
)

func main() {
	db := app.NewDB()
}