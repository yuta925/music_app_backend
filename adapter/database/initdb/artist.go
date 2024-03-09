package initdb

import (
	"music-app/adapter/database/model"

	"gorm.io/gorm"
)

func CreateArtists(db *gorm.DB) ([]model.Artist, error) {
	retList := make([]model.Artist, 15)
	u := []model.Artist{
		{ArtistId: "1", Name: "The beatls"},
		{ArtistId: "2", Name: "SuperBeaver"},
		{ArtistId: "3", Name: "ive"},
		{ArtistId: "4", Name: "BTS"},
		{ArtistId: "5", Name: "嵐"},
		{ArtistId: "6", Name: "伊藤史人"},
		{ArtistId: "7", Name: "かなぶん"},
		{ArtistId: "8", Name: "あいこ"},
		{ArtistId: "9", Name: "queen"},
		{ArtistId: "10", Name: "radwimps"},
		{ArtistId: "11", Name: "The Fuhii"},
		{ArtistId: "12", Name: "いとう"},
		{ArtistId: "13", Name: "ふひいいい"},
		{ArtistId: "14", Name: "バレーボール"},
		{ArtistId: "15", Name: "プリン"},
	}
	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}

	return retList, nil

}
