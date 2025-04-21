package enum

type TagType byte

const (
	Tag TagType = iota
	_           // unknown
	Artist
	Copyright
	Character
	Metadata
	Deprecated
)
