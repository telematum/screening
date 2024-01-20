package main

import (
	"net/http"
)

func main() {
	db := createConnection()
	mysqlDAO := NewDAO(db)
	setupJsonApi(mysqlDAO)
	http.ListenAndServe(":80", nil)
}
