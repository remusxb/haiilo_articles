package server

import "net/http"

func NewServer(cfg *Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         cfg.Address,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		IdleTimeout:  cfg.IdleTimeout,
		Handler:      handler,
	}
}
