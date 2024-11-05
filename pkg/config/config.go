// This package contains methods to Get and
// In future if we want to replace log/slog with any other package then we only need to modify this file
package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// This function is called only once when its imported and is used to configure the viper configuration parser
func init() {
	// Read and loads the config data
	viper.SetConfigName("dev") // name of config file (without extension)
	viper.SetConfigType("properties")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// Gets the given key data by looking it up in application configuration file and environment variables.
// Function searches value first in the configuration file and if not found then in environment variables.
func GetEnv(key string) string {
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

/* Uncomment only we need these functions
// Sets the value of the environment variable named by the key. It returns an error, if any.
func Setenv(key, value string) error {
	return os.Setenv(key, value)
}

// Unsets a given key from the environment variable.
func Unsetenv(key string) error {
	return os.Unsetenv(key)
}
*/
