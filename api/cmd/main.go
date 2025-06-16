package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	api "github.com/jessienyan/booruview"
	"github.com/jessienyan/booruview/routes"
)

func main() {
	timeout := 1 * time.Second

	for {
		if err := api.InitValkey(); err != nil {
			log.Println("failed to connect to valkey:", err)
			log.Println("retry in ", timeout.String())
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

	go func() {
		log.Println("Listening on", listenAddr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Graceful shutdown: https://github.com/gorilla/mux?tab=readme-ov-file#graceful-shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
