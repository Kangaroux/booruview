package api

type TagType string

const (
	Tag       TagType = "tag"
	Artist    TagType = "artist"
	Copyright TagType = "copyright"
	Character TagType = "character"
	Metadata  TagType = "metadata"
	Unknown   TagType = "unknown"
)

func ParseTagType(val string) TagType {
	switch val {
	case string(Tag):
		return Tag
	case string(Artist):
		return Artist
	case string(Copyright):
		return Copyright
	case string(Character):
		return Character
	case string(Metadata):
		return Metadata
	case string(Unknown):
		return Unknown
	default:
		return Unknown
	}
}

type TagResponse struct {
	Name  string  `json:"name"`
	Type  TagType `json:"type"`
	Count int     `json:"count"`
}

type PostResponse struct {
	Id                 int      `json:"id"`
	CreatedAtTimestamp int64    `json:"created_at"`
	Score              int      `json:"score"`
	Rating             string   `json:"rating"`
	SourceUrl          string   `json:"source_url"`
	Uploader           string   `json:"uploader"`
	UploaderUrl        string   `json:"uploader_url"`
	Tags               []string `json:"tags"`
	ThumbnailUrl       string   `json:"thumbnail_url"`
	ThumbnailWidth     int      `json:"thumbnail_width"`
	ThumbnailHeight    int      `json:"thumbnail_height"`
	LowResUrl          string   `json:"lowres_url"`
	LowResWidth        int      `json:"lowres_width"`
	LowResHeight       int      `json:"lowres_height"`
	ImageUrl           string   `json:"image_url"`
	Width              int      `json:"width"`
	Height             int      `json:"height"`
}
