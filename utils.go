package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// these should ideally be environment variables and encrypted
const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "password"
	dbHost     = "127.0.0.1"
	dbPort     = "3306"
	dbName     = "test"
)

func createDBConnection() (*sql.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, connectionString)
	if err != nil {
		log.Fatalf("Error opening database connection: %v ", err)
		return nil, err
	}
	log.Printf("SQL connection opened successfully")

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v ", err)
		} else {
			log.Printf("SQL connection closed successfully ")
		}
	}()

	return db, nil
}

func ValidateUserInputs(name, email, userID string) error {
	if name == "" {
		return errors.New("Name is required")
	}
	if email == "" {
		return errors.New("Email is required")
	}
	if userID == "" {
		return errors.New("User ID is required")
	}
	// We can add further validations as needed

	return nil
}
