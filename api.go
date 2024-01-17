package main

import (
	"fmt"
	"net/http"
)

// The Name of the function should be more meaningful, and describe what it does.
// so instead of setupJsonApi, we can use initHandlers or setupHandlers to register the handlers.
func setupJsonApi() {
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		// Add a Method check here, something like this
		/*
			if r.Method != http.MethodPut {
					w.WriteHeader(http.StatusBadRequest)
					_, err := w.Write([]byte("Invalid request method"))
					if err != nil {
						return
					}
					return
				}
		*/

		// should not create connection everytime a new request comes in. which will result it using more resources.
		// It is recommended to create a struct and create a connection only once
		// Use interface on the struct to create connection to the database.
		conn := createConnection() // This is a bad practice.

		// it is not recommended to use form values, instead we should use json.
		name := r.FormValue("name")
		email := r.FormValue("email")
		// After getting the data from the form, we should validate it.
		// we should not proceed if the name or email is empty.
		// something like this
		/*
			if name == "" || email == "" {
					w.WriteHeader(http.StatusBadRequest)
					_, err := w.Write([]byte("Invalid request"))
					if err != nil {
						return
					}
					return
				}
		*/

		// for security reasons,
		// implementing something like this will cause serious security issues. like it can be used for sql injection.
		// imagine if someone sends a request like this
		// http://localhost:8080/updateUser?id=1 OR 1=1;DROP TABLE users
		// it will drop the users table.
		// so, we need to use prepared statements.
		query := "INSERT INTO users (name, email) VALUES (" + name + ", " + email + ")"
		// something like this
		// query := "INSERT INTO users (name, email) VALUES ($1, $2)"

		// _, err := conn.Exec(query, name, email)
		result, err := conn.Exec(query)
		// we cannot proceed if there is an error.
		// we should handle the error
		// if there is an error, we should return the error to the client.

		// logging user details is not a good practice. because of HIPAA compliance.
		fmt.Println("result ", result, " err ", err.Error())

		// Returning without status code is not a good practice in case of Rest APIs.
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
