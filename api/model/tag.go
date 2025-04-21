package model

import "github.com/kangaroux/booru-viewer/enum"

type Tag struct {
	BaseModel

	Count int
	Type  enum.TagType
	Name  string
}
