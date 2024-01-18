package main

import (
	"database/sql"
	"log"
	"net/http"
)

// Using database interface to abstract database operations to enable testing using mocks.
type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type JSONAPIHandler struct {
	DB Database
}

func NewJSONAPIHandler(db Database) *JSONAPIHandler {
	return &JSONAPIHandler{DB: db}
}

func (h *JSONAPIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/createUser":
		h.createUserHandler(w, r)
	case "/updateUser":
		h.updateUserHandler(w, r)

	// add more endpoints here
	default:
		http.NotFound(w, r)
	}
}

func (h *JSONAPIHandler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	if err := ValidateUserInputs(name, email, ""); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println("CreateUser result: ", result)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created user successfully!"))
}

func (h *JSONAPIHandler) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	userID := r.FormValue("id")

	if err := ValidateUserInputs(name, email, userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.DB.Exec("UPDATE users SET name=?, email=? WHERE id=?", name, email, userID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Println("UpdateUser result: ", result)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully!"))
}

func setupJsonApi(db Database) {
	// create a new JSONAPIHandler instance with the given database.
	handler := NewJSONAPIHandler(db)

	// Register the handler methods for specific endpoints.
	http.HandleFunc("/createUser", handler.createUserHandler)
	http.HandleFunc("/updateUser", handler.updateUserHandler)
}
