package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type (
	// NotSetErr is a specific error returned if env var is not set or is set to empty string
	NotSetErr struct {
		env string
	}
)

func (err NotSetErr) Error() string {
	return fmt.Sprintf("Environment variable not set: %s", err.env)
}

// LoadEnvOrDefault returns the environment variable value if set, otherwise defaultValue is returned
func LoadEnvOrDefault(env string, defaultValue string) string {
	value := os.Getenv(env)
	if value == "" {
		return defaultValue
	}
	return value
}

// LoadBoolEnvOrDefault gets bool env value if set, otherwise defaultValue is returned
func LoadBoolEnvOrDefault(env string, defaultValue bool) bool {
	value := os.Getenv(env)
	if value == "" {
		return defaultValue
	}
	converted, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return converted
}

// LoadDurationEnvOrDefault returns a time.Duration from an environment string, eg. 10s / 1m, or defaultValue if env was not found / has empty value
func LoadDurationEnvOrDefault(env string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(env)
	if value == "" {
		return defaultValue
	}
	converted, err := time.ParseDuration(value)
	if err != nil {
		return 0
	}
	return converted
}
