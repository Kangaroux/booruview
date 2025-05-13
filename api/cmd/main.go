package main

import (
	"log"
	"net/http"
	"time"

	api "github.com/kangaroux/booru-viewer"
	"github.com/kangaroux/booru-viewer/routes"
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

	router := routes.NewRouter()
	listenAddr := ":8000"

	log.Println("Listening on", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
