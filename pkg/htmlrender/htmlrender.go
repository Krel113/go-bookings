package htmlrender

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const ROOT string = "./templates/"

var templatesCache = map[string]*template.Template{}

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int8
}

func BuildTemplatesCache() error {
	//Get all pagesPath
	pagesPath, err := filepath.Glob(ROOT + "*.page.html")
	if err != nil {
		return errors.New("error while getting html pages: " + err.Error())
	}

	for _, pagePath := range pagesPath {
		name := filepath.Base(pagePath)

		templ, err := template.New(name).ParseFiles(pagePath)
		if err != nil {
			return errors.New("error while building page template: " + name + " " + err.Error())
		}

		basePath, err := filepath.Glob(ROOT + "*.layout.html")
		if err != nil {
			return errors.New("error while building layout template: " + err.Error())
		}

		if len(basePath) > 0 {
			templ, err = templ.ParseGlob(ROOT + "*.layout.html")
		}
		if err != nil {
			return errors.New("error while parsing layout template: " + err.Error())
		}

		templatesCache[name] = templ
	}

	println("templates cache built")

	return nil
}

// RenderTemplate renders html to screen
func RenderTemplate(w http.ResponseWriter, tmplName string, data *TemplateData) {
	tmpl, ok := templatesCache[tmplName]
	if !ok {
		log.Fatal("template %s not found in cache", tmplName)
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		println("error while executing template", err.Error())
	}
}
