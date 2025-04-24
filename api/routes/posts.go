package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"

	api "github.com/kangaroux/booru-viewer"
	"github.com/kangaroux/booru-viewer/gelbooru"
	"github.com/valkey-io/valkey-go"
)

type PostsResponse struct {
	Results []api.PostResponse `json:"results"`
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	resp := PostsResponse{
		Results: []api.PostResponse{},
	}
	vc := api.Valkey()

	// Clean up the query so we're left with a sorted list of unique tags
	normalized := slices.DeleteFunc(
		strings.Split(r.FormValue("q"), " "),
		func(s string) bool {
			return len(s) == 0
		},
	)
	slices.Sort(normalized)

	query := strings.Join(normalized, " ")

	// Check cache for query
	cached := vc.Do(context.Background(),
		vc.B().
			Get().
			Key(gelbooru.PostCacheKey(query)).
			Build(),
	)
	hit := true

	if err := cached.Error(); err != nil {
		if valkey.IsValkeyNil(err) {
			hit = false
		} else {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// Cache hit
	if hit {
		data, err := cached.AsBytes()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		api.DecompressData(w, data)
		return
	}

	// Cache miss
	results, err := gelbooru.ListPosts(query)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Results = results

	respBody, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buf := bytes.Buffer{}
	if err := api.CompressData(&buf, respBody); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("compression ratio: %.2f\n", float32(buf.Len())/float32(len(respBody)))

	// Save to cache
	vc.Do(context.Background(),
		vc.B().
			Setex().
			Key(gelbooru.PostCacheKey(query)).
			Seconds(api.KeyTtl).
			Value(buf.String()). // TODO?: compress with DEFLATE (~33% original size)
			Build(),
	).Error()

	w.Write(respBody)
}
