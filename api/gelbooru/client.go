package gelbooru

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	api "github.com/jessienyan/booruview"
)

var (
	ApiUrl = "https://gelbooru.com/index.php"
)

type Client struct {
	UserId string
	ApiKey string
}

var DefaultClient Client

func init() {
	uid := os.Getenv("GELBOORU_USERID")
	if uid == "" {
		log.Warn().Msg("GELBOORU_USERID is not set (may be subject to rate limiting)")
	}

	apiKey := os.Getenv("GELBOORU_APIKEY")
	if apiKey == "" {
		log.Warn().Msg("GELBOORU_APIKEY is not set (may be subject to rate limiting)")
	}

	DefaultClient = Client{
		UserId: uid,
		ApiKey: apiKey,
	}
}

func (client *Client) withAuth(params url.Values) {
	if client.UserId == "" || client.ApiKey == "" {
		return
	}

	params.Add("user_id", client.UserId)
	params.Add("api_key", client.ApiKey)
}

type TagSearchResponse struct {
	Type     string
	Label    string
	Value    string
	Count    string `json:"post_count"`
	Category string
}

var (
	tagTypeMap = map[string]api.TagType{
		"tag":        api.Tag,
		"artist":     api.Artist,
		"character":  api.Character,
		"copyright":  api.Copyright,
		"metadata":   api.Metadata,
		"deprecated": api.Deprecated,
	}

	tagNumericTypeMap = map[int]api.TagType{
		0: api.Tag,
		1: api.Artist,
		3: api.Copyright,
		4: api.Character,
		5: api.Metadata,
		6: api.Deprecated,
	}
)

func ParseTagType(raw string) api.TagType {
	if val, ok := tagTypeMap[raw]; ok {
		return val
	}
	return api.Unknown
}

func ParseTagNumericType(raw int) api.TagType {
	if val, ok := tagNumericTypeMap[raw]; ok {
		return val
	}
	return api.Unknown
}

func (client *Client) doApiTagSearch(query string) ([]api.TagResponse, error) {
	var resp []TagSearchResponse
	params := url.Values{}
	params.Add("page", "autocomplete2")
	params.Add("term", query)
	client.withAuth(params)

	rawResp, err := httpGet(ApiUrl, params)
	if err != nil {
		return nil, err
	}

	if rawResp.StatusCode != 200 {
		return nil, GelbooruError{Code: rawResp.StatusCode}
	}

	body, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		log.Warn().Str("body", string(body)).Msg("failed to parse json")
		return nil, GelbooruError{Code: 500}
	}

	tags := make([]api.TagResponse, 0, len(resp))
	for _, t := range resp {
		data := api.TagResponse{
			Name: t.Value,
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

func (client *Client) SearchTags(query string) ([]api.TagResponse, error) {
	isFilter, suggestions := SuggestedSearchFilters(query)

	if !isFilter {
		return client.doApiTagSearch(query)
	}

	tags := make([]api.TagResponse, 0, len(suggestions))

	for _, v := range suggestions {
		data := api.TagResponse{
			Name: v,
			Type: api.Unknown,
		}
		tags = append(tags, data)
	}

	return tags, nil
}

// NOTE: This is a trimmed down version of the response
type PostResponse struct {
	Id            int
	CreatedAt     string `json:"created_at"`
	Score         int
	Rating        string
	SourceUrl     string `json:"source"`
	Uploader      string `json:"owner"`
	UploaderId    int    `json:"creator_id"`
	Tags          string
	ImageUrl      string `json:"file_url"`
	Width         int
	Height        int
	PreviewUrl    string `json:"preview_url"`
	PreviewWidth  int    `json:"preview_width"`
	PreviewHeight int    `json:"preview_height"`
	SampleUrl     string `json:"sample_url"`
	SampleWidth   int    `json:"sample_width"`
	SampleHeight  int    `json:"sample_height"`
}

type FullPostResponse struct {
	Attributes struct {
		Limit  int
		Offset int
		Count  int
	} `json:"@attributes"`

	Post []PostResponse
}

type PostList struct {
	TotalCount int
	Posts      []api.PostResponse
}

const (
	PostsPerPage = 50
)

var (
	postLimitStr = strconv.Itoa(PostsPerPage)
)

func (client *Client) ListPosts(tags string, page int) (*PostList, error) {
	params := url.Values{}
	params.Add("page", "dapi")
	params.Add("s", "post")
	params.Add("q", "index")
	params.Add("json", "1")
	params.Add("limit", postLimitStr)
	params.Add("tags", tags)
	params.Add("pid", strconv.Itoa(page-1)) // Pages are 0-indexed
	client.withAuth(params)

	rawResp, err := httpGet(ApiUrl, params)
	if err != nil {
		return nil, err
	}

	if rawResp.StatusCode != 200 {
		return nil, GelbooruError{Code: rawResp.StatusCode}
	}

	body, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}

	resp := FullPostResponse{
		Post: make([]PostResponse, 0, PostsPerPage),
	}
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Warn().Str("body", string(body)).Msg("failed to parse json")
		return nil, GelbooruError{Code: 500}
	}

	posts := make([]api.PostResponse, 0, len(resp.Post))
	for _, p := range resp.Post {
		data := api.PostResponse{
			Id:    p.Id,
			Score: p.Score,

			Rating:          p.Rating,
			SourceUrl:       p.SourceUrl,
			Uploader:        p.Uploader,
			UploaderUrl:     fmt.Sprintf("https://gelbooru.com/index.php?page=account&s=profile&id=%d", p.UploaderId),
			Tags:            strings.Split(p.Tags, " "),
			ThumbnailUrl:    p.PreviewUrl,
			ThumbnailWidth:  p.PreviewWidth,
			ThumbnailHeight: p.PreviewHeight,
			LowResUrl:       p.SampleUrl,
			LowResWidth:     p.SampleWidth,
			LowResHeight:    p.SampleHeight,
			ImageUrl:        p.ImageUrl,
			Width:           p.Width,
			Height:          p.Height,
		}

		if createdAt, err := time.Parse(time.RubyDate, p.CreatedAt); err == nil {
			data.CreatedAtTimestamp = createdAt.Unix()
		} else {
			log.Warn().Str("val", p.CreatedAt).Err(err).Msg("failed to parse post created_at")
		}

		posts = append(posts, data)
	}

	return &PostList{
		TotalCount: resp.Attributes.Count,
		Posts:      posts,
	}, nil
}

type TagInfo struct {
	Id        int
	Name      string
	Count     int
	Type      int
	Ambiguous int
}

type FullTagInfoResponse struct {
	Attributes struct {
		Limit  int
		Offset int
		Count  int
	} `json:"@attributes"`

	Tag []TagInfo
}

// ListTags returns a list of info found on the given tags (e.g. count, type).
// tags should be one or more tags separated by a space.
func (client *Client) ListTags(tags string) ([]api.TagResponse, error) {
	params := url.Values{}
	params.Add("page", "dapi")
	params.Add("s", "tag")
	params.Add("q", "index")
	params.Add("json", "1")
	params.Add("names", tags)
	client.withAuth(params)

	rawResp, err := httpGet(ApiUrl, params)
	if err != nil {
		return nil, err
	}

	if rawResp.StatusCode != 200 {
		return nil, GelbooruError{Code: rawResp.StatusCode}
	}

	body, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}

	var resp FullTagInfoResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Warn().Str("body", string(body)).Msg("failed to parse json")
		return nil, GelbooruError{Code: 500}
	}

	tagInfo := make([]api.TagResponse, resp.Attributes.Count)
	for i, t := range resp.Tag {
		if t.Name == "" {
			continue
		}

		tagInfo[i] = api.TagResponse{
			Name:  t.Name,
			Type:  ParseTagNumericType(t.Type),
			Count: t.Count,
		}
	}

	return tagInfo, nil
}
