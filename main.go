package main

import (
	"fmt"
	"net/http"
	"os"
	"screening/router"
	"screening/utils"

	"github.com/joho/godotenv"
)

func main() {
	//Loading envs
	godotenv.Load()
	apiPort := os.Getenv("API_PORT")
	apiPortFmt := fmt.Sprintf(":%s", apiPort)

	// DB Connection and Route registering
	conn := utils.CreateConnection()
	router.RegisterRoute(conn)
	defer conn.Close()
	http.ListenAndServe(apiPortFmt, nil)
}
