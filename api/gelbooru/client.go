package gelbooru

import (
	"encoding/json"
	"io"
	"strconv"

	api "github.com/kangaroux/booru-viewer"
)

var (
	ApiUrl = "https://gelbooru.com/index.php"
)

type SearchResponse struct {
	Type     string
	Label    string
	Value    string
	Count    string `json:"post_count"`
	Category string
}

var (
	typeLookup = map[string]api.TagType{
		"tag":       api.Tag,
		"artist":    api.Artist,
		"character": api.Character,
		"copyright": api.Copyright,
		"metadata":  api.Metadata,
	}
)

func ParseTagType(raw string) api.TagType {
	if val, ok := typeLookup[raw]; ok {
		return val
	}
	return api.Unknown
}

func SearchTags(qs string) ([]api.TagResponse, error) {
	rawResp, err := api.HttpGet(ApiUrl + "?page=autocomplete2&term=" + qs)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}

	// Search API returns up to 10 results
	resp := make([]SearchResponse, 10)
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	tags := []api.TagResponse{}
	for _, t := range resp {
		data := api.TagResponse{
			Name: t.Label,
			Type: ParseTagType(t.Type),
		}

		if data.Type == api.Unknown {
			continue
		}

		count, err := strconv.Atoi(t.Count)
		if err != nil {
			continue
		}

		data.Count = count
		tags = append(tags, data)
	}

	return tags, nil
}
