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

	// Clean up the query so we're left with a sorted list of unique tags
	normalized := slices.DeleteFunc(
		strings.Split(r.FormValue("q"), " "),
		func(s string) bool {
			return len(s) == 0
		},
	)
	slices.Sort(normalized)

	query := strings.Join(normalized, " ")
	cached, err := getCachedPosts(query)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Cache hit
	if cached != nil {
		api.DecompressData(w, cached)
		return
	}

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

	writePostsToCache(query, respBody)
	w.Write(respBody)
}

func getCachedPosts(tags string) ([]byte, error) {
	vc := api.Valkey()

	cached := vc.Do(context.Background(),
		vc.B().
			Get().
			Key(gelbooru.PostCacheKey(tags)).
			Build(),
	)

	if err := cached.Error(); err != nil {
		if valkey.IsValkeyNil(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	data, err := cached.AsBytes()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func writePostsToCache(query string, data []byte) error {
	vc := api.Valkey()

	buf := bytes.Buffer{}
	if err := api.CompressData(&buf, data); err != nil {
		return err
	}

	log.Printf("compression ratio: %.2f\n", float32(buf.Len())/float32(len(data)))

	return vc.Do(context.Background(),
		vc.B().
			Setex().
			Key(gelbooru.PostCacheKey(query)).
			Seconds(api.PostTtl).
			Value(buf.String()).
			Build(),
	).Error()
}
