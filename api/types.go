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
	Id                 int
	CreatedAtTimestamp int64 `json:"created_at"`
	Score              int
	Width              int
	Height             int
	Rating             string
	SourceUrl          string `json:"source_url"`
	Uploader           string
	UploaderUrl        string `json:"uploader_url"`
	Tags               string
	ImageUrl           string `json:"file_url"`
	PreviewUrl         string `json:"preview_url"`
	SampleUrl          string `json:"sample_url"`
}
