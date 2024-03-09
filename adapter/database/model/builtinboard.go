package model

import "time"


type BuiltinBoard struct {
	BuiltinBoardId string `gorm:"primaryKey;size:26"`
	ImageUrl       string `gorm:"size:26; not null"`
	Date           *time.Time
	LocationId     string
	ArtistId       string
	Location       Location   `gorm:"foreignKey:LocationId"`
	Artist         Artist     `gorm:"foreignKey:ArtistId"`
}

