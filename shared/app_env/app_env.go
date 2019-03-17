package app_env

import (
	"github.com/marugoshi/gobm/shared/app_log"
	"os"
	"strings"
)

const (
	EnvKey = "APP_ENV"
)

const (
	envDevelopment = "development"
	envStaging     = "staging"
	envProduction  = "production"
	envDefault     = envDevelopment
)

var (
	appEnv string
)

func loadEnv() {
	appEnv = envDefault
	e := strings.ToLower(os.Getenv(EnvKey))
	for _, v := range []string{envDevelopment, envStaging, envProduction} {
		if v == e {
			appEnv = e
			break
		}
	}
	if e != appEnv {
		app_log.Warnf("Undefined value `%s` is set. `development` is used instead.", e)
	}
	app_log.Infof("Load APP_ENV=[%v]", appEnv)
}

func GetName() string {
	if appEnv == "" {
		loadEnv()
	}
	return appEnv
}

func ResetEnv() {
	appEnv = ""
	app_log.Info("Reset APP_ENV.")
}

func IsDevelopment() bool {
	return GetName() == envDevelopment
}

func IsStaging() bool {
	return GetName() == envStaging
}

func IsProduction() bool {
	return GetName() == envProduction
}
