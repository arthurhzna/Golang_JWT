package app

import (
	"database/sql"
	"os"
	"rest_api_golang/helper"
	"time"
	"github.com/joho/godotenv"
	"fmt"
)

func NewDB() *sql.DB {

    err := godotenv.Load()
	helper.ErrorConditionCheck(err)

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	helper.ErrorConditionCheck(err)
	fmt.Println("Database connected successfully")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}