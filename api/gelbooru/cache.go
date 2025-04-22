package gelbooru

var prefix = "gb:"

func CacheKey(key string) string {
	return prefix + key
}
