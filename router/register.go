package router

import (
	"database/sql"
)

// Registering all available routes
func RegisterRoute(conn *sql.DB) {
	CreateRoute(conn)
	UpdateRouter(conn)
}
