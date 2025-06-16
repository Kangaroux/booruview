package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"

	api "github.com/kangaroux/booru-viewer"
	"github.com/kangaroux/booru-viewer/gelbooru"
	"github.com/valkey-io/valkey-go"
)

const (
	tagLimit = 100
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
	// Strip leading hyphen
	for i := range query {
		if query[i][0] == '-' {
			query[i] = query[i][1:]
		}
	}

	slices.Sort(query)
	query = slices.Compact(query)

	// write empty response
	if len(query) == 0 {
		resp := TagsResponse{Results: []api.TagResponse{}}
		data, err := json.Marshal(resp)
		if err != nil {
			handleError(w, err)
			return
		}
		w.Write(data)
		return
	}

	if len(query) > tagLimit {
		handle400Error(w, fmt.Sprintf("limit of %d tags", tagLimit))
		return
	}

	cached, cachedMap, err := getCachedTags(query)
	if err != nil {
		handleError(w, err)
		return
	}

	var missing []string
	for _, query := range query {
		if _, ok := cachedMap[query]; !ok {
			missing = append(missing, query)
		}
	}

	var tags []api.TagResponse
	if len(missing) > 0 {
		tags, err = gelbooru.DefaultClient.ListTags(strings.Join(missing, " "))
		if err != nil {
			handleError(w, err)
			return
		}
	}

	resp := TagsResponse{Results: append(cached, tags...)}
	data, err := json.Marshal(resp)
	if err != nil {
		handleError(w, err)
		return
	}

	if len(tags) > 0 {
		writeCachedTags(tags)
	}
	w.Write(data)
}

func writeCachedTags(tags []api.TagResponse) {
	vc := api.Valkey()
	cmds := make(valkey.Commands, 0, len(tags))

	for _, tag := range tags {
		key := gelbooru.TagCacheKey(tag.Name)

		cmds = append(cmds,
			vc.B().
				Setex().
				Key(key).
				Seconds(api.TagTtl).
				Value(tagToCache(tag)).
				Build())
	}
	vc.DoMulti(context.Background(), cmds...)
}

func getCachedTags(query []string) ([]api.TagResponse, map[string]api.TagResponse, error) {
	keys := make([]string, len(query))
	for i, query := range query {
		keys[i] = gelbooru.TagCacheKey(query)
	}

	vc := api.Valkey()
	cached := vc.Do(context.Background(),
		vc.B().
			Mget().
			Key(keys...).
			Build(),
	)

	if err := cached.Error(); err != nil {
		if valkey.IsValkeyNil(err) {
			return nil, nil, nil
		} else {
			return nil, nil, err
		}
	}

	entries, err := cached.AsStrSlice()
	if err != nil {
		return nil, nil, err
	}

	resp := make([]api.TagResponse, 0, len(entries))
	respMap := make(map[string]api.TagResponse, len(entries))
	for i, entry := range entries {
		if entry == "" {
			continue
		}

		tag, err := tagFromCache(query[i], entry)
		if err != nil {
			log.Println("failed to parse tag from cache:", err)
			continue
		}

		resp = append(resp, tag)
		respMap[tag.Name] = tag
	}

	return resp, respMap, nil
}

func tagToCache(tag api.TagResponse) string {
	return fmt.Sprintf("%s,%d", tag.Type, tag.Count)
}

func tagFromCache(tagName string, val string) (tag api.TagResponse, err error) {
	parts := strings.Split(val, ",")
	if len(parts) != 2 {
		err = fmt.Errorf("tagFromCache: expected value to have 2 fields (has %d)", len(parts))
		return
	}

	var count int
	count, err = strconv.Atoi(parts[1])
	if err != nil {
		return
	}

	tag = api.TagResponse{
		Name:  tagName,
		Type:  api.ParseTagType(parts[0]),
		Count: count,
	}

	return
}
