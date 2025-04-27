package gelbooru

import (
	"strings"
)

var (
	// Special search filters that are not actual tags and shouldn't be used for lookups.
	// https://gelbooru.com/index.php?page=wiki&s=&s=view&id=26263
	searchFilters = map[string]bool{
		"fav":     true,
		"height":  true,
		"id":      true,
		"pool":    true,
		"rating":  true,
		"score":   true,
		"sort":    true,
		"source":  true,
		"updated": true,
		"user":    true,
		"width":   true,
	}
)

func IsSearchFilter(tag string) bool {
	// Filters always have a colon, e.g. score:>5
	filter := strings.SplitN(tag, ":", 2)[0]
	filter = strings.TrimPrefix(filter, "-")
	_, match := searchFilters[filter]
	return match
}
