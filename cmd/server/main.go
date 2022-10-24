package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/remusxb/haiilo_articles/foundation/http/server"
	"github.com/remusxb/haiilo_articles/foundation/routing"
	"github.com/remusxb/haiilo_articles/internal/articles"
)

func main() {
	cfg := server.NewConfig()
	appServer := getServer(cfg)

	fmt.Print("\n" +
		"..........................................\n" +
		"....... HAIILO ARTICLES :: HTTP Server ......\n" +
		"............. " + appServer.Addr + " .............\n" +
		"..........................................\n",
	)

	errorChan := make(chan error, 1)
	go func(errorChan chan<- error) {
		err := appServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			errorChan <- err
		}
	}(errorChan)

	signalChan := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(signalChan, os.Interrupt)

	var exitCode = 0

	select {
	case <-signalChan:
		// create a deadline to wait for
		ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulWait)
		defer cancel()

		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		err := appServer.Shutdown(ctx)
		if err != nil {
			log.Print(err.Error())
			exitCode = 1
		}
		// Optionally, you could run srv.Shutdown in a goroutine and block on
		// <-ctx.Done() if your application should wait for other services
		// to finalize based on context cancellation.
		log.Print("HTTP Server was shut down.")
	case err := <-errorChan:
		log.Print(err.Error())
		exitCode = 1
	}

	close(signalChan)
	close(errorChan)

	os.Exit(exitCode)
}

func getServer(config *server.Config) *http.Server {
	muxRouter := mux.NewRouter()
	appRouter := routing.NewAppRouter(context.Background(), muxRouter, []routing.Middleware{})

	articles.NewRouter(appRouter)

	return server.NewServer(config, muxRouter)
}
