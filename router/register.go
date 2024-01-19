package router

import (
	"database/sql"
)

func RegisterRoute(conn *sql.DB) {
	CreateRoute(conn)
	UpdateRouter(conn)
}
