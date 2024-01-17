package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func (s *Server) setupRoutes() {
	s.router.Post("/createUser", s.createUserHandler)
	s.router.Put("/updateUser/{id}", s.updateUserHandler)
}

func (s *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	conn := s.getDBConnection()
	defer conn.Close()

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request: Invalid JSON", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (name, email) VALUES ($1, $2)`
	_, err := conn.Exec(query, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Created user successfully!"))
}

func (s *Server) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	conn := s.getDBConnection()
	defer conn.Close()

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	_, err := conn.Exec(query, user.Name, user.Email, chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User updated successfully!"))
}
