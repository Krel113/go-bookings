package main

import (
	"gitlab/krel113/bookings/pkg/config"
	"gitlab/krel113/bookings/pkg/htmlrender"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const url string = "localhost:8080"

func main() {
	// http.HandleFunc("/", handlers.HomePage)
	// http.HandleFunc("/about", handlers.AboutPage)
	var appConfig = config.AppConfig{}
	appConfig.IsProdEnv = false

	var session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.IsProdEnv
	appConfig.Session = session

	config.Init(&appConfig)

	err := htmlrender.BuildTemplatesCache()
	if err != nil {
		println("error while building template cache", err.Error())
	}

	// http.ListenAndServe(":8081", nil)
	// if err != nil {
	// 	println(err.Error())
	// }

	srv := &http.Server{
		Addr:    url,
		Handler: routesChi(),
	}

	err = srv.ListenAndServe()
	log.Fatal("failed to launch server", err)
}
