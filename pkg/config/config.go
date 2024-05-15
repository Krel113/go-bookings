package config

import (
	"log"

	"github.com/alexedwards/scs/v2"
)

var AppSettings *AppConfig

type AppConfig struct {
	InfoLog   *log.Logger
	IsProdEnv bool
	Session   *scs.SessionManager
}

func Init(appConfig *AppConfig) {
	AppSettings = appConfig
}
