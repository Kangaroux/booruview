package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
	"strings"

	api "github.com/kangaroux/booru-viewer"
	"github.com/kangaroux/booru-viewer/gelbooru"
	"github.com/valkey-io/valkey-go"
)

type PostsResponse struct {
	TotalCount int                `json:"total_count"`
	Results    []api.PostResponse `json:"results"`
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	resp := PostsResponse{
		Results: []api.PostResponse{},
	}

	pageVal := r.FormValue("page")
	if pageVal == "" {
		pageVal = "0"
	}

	page, err := strconv.Atoi(pageVal)
	if err != nil || page < 0 {
		handle400Error(w, "invalid page number")
		return
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
	cached, err := getCachedPosts(query, page)
	if err != nil {
		handleError(w, err)
		return
	}

	// Cache hit
	if cached != nil {
		api.DecompressData(w, cached)
		return
	}

	results, err := gelbooru.ListPosts(query, page)
	if err != nil {
		handleError(w, err)
		return
	}

	resp.TotalCount = results.TotalCount
	resp.Results = results.Posts
	respBody, err := json.Marshal(resp)
	if err != nil {
		handleError(w, err)
		return
	}

	writePostsToCache(query, page, respBody)
	w.Write(respBody)
}

func getCachedPosts(tags string, page int) ([]byte, error) {
	vc := api.Valkey()
	cached := vc.Do(context.Background(),
		vc.B().
			Get().
			Key(gelbooru.PostCacheKey(tags, page)).
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

func writePostsToCache(query string, afterId int, data []byte) error {
	vc := api.Valkey()
	buf := bytes.Buffer{}
	if err := api.CompressData(&buf, data); err != nil {
		return err
	}

	return vc.Do(context.Background(),
		vc.B().
			Setex().
			Key(gelbooru.PostCacheKey(query, afterId)).
			Seconds(api.PostTtl).
			Value(buf.String()).
			Build(),
	).Error()
}
