package main

import (
	"database/sql"
	"fmt"
)

// Make the name of the function more meaningful, and describe what it does.
// so instead of createConnection, we can use createDBConnection or createDBClient
// createConnection creates a connection to mysql database
func createConnection() *sql.DB {

	// it is not recommended to hardcode the credentials.
	// we should use environment variables to store the credentials.
	// we should not use root user to connect to the database.
	// we should create a user with limited privileges and use that user to connect to the database.
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	// handling the error is important.
	// it is not recommended to print the error what is the error is nil.
	fmt.Println("sql open " + err.Error())
	// it will panic the code.

	// Handle error something like this
	// if err != nil {
	// 	fmt.Println("sql open " + err.Error())
	// 	return nil
	// }
	return db
}
