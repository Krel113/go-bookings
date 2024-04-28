package main

import (
	"gitlab/krel113/hello_web/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		println(err.Error())
	}
}
