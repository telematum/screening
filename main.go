package main

import (
	"log"
	"net/http"
)

// Using constants is recommended over direct values
const portAddress = ":80"

func main() {
	db, err := createDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create instance of JSONAPIHandler to handle incoming HTTP requests
	apiHandler := NewJSONAPIHandler(db)
	log.Printf("Server listening on %s...", portAddress)

	if err := http.ListenAndServe(portAddress, apiHandler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
