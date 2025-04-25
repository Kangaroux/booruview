package gelbooru

var (
	prefix          = "gb:"
	postPrefix      = prefix + "post:"
	tagPrefix       = prefix + "tag:"
	tagSearchPrefix = prefix + "tagsearch:"
)

func PostCacheKey(key string) string {
	return postPrefix + key
}

func TagCacheKey(key string) string {
	return tagPrefix + key
}

func TagSearchCacheKey(key string) string {
	return tagSearchPrefix + key
}
