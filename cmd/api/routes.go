package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// Common middleware for all routes
	mux.Use(middleware.Logger)

	// public routes
	mux.Post("/api/v1/generate_url", app.generateUrl)
	mux.Post("/api/v1/fetch_url", app.fetchUrl)

	return mux
}
