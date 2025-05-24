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

type TagSearchResponse struct {
	Type     string
	Label    string
	Value    string
	Count    string `json:"post_count"`
	Category string
}

var (
	tagTypeMap = map[string]api.TagType{
		"tag":       api.Tag,
		"artist":    api.Artist,
		"character": api.Character,
		"copyright": api.Copyright,
		"metadata":  api.Metadata,
	}

	tagNumericTypeMap = map[int]api.TagType{
		0: api.Tag,
		1: api.Artist,
		3: api.Copyright,
		4: api.Character,
		5: api.Metadata,
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

func doApiTagSearch(query string) ([]api.TagResponse, error) {
	var resp []TagSearchResponse
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

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
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

func SearchTags(query string) ([]api.TagResponse, error) {
	isFilter, suggestions := SuggestedSearchFilters(query)

	if !isFilter {
		return doApiTagSearch(query)
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

func ListPosts(tags string, page int) (*PostList, error) {
	params := url.Values{}
	params.Add("page", "dapi")
	params.Add("s", "post")
	params.Add("q", "index")
	params.Add("json", "1")
	params.Add("limit", postLimitStr)
	params.Add("tags", tags)
	params.Add("pid", strconv.Itoa(page-1)) // Pages are 0-indexed

	rawResp, err := api.HttpGet(ApiUrl + "?" + params.Encode())
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}

	resp := FullPostResponse{
		Post: make([]PostResponse, PostsPerPage),
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	posts := make([]api.PostResponse, len(resp.Post))
	for i, p := range resp.Post {
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
			log.Println("warning: failed to parse post created_at:", err)
		}

		posts[i] = data
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
func ListTags(tags string) ([]api.TagResponse, error) {
	params := url.Values{}
	params.Add("page", "dapi")
	params.Add("s", "tag")
	params.Add("q", "index")
	params.Add("json", "1")
	params.Add("names", tags)

	rawResp, err := api.HttpGet(ApiUrl + "?" + params.Encode())
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}

	var resp FullTagInfoResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
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
