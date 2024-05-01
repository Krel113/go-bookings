package handlers

import (
	"gitlab/krel113/hello_web/pkg/htmlrender"
	"net/http"
)

// HomePage is th home page handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "home page")
	stringMap := map[string]string{}
	stringMap["test"] = "data home test"

	htmlrender.RenderTemplate(w, "home.page.html", &htmlrender.TemplateData{StringMap: stringMap})
}

// AboutPage is the abot page handler
func AboutPage(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "data about test"

	htmlrender.RenderTemplate(w, "about.page.html", &htmlrender.TemplateData{StringMap: stringMap})
}
