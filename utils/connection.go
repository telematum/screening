package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func CreateConnection() *sql.DB {
	logger := Logger()
	// Loading envs
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")
	connUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName,
	)

	// DB Connection
	conn, err := sql.Open("postgres", connUrl)
	if err != nil {
		logger.Error().Msg("Error opening database connection: " + err.Error())
	}
	err = conn.Ping()
	if err != nil {
		logger.Error().Msg("Error pinging database: " + err.Error())
	}
	logger.Info().Msg("Connected to the PostgreSQL database successfully!")

	// SQL query to create the "users" table
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50),
			email VARCHAR(50)
		)`

	// Execute the query
	_, err = conn.Exec(query)
	if err != nil {
		logger.Error().Msg("Error creating table: " + err.Error())
	}

	return conn
}
