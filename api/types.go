package api

type TagType string

const (
	Tag        TagType = "tag"
	Artist     TagType = "artist"
	Copyright  TagType = "copyright"
	Character  TagType = "character"
	Metadata   TagType = "metadata"
	Deprecated TagType = "deprecated"
	Unknown    TagType = "unknown"
)

type TagResponse struct {
	Name  string  `json:"name"`
	Type  TagType `json:"type"`
	Count int     `json:"count"`
}

type PostResponse struct {
	Id                 int    `json:"id"`
	CreatedAtTimestamp int64  `json:"created_at"`
	Score              int    `json:"score"`
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	Rating             string `json:"rating"`
	SourceUrl          string `json:"source_url"`
	Uploader           string `json:"uploader"`
	UploaderUrl        string `json:"uploader_url"`
	Tags               string `json:"tags"`
	ImageUrl           string `json:"file_url"`
	PreviewUrl         string `json:"preview_url"`
	SampleUrl          string `json:"sample_url"`
}
