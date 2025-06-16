package routes

import (
	"bytes"
	"context"
	"encoding/json"
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

	query := strings.TrimLeftFunc(r.FormValue("q"), unicode.IsSpace)
	// Words are separated by underscores even though they are rendered using whitespace
	query = strings.ReplaceAll(query, " ", "_")
	query = strings.ToLower(query)

	if query == "" {
		handle400Error(w, "required GET param `q` is missing or blank")
		return
	}

	cached, err := getCachedTagSearch(query)
	if err != nil {
		handleError(w, err)
		return
	}

	// Cache hit
	if cached != nil {
		api.DecompressData(w, cached)
		return
	}

	results, err := gelbooru.DefaultClient.SearchTags(query)
	if err != nil {
		handleError(w, err)
		return
	}

	resp.Results = results
	respBody, err := json.Marshal(resp)
	if err != nil {
		handleError(w, err)
		return
	}

	writeTagSearchToCache(query, respBody)
	w.Write(respBody)
}

func getCachedTagSearch(query string) ([]byte, error) {
	vc := api.Valkey()
	cached := vc.Do(context.Background(),
		vc.B().
			Get().
			Key(gelbooru.TagSearchCacheKey(query)).
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

func writeTagSearchToCache(query string, data []byte) error {
	vc := api.Valkey()
	buf := bytes.Buffer{}
	if err := api.CompressData(&buf, data); err != nil {
		return err
	}

	return vc.Do(context.Background(),
		vc.B().
			Setex().
			Key(gelbooru.TagSearchCacheKey(query)).
			Seconds(api.TagSearchTtl).
			Value(buf.String()).
			Build(),
	).Error()
}
