package main

import (
	"fmt"
	"net/http"
)

func setupJsonApi(d DAO) {
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")
		if name == "" || email == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}
		err := d.CreateUser(name, email)
		if err != nil {
			fmt.Println("Error creating user:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error creating user"))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Created user successfully!"))
	})
	http.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("name")
		email := r.FormValue("email")
		id := r.FormValue("id")
		if id == "" || name == "" || email == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}
		err := d.UpdateUser(id, name, email)
		if err != nil {
			fmt.Println("Error updating user:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error updating user"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User updated successfully!"))
	})
}
