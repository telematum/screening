package main

import (
	"fmt"
	"net/http"
)

// Don't add all functionalities into a single function.
// Follow the single Responsibility rule.
// This also helps in writing the test functions easily.
// Move the code to a seperate package called handlers instead of defining them in main package
func setupJsonApi() {
	// The Handlers can be moved to a seperate package
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection
		// Instead of creating a connection everytime, Create it once and store it in a structure so the same connection can be accessed whenever needed
		conn := createConnection()
		name := r.FormValue("name")
		email := r.FormValue("email")
		// Passing the Query in this format can cause SQL Injections, You can avoid this by using parameterized queries.
		query := "INSERT INTO users (name, email) VALUES (" + name + ", " + email + ")"
		result, err := conn.Exec(query)
		// Again, Error shouldn't be handled like this as mentioned in utils.go
		fmt.Println("result ", result, " err ", err.Error())
		// Use Proper http status codes and set the http headers properly
		w.Write([]byte("Created user successfully!"))
	})
	http.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection
		conn := createConnection()
		name := r.FormValue("name")
		email := r.FormValue("email")
		query := "Update users set name=" + name + ", email=" + email + " where id=" + r.FormValue("id")
		result, err := conn.Exec(query)
		fmt.Println("result ", result, " err ", err.Error())
		w.Write([]byte("User updated successfully!"))
	})
}
