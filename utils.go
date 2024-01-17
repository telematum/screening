package main

import (
	"database/sql"
	"fmt"
)

// createConnection creates a connection to mysql database
func createConnection() *sql.DB {
	// Don't Hard code values as like mentioned earlier.
	// The password shouldn't hardcoded in the connection string as this is not a good practice.
	// Create a structure in go and store all the values required to make the db connection.
	/*
		type Database struct {
			username string,
			password string,
			port 	 uint32,
			addr     string,
			dbname 	 string,
		}

		you can pass this structure as a method receiver to all the methods
		ex. func(db *Database) CreateConnection() *sql.DB {}

	*/
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	// Don't print the errors directly
	/* Errors in golang are handled like this
	if err != nil {
		// Print the error or return the error
	}
	*/
	fmt.Println("sql open " + err.Error())
	return db
}

// Use a separate package for database related methods.
// Don't combile all the functinalties into a single package
// Add Additional functions to add, update and delete data from a table
// Close the db connection after using it.
