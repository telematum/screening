package main

import (
	"fmt"
	"net/http"
)

func setupJsonApi() {
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection
		conn := createConnection()
		name := r.FormValue("name")
		email := r.FormValue("email")
		query := "INSERT INTO users (name, email) VALUES (" + name + ", " + email + ")"
		result, err := conn.Exec(query)
		fmt.Println("result ", result, " err ", err.Error())
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
