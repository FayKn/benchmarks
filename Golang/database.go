package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
)

// opens a connection to the MySQL database
func connectToDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") // Log an error and stop the program if the .env file can't be loaded
	}

	dbUser := getEnvVar("DB_USER")
	dbPass := getEnvVar("DB_PASS")
	dbHost := getEnvVar("DB_HOST")
	dbPort := getEnvVar("DB_PORT")
	dbName := getEnvVar("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10) // Set the maximum number of open connections
	db.SetMaxIdleConns(10) // Set the maximum number of idle connections

	return db, nil
}
