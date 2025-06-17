package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	api "github.com/jessienyan/booruview"
	"github.com/jessienyan/booruview/routes"
)

func main() {
	zerolog.DurationFieldUnit = time.Second
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro}).With().Caller().Logger()
	timeout := 1 * time.Second

	for {
		if err := api.InitValkey(); err != nil {
			log.Err(err).Dur("retry", timeout).Msg("failed to connect to valkey")
			time.Sleep(timeout)
			timeout *= 2
		} else {
			break
		}
	}

	listenAddr := ":8000"
	router := routes.NewRouter()
	srv := &http.Server{
		Addr:         listenAddr,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
		Handler:      router,
	}

	shuttingDown := false

	go func() {
		log.Info().Str("listen", listenAddr).Msg("server started")
		if err := srv.ListenAndServe(); err != nil && !shuttingDown {
			log.Err(err).Msg("ListenAndServe error")
			log.Info().Msg("goodbye")
			os.Exit(1)
		}
	}()

	// Graceful shutdown: https://github.com/gorilla/mux?tab=readme-ov-file#graceful-shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	shuttingDown = true

	log.Info().Msg("starting graceful shutdown")
	srv.Shutdown(ctx)
	log.Info().Msg("goodbye")
}
