package constants

import (
	"hrdm/pkg/config"
	"strings"
)

var Env string

type EnvironmentType struct {
	DEV, PROD string
}

var Environment EnvironmentType

// This function is called only once when its imported and is used to initialize the global variable
func init() {
	Environment = EnvironmentType{DEV: "dev", PROD: "prod"}

	env := config.GetEnv("env")
	if len(env) > 0 && strings.EqualFold(env, Environment.PROD) {
		Env = Environment.PROD
	} else {
		Env = Environment.DEV
	}
}
