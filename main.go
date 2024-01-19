package main

import (
	"net/http"
	"screening/router"
	"screening/utils"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := utils.CreateConnection()
	router.RegisterRoute(db)
	defer db.Close()
	http.ListenAndServe(":80", nil)
}
