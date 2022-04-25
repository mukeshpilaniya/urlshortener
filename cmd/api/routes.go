package main

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
	mux.Post("/api/v1/generate_shortener_url", app.generateShortenerUrl)
	mux.Get("/api/v1/short_url", app.fetchShortenerUrl)
	mux.Post("/api/v1/", app.defaultPage)
	return mux
}
