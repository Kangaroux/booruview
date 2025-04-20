package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test\n"))
	})

	log.Println("Listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
