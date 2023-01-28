package config

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB, err error) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
