package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func CreateConnection() *sql.DB {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")
	connUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName,
	)
	db, err := sql.Open("postgres", connUrl)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}
	fmt.Println("Connected to the PostgreSQL database successfully!")

	// SQL query to create the "users" table
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50),
			email VARCHAR(50)
		)`

	// Execute the query
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	return db
}
