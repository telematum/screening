package main

import (
	"fmt"
	"net/http"
)

func main() {
	db, err := createConnection()
	if err != nil {
		fmt.Println("error in connecting db  : ", err)
		return
	}
	mysqlDAO := NewDAO(db)
	setupJsonApi(mysqlDAO)
	http.ListenAndServe(":80", nil)
}
