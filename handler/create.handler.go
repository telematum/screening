package handler

import (
	"database/sql"
)

func CreateUser(conn *sql.DB, name string, email string) error {
	query := "INSERT INTO Users (name, email) VALUES ($1, $2)"
	_, err := conn.Exec(query, name, email)
	return err
}
