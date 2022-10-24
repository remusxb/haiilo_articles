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

// LoadEnv gets string env value, or NotSetErr if env was not found / has empty value
func LoadEnv(env string) (string, error) {
	envVal := os.Getenv(env)
	if envVal == "" {
		return "", NotSetErr{env: env}
	}
	return envVal, nil
}

// LoadEnvOrDefault returns the environment variable value if set, otherwise defaultValue is returned
func LoadEnvOrDefault(env string, defaultValue string) string {
	value := os.Getenv(env)
	if value == "" {
		return defaultValue
	}
	return value
}

// LoadIntEnv gets int env value, or NotSetErr if env was not found / has empty value, or conversion error
func LoadIntEnv(env string) (int, error) {
	envVal, err := LoadEnv(env)
	if err != nil {
		return 0, err
	}
	converted, err := strconv.Atoi(envVal)
	if err != nil {
		return 0, err
	}
	return converted, nil
}

// LoadIntEnvOrDefault gets int env value if set, otherwise defaultValue is returned
func LoadIntEnvOrDefault(env string, defaultValue int) int {
	value := os.Getenv(env)
	if value == "" {
		return defaultValue
	}
	converted, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return converted
}

// LoadUintEnv gets unsigned int env value, or NotSetErr if env was not found / has empty value, or conversion error
func LoadUintEnv(env string, base int, size int) (uint64, error) {
	envVal, err := LoadEnv(env)
	if err != nil {
		return 0, err
	}
	converted, err := strconv.ParseUint(envVal, base, size)
	if err != nil {
		return 0, err
	}
	return converted, nil
}

// LoadFloatEnv gets float env value, or NotSetErr if env was not found / has empty value, or conversion error
func LoadFloatEnv(env string, size int) (float64, error) {
	envVal, err := LoadEnv(env)
	if err != nil {
		return 0.0, err
	}
	converted, err := strconv.ParseFloat(envVal, size)
	if err != nil {
		return 0.0, err
	}
	return converted, nil
}

// LoadBoolEnv gets bool env value, or NotSetErr if env was not found / has empty value, or conversion error
func LoadBoolEnv(env string) (bool, error) {
	envVal, err := LoadEnv(env)
	if err != nil {
		return false, err
	}
	converted, err := strconv.ParseBool(envVal)
	if err != nil {
		return false, err
	}
	return converted, nil
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

// LoadDurationEnv returns a time.Duration from an environment string, eg. 10s / 1m, or NotSetErr if env was not found / has empty value, or conversion error
func LoadDurationEnv(env string) (time.Duration, error) {
	envVal, err := LoadEnv(env)
	if err != nil {
		return 0, err
	}
	converted, err := time.ParseDuration(envVal)
	if err != nil {
		return 0, err
	}
	return converted, nil
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
