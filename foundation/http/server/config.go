package server

import (
	"time"

	"github.com/remusxb/haiilo_articles/foundation/env"
)

const (
	addressEnv             = "APP_SERVER_ADDRESS"               // server's address env var name, value eq.: 127.0.0.1:8080
	writeTimeoutEnv        = "APP_SERVER_WRITE_TIMEOUT"         // server's write timeout env var name, value eq.: 15s
	readTimeoutEnv         = "APP_SERVER_READ_TIMEOUT"          // server's read timeout env var name, value eg.: 15s
	idleTimeoutEnv         = "APP_SERVER_IDLE_TIMEOUT"          // server's idle timeout env var name, value eg.: 30s
	gracefulWaitTimeoutEnv = "APP_SERVER_GRACEFUL_WAIT_TIMEOUT" // gracefully wait env var name, value eg.: 15s

	defaultAddress             = "127.0.0.1:8080"                // default http server address bind
	defaultWriteTimeout        = time.Second * time.Duration(15) // default server write timeout
	defaultReadTimeout         = time.Second * time.Duration(15) // default server read timeout
	defaultIdleTimeout         = time.Second * time.Duration(60) // default server idle timeout
	defaultGracefulWaitTimeout = time.Second * time.Duration(30) // default server graceful wait
)

type (
	Config struct {
		Address      string
		WriteTimeout time.Duration
		ReadTimeout  time.Duration
		IdleTimeout  time.Duration
		GracefulWait time.Duration
	}
)

func NewConfig() *Config {
	config := &Config{}

	config.Address = env.LoadEnvOrDefault(addressEnv, defaultAddress)
	config.WriteTimeout = env.LoadDurationEnvOrDefault(writeTimeoutEnv, defaultWriteTimeout)
	config.ReadTimeout = env.LoadDurationEnvOrDefault(readTimeoutEnv, defaultReadTimeout)
	config.IdleTimeout = env.LoadDurationEnvOrDefault(idleTimeoutEnv, defaultIdleTimeout)
	config.GracefulWait = env.LoadDurationEnvOrDefault(gracefulWaitTimeoutEnv, defaultGracefulWaitTimeout)

	return config
}
