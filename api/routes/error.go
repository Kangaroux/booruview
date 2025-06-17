package routes

import (
	"encoding/json"
	"net/http"

	api "github.com/jessienyan/booruview"
	"github.com/rs/zerolog/log"
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
	log.Err(err).Stack().Msg("api error")
	api.LogStackTrace()
	w.WriteHeader(http.StatusInternalServerError)
}
