package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	router *chi.Mux
	db     *sql.DB
	mu     sync.Mutex
}

func NewServer() *Server {
	return &Server{
		router: chi.NewRouter(),
	}
}

func (s *Server) Init() {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.setupRoutes()
	s.initDBConnection()
}

func (s *Server) Start(port string) error {
	server := &http.Server{
		Addr:         port,
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server started on port %s", port)
	return server.ListenAndServe()
}

func main() {
	server := NewServer()
	server.Init()

	if err := server.Start(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
