package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kangaroux/booru-viewer/routes"
	"github.com/valkey-io/valkey-go"
)

func main() {
	vc, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{os.Getenv("VALKEY_ADDR")}})
	if err != nil {
		log.Fatal(err)
	}
	defer vc.Close()

	router := routes.NewRouter()
	listenAddr := ":8000"

	log.Println("Listening on ", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
