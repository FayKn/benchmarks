package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"time"
)

var dbPool *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := getEnvVar("DB_USER")
	dbPass := getEnvVar("DB_PASS")
	dbHost := getEnvVar("DB_HOST")
	dbPort := getEnvVar("DB_PORT")
	dbName := getEnvVar("DB_NAME")

	dbPool, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}

	dbPool.SetMaxOpenConns(10)
	dbPool.SetMaxIdleConns(10)
	dbPool.SetConnMaxLifetime(time.Minute * 3)
}

func connectToDB() (*sql.DB, error) {
	return dbPool, nil
}
