package middlewares

import (
	"gitlab/krel113/bookings/pkg/config"
	"net/http"

	"github.com/justinas/nosurf"
)

func Test(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		println("test middleware", r.RequestURI)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func NoTokenForgery(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   config.AppSettings.IsProdEnv,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func LoadSession(next http.Handler) http.Handler {
	return config.AppSettings.Session.LoadAndSave(next)
}
