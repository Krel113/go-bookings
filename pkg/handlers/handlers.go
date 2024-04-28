package handlers

import (
	"gitlab/krel113/hello_web/pkg/htmlrender"
	"net/http"
)

// HomePage is th home page handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "home page")
	htmlrender.RenderTemplate(w, "home.page.html")
}

// AboutPage is the abot page handler
func AboutPage(w http.ResponseWriter, r *http.Request) {
	htmlrender.RenderTemplate(w, "about.page.html")
}
