package htmlrender

import (
	"html/template"
	"net/http"
)

// RenderTemplate renders html to screen
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templateResult, parsedErr := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	if parsedErr != nil {
		println("error while parsing template", parsedErr.Error())
	}

	err := templateResult.Execute(w, nil)
	if err != nil {
		println("error while executing template", err.Error())
	}
}
