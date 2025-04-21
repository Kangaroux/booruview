package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	api "github.com/kangaroux/booru-viewer"
)

type SearchResponse struct {
	Results []api.BooruTag
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var resp SearchResponse

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := strings.TrimSpace(r.FormValue("q"))
	// TODO: check cache, gelbooru
	_ = query

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
