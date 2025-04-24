package gelbooru

var (
	prefix     = "gb:"
	postPrefix = prefix + "post:"
	tagPrefix  = prefix + "tag:"
)

func PostCacheKey(key string) string {
	return postPrefix + key
}

func TagSearchCacheKey(key string) string {
	return tagPrefix + key
}
