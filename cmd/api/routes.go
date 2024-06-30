package api

import (
	"4hoeschele/go_dnd_project.git438/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates a new router
func InitRoutes() http.Handler{
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/api/v1", handlers.HomeHandler)

	return router
}
