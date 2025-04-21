package api

type TagType byte

const (
	Tag TagType = iota
	Artist
	Copyright
	Character
	Metadata
	Deprecated
	Unknown
)

type BooruTag struct {
	Name  string
	Type  TagType
	Count int
}
