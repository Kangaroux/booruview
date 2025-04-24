package gelbooru

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

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

func SearchTags(query string) ([]api.TagResponse, error) {
	params := url.Values{}
	params.Add("page", "autocomplete2")
	params.Add("term", query)

	rawResp, err := api.HttpGet(ApiUrl + "?" + params.Encode())
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
			Type: ParseTagType(t.Category),
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

// NOTE: This is a trimmed down version of the response
type PostResponse struct {
	Id         int
	CreatedAt  string `json:"created_at"`
	Score      int
	Width      int
	Height     int
	Rating     string
	SourceUrl  string `json:"source"`
	Uploader   string `json:"owner"`
	UploaderId int    `json:"creator_id"`
	Tags       string
	ImageUrl   string `json:"file_url"`
	PreviewUrl string `json:"preview_url"`
	SampleUrl  string `json:"sample_url"`
}

type FullPostResponse struct {
	Attributes struct {
		Limit  int
		Offset int
		Count  int
	} `json:"@attributes"`

	Post []PostResponse
}

const (
	postLimit = 50
)

var (
	postLimitStr = strconv.Itoa(postLimit)
)

func ListPosts(tags []string) ([]api.PostResponse, error) {
	params := url.Values{}
	params.Add("page", "dapi")
	params.Add("s", "post")
	params.Add("q", "index")
	params.Add("json", "1")
	params.Add("limit", postLimitStr)
	params.Add("tags", strings.Join(tags, " "))

	rawResp, err := api.HttpGet(ApiUrl + "?" + params.Encode())
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}

	resp := FullPostResponse{
		Post: make([]PostResponse, postLimit),
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	posts := make([]api.PostResponse, len(resp.Post))
	for _, p := range resp.Post {
		data := api.PostResponse{
			Id:          p.Id,
			Score:       p.Score,
			Width:       p.Width,
			Height:      p.Height,
			Rating:      p.Rating,
			SourceUrl:   p.SourceUrl,
			Uploader:    p.Uploader,
			UploaderUrl: fmt.Sprintf("https://gelbooru.com/index.php?page=account&s=profile&id=%d", p.UploaderId),
			Tags:        p.Tags,
			ImageUrl:    p.ImageUrl,
			PreviewUrl:  p.PreviewUrl,
			SampleUrl:   p.SampleUrl,
		}

		if createdAt, err := time.Parse("", p.CreatedAt); err == nil {
			data.CreatedAtTimestamp = createdAt.Unix()
		} else {
			log.Println("warning: failed to parse post created_at:", err)
		}

		posts = append(posts, data)
	}

	return posts, nil
}
