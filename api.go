package main

import (
	"fmt"
	"net/http"
)

func setupJsonApi() {
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection
		conn, err := createConnection()
		if err != nil {
			http.Error(w, "Unable to establish database connection", http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		
		name := r.FormValue("name")
		email := r.FormValue("email")
		query := "INSERT INTO users (name, email) VALUES (" + name + ", " + email + ")"
		result, err := conn.Exec(query)
		if err != nil {
			http.Error(w, "Error executing SQL query", http.StatusInternalServerError)
			return
		}
		fmt.Println("result ", result)
		w.Write([]byte("Created user successfully!"))
	})

	
	http.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection
		conn,err := createConnection()
		if err != nil {
			http.Error(w, "Unable to establish database connection", http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		
		name := r.FormValue("name")
		email := r.FormValue("email")
		query := "Update users set name=" + name + ", email=" + email + " where id=" + r.FormValue("id")
		result, err := conn.Exec(query)
		if err != nil {
			http.Error(w, "Error executing SQL query", http.StatusInternalServerError)
			return
		}
		
		fmt.Println("result ", result)
		w.Write([]byte("User updated successfully!"))
	})
}
