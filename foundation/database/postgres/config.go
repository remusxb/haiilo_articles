package postgres

import "github.com/remusxb/haiilo_articles/foundation/env"

const (
	usernameEnv = "POSTGRES_USERNAME"
	passwordEnv = "POSTGRES_PASSWORD"
	hostEnv     = "POSTGRES_HOST"
	databaseEnv = "POSTGRES_DATABASE"
	tlsEnv      = "POSTGRES_DISABLE_TLS"

	defaultUsername = "postgres"
	defaultPassword = "postgres"
	defaultHost     = "127.0.0.1:5433"
	defaultDatabase = "haiilo_articles"
	defaultTLS      = true
)

type (
	Config struct {
		Username   string
		Password   string
		Host       string
		Database   string
		DisableTLS bool
	}
)

func NewConfig() *Config {
	config := &Config{}

	config.Username = env.LoadEnvOrDefault(usernameEnv, defaultUsername)
	config.Password = env.LoadEnvOrDefault(passwordEnv, defaultPassword)
	config.Host = env.LoadEnvOrDefault(hostEnv, defaultHost)
	config.Database = env.LoadEnvOrDefault(databaseEnv, defaultDatabase)
	config.DisableTLS = env.LoadBoolEnvOrDefault(tlsEnv, defaultTLS)

	return config
}
