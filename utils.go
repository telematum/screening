package main

import (
	"database/sql"
	"fmt"
)

// createConnection creates a connection to mysql database
func createConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
