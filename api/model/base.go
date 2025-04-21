package model

import "time"

type BaseModel struct {
	Id        int64
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
