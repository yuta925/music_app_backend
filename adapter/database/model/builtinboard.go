package model

import "time"

type BuiltinBoard struct {
	BuiltinBoardId string `gorm:"primaryKey;size:26"`
	ImageUrl       string `gorm:"size:255; not null"`
	Date           time.Time
	LocationId     string
	ArtistId       string
}
