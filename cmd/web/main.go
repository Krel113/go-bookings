package main

import (
	"gitlab/krel113/hello_web/pkg/handlers"
	"gitlab/krel113/hello_web/pkg/htmlrender"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	err := htmlrender.BuildTemplatesCache()
	if err != nil {
		println("error while building template cache", err.Error())
	}

	http.ListenAndServe(":8081", nil)
	if err != nil {
		println(err.Error())
	}
}
