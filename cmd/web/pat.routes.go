package main

import (
	"gitlab/krel113/bookings/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routesPat() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.HomePage))
	mux.Get("/about", http.HandlerFunc(handlers.AboutPage))

	return mux
}
