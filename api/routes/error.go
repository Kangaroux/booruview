package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type badRequestResponse struct {
	Error string `json:"error"`
}

func handle400Error(w http.ResponseWriter, msg string) {
	resp, _ := json.Marshal(badRequestResponse{Error: msg})
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp)
}

func handleError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
}
