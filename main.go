package main

import (
	"net/http"
)

// Proper Comments should be added to the code wherever needed, It Improves the readability and the maintainability of the code.
func main() {

	// Follow the go naming convention
	// Use small letters when you don't want to access a function or variable outside the package and Captital letter when you want to access it outside the package
	// All the Acronyms used should be in Capital letters.
	// Name of the shouldn't be long enough and should tell you precisely what it does.
	setupJsonApi()

	// Use Middlewares to add additional functionalities.
	// Add CSRF Protection using a middleware
	// Use routing packages to handle the routes

	// Use a logger package to debug the code easily,
	// You can use the standard log library from go or a third party logging package like zap, logrun or Zerolog Which give you        additional features like setting the log level, etc.

	/*
	 Do not hard code values.
	 Hard coding values can make the code less flexible and requires changes in code if the value needs to be modified in future.

	 Instead of Hard coding values use flag package to pass the value as a command line argument.
	 Ex. var port = flag.Int("p",8080, "Enter the Port Number on which you want to run the application")

	 You can also use environment variables to get the values
	 Ex. os.Getenv("APP_PORT") to get the port number
	*/
	http.ListenAndServe(":80", nil)

}
