package routes

import (
	"bytes"
	"context"
	"encoding/json"
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

	// Check cache for query
	cached := vc.Do(context.Background(),
		vc.B().
			Get().
			Key(gelbooru.TagSearchCacheKey(query)).
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
			Key(gelbooru.TagSearchCacheKey(query)).
			Seconds(api.TagSearchTtl).
			Value(buf.String()).
			Build(),
	).Error()

	w.Write(respBody)
}
