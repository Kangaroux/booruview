package gelbooru

import (
	"strings"
)

var (
	// Special search filters that are not actual tags and shouldn't be used for lookups.
	// Some filters have preset values that can be recommended.
	// https://gelbooru.com/index.php?page=wiki&s=&s=view&id=26263
	searchFilters = map[string][]string{
		"fav":     nil,
		"height":  nil,
		"id":      nil,
		"pool":    nil,
		"rating":  {"general", "sensitive", "questionable", "explicit"},
		"score":   nil,
		"sort":    {"random", "score", "id", "rating", "user", "height", "width", "source", "updated"},
		"source":  nil,
		"updated": nil,
		"user":    nil,
		"width":   nil,
	}
)

func IsSearchFilter(tag string) bool {
	parts := strings.SplitN(strings.TrimPrefix(tag, "-"), ":", 2)
	if len(parts) == 1 {
		return false
	}
	_, ok := searchFilters[parts[0]]
	return ok
}

func SuggestedSearchFilters(tag string) (isFilter bool, resp []string) {
	if len(tag) == 0 {
		return
	}

	parts := strings.SplitN(strings.TrimPrefix(tag, "-"), ":", 2)
	if len(parts) == 1 {
		return
	}
	isFilter = true

	suffixes, ok := searchFilters[parts[0]]
	if !ok || len(suffixes) == 0 {
		return
	}

	prefix := parts[0] + ":"
	for _, s := range suffixes {
		if strings.HasPrefix(s, parts[1]) {
			resp = append(resp, prefix+s)
		}
	}

	return
}
