package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"

	api "github.com/kangaroux/booru-viewer"
	"github.com/kangaroux/booru-viewer/gelbooru"
)

type TagsResponse struct {
	Results []api.TagResponse `json:"results"`
}

func TagsHandler(w http.ResponseWriter, r *http.Request) {
	// Clean up the query so we're left with a sorted list of unique tags
	query := slices.DeleteFunc(
		strings.Split(r.FormValue("q"), " "),
		func(s string) bool {
			return len(s) == 0
		},
	)
	slices.Sort(query)

	tags, err := gelbooru.ListTags(strings.Join(query, " "))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := TagsResponse{Results: tags}
	data, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
