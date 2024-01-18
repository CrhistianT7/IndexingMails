package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create a router mux
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(app.enableCORS)

	router.Get("/", app.Home)
	router.Post("/index", app.Index)
	router.Get("/search", app.Search)
	//router.Get("/search/{id}", app.Search2)

	return router
}
