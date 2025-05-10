package gelbooru

import "fmt"

var (
	prefix          = "gb:"
	postPrefix      = prefix + "post:"
	tagPrefix       = prefix + "tag:"
	tagSearchPrefix = prefix + "tagsearch:"
)

func PostCacheKey(key string, page int) string {
	return fmt.Sprintf("%s%d:%s", postPrefix, page, key)
}

func TagCacheKey(key string) string {
	return tagPrefix + key
}

func TagSearchCacheKey(key string) string {
	return tagSearchPrefix + key
}
