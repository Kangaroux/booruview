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
