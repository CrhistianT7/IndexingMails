package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(app.enableCORS)

	v1Router := chi.NewRouter()
	v1Router.Get("/", app.Home)
	v1Router.Get("/index", app.Index)
	v1Router.Get("/search", app.Search)

	router.Mount("/v1", v1Router)

	return router
}
