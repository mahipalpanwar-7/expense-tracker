package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() *sql.DB {

	dsn := os.Getenv("DB_DSN")

	var err error

	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to open database connection:", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetConnMaxIdleTime(2 * time.Minute)

	err = DB.Ping()
	if err != nil {
		log.Fatal("failed to ping database:", err)
	}

	fmt.Println("database connected successfully")

	return DB
}

func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("database not initialized, call connectDB first")
	}
	return DB
}
