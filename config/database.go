package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env files")
	}
}

func ConnectDB() *sql.DB {
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("Error connecting to database ::", err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
