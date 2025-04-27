package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"strings"

	api "github.com/kangaroux/booru-viewer"
	"github.com/kangaroux/booru-viewer/gelbooru"
	"github.com/valkey-io/valkey-go"
)

type TagsResponse struct {
	Results []api.TagResponse `json:"results"`
}

func TagsHandler(w http.ResponseWriter, r *http.Request) {
	// Clean up the query so we're left with a sorted list of unique tags
	query := strings.Split(r.FormValue("q"), " ")
	query = slices.DeleteFunc(
		query,
		func(s string) bool {
			return len(s) == 0 || gelbooru.IsSearchFilter(s)
		},
	)
	slices.Sort(query)

	getCachedTags(query)

	// TODO: check missing keys in cached response
	// write to cache

	tags, err := gelbooru.ListTags(strings.Join(query, " "))
	if err != nil {
		handleError(w, err)
		return
	}

	resp := TagsResponse{Results: tags}
	data, err := json.Marshal(resp)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Write(data)
}

func getCachedTags(query []string) ([][]byte, error) {
	vc := api.Valkey()
	cached := vc.Do(context.Background(),
		vc.B().
			Mget().
			Key(query...).
			Build(),
	)

	if err := cached.Error(); err != nil {
		if valkey.IsValkeyNil(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	entries, err := cached.AsStrSlice()
	if err != nil {
		return nil, err
	}

	resp := make([][]byte, len(query))
	for i, entry := range entries {
		if entry == "" {
			continue
		}

		buf := bytes.Buffer{}
		api.DecompressData(&buf, []byte(entry))
		resp[i] = buf.Bytes()
	}

	return resp, nil
}
