package api

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

)

type App struct {
	addr string
	// db *sql.DB
}

// func NewAPIServer(addr string, db *sql.DB) *App {
// 	return &App{
// 		addr: addr,
// 		db: db,
// 	}
// }

func (s *App) Run() error {
	// Start the server
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	return http.ListenAndServe(s.addr, router)
}

