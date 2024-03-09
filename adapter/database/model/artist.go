package model

type Artist struct {
	ArtistId string `gorm:"primaryKey;size:26"`
	Name     string `gorm:"size:255;not null"`
}
