package model

type User struct {
	UserID         string `gorm:"primaryKey;size:26"`
	Name           string `gorm:"size:26;not null"`
	Email          string `gorm:"size:255;unique;not null"`
	UserIcon       string `gorm:"size:255; not null"`
	HashedPassword	   string `gorm:"size:255;not null"`
}



