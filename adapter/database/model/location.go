package model

type Location struct {
	LocationId string `gorm:"primaryKey;size:26"`
	Name     string `gorm:"size:255;not null"`
}
