package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func (s *Server) getDBConnection() *sql.DB {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.db
}

func (s *Server) initDBConnection() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	s.db = conn
}
