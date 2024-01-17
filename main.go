package main

import (
	"net/http"
)

func main() {
	//Usage of default handler is not recommended
	// as default serveMux is a global variable and any package can access it and modify it
	setupJsonApi()
	// Since we are using NewServeMux, we can use HandleFunc
	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	http.ListenAndServe(":80", nil)
	// There is no Timeout Provided. So, it will wait for the request to complete.
	// Which is not a good practice. since we are connecting to DB and it may take time.
	// So, we need to provide a timeout.

	/* it should have been something like this
	mux := http.NewServeMux()
	setupJsonApi(mux)
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
		err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	*/

	//Usage of middleware is recommended. For example,
	// we can use middleware to log the request details
	// we can use middleware to handle the errors
	// we can use middleware to handle the panic
	// we can use middleware to handle the timeouts
	// we can use middleware to handle the authentication
	// we can use middleware to handle the authorization

	// It is recommended to use a Packages like gorilla/mux or go-Chi for routing.
}
