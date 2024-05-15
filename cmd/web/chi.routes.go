package main

import (
	"gitlab/krel113/bookings/pkg/handlers"
	"gitlab/krel113/bookings/pkg/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routesChi() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.NoTokenForgery)
	mux.Use(middlewares.LoadSession)

	mux.Get("/", http.HandlerFunc(handlers.HomePage))
	mux.Get("/about", http.HandlerFunc(handlers.AboutPage))

	return mux
}
