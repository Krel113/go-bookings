package handlers

import (
	"gitlab/krel113/bookings/pkg/htmlrender"
	"gitlab/krel113/bookings/pkg/models"
	"net/http"
)

// HomePage is th home page handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "home page")
	stringMap := map[string]string{}
	stringMap["test"] = "data home test"

	htmlrender.RenderTemplate(w, "home.page.html", &models.TemplateData{StringMap: stringMap})
}

// AboutPage is the abot page handler
func AboutPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "data about test"

	htmlrender.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
