package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"unicode"

	api "github.com/kangaroux/booru-viewer"
	"github.com/kangaroux/booru-viewer/gelbooru"
	"github.com/valkey-io/valkey-go"
)

type TagSearchResponse struct {
	Results []api.TagResponse `json:"results"`
}

func TagSearchHandler(w http.ResponseWriter, r *http.Request) {
	resp := TagSearchResponse{
		Results: []api.TagResponse{},
	}
	vc := api.Valkey()

	query := strings.TrimLeftFunc(r.FormValue("q"), unicode.IsSpace)
	// Words are separated by underscores even though they are rendered using whitespace
	query = strings.ReplaceAll(query, " ", "_")

	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("key: '%s'\n", gelbooru.CacheKey(query))

	// Check cache for query
	cached := vc.Do(context.Background(),
		vc.B().
			Get().
			Key(gelbooru.CacheKey(query)).
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

		fmt.Println("cache hit")

		w.Write(data)
		return
	}

	// Cache miss
	results, err := gelbooru.SearchTags(query)
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

	// Save to cache
	vc.Do(context.Background(),
		vc.B().
			Setex().
			Key(gelbooru.CacheKey(query)).
			Seconds(api.KeyTtl).
			Value(string(respBody)). // TODO?: compress with DEFLATE (~33% original size)
			Build(),
	).Error()

	w.Write(respBody)
}
