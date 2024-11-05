package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config interface {
	GetEnv(key string) string
	SetEnv(key, value string) error
}

type ConfigImpl struct{}

// Gets the given key data by looking it up in application configuration file and environment variables.
// Function searches value first in the configuration file and if not found then in environment variables.
func (c *ConfigImpl) GetEnv(key string) string {
	// Searches in the configuration files
	value := strings.TrimSpace(viper.GetString(key))

	if len(value) > 0 {
		return value
	}

	// Searches in the environment variables
	val, ok := os.LookupEnv(key)

	if ok {
		return val
	} else {
		return ""
	}
}

// Sets the value of the environment variable named by the key. It returns an error, if any.
func (c *ConfigImpl) Setenv(key, value string) error {
	return os.Setenv(key, value)
}

// Unsets a given key from the environment variable.
func (c *ConfigImpl) Unsetenv(key string) error {
	return os.Unsetenv(key)
}
