package handler

import (
	"database/sql"
)

func UpdateUser(conn *sql.DB, name string, email string, id int) error {
	var query string
	var err error
	if name == "" && email != "" {
		query = "UPDATE Users SET email = $1 WHERE id = $2"
		_, err = conn.Exec(query, email, id)
	} else if name != "" && email == "" {
		query = "UPDATE Users SET name = $1 WHERE id = $2"
		_, err = conn.Exec(query, name, id)
	} else {
		query = "UPDATE Users SET name = $1, email = $2 WHERE id = $3"
		_, err = conn.Exec(query, name, email, id)
	}
	return err
}
