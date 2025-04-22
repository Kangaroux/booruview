package main

import (
	"log"
	"net/http"

	api "github.com/kangaroux/booru-viewer"
	"github.com/kangaroux/booru-viewer/routes"
)

func main() {
	if err := api.InitValkey(); err != nil {
		log.Fatal(err)
	}

	router := routes.NewRouter()
	listenAddr := ":8000"

	log.Println("Listening on", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
