package initdb

import (
	"music-app/adapter/database/model"

	"gorm.io/gorm"
)

func CreateLocations(db *gorm.DB) ([]model.Location, error) {
	retList := make([]model.Location, 15)
	u := []model.Location{
		{LocationId: "1", Name: "東京ドーム"},
		{LocationId: "2", Name: "武道館"},
		{LocationId: "3", Name: "大阪城ホーム"},
		{LocationId: "4", Name: "paypayドーム"},
		{LocationId: "5", Name: "名古屋ドーム"},
		{LocationId: "6", Name: "岡山ドーム"},
		{LocationId: "7", Name: "俺の家"},
		{LocationId: "8", Name: "関西学院大学"},
		{LocationId: "9", Name: "サウナ"},
		{LocationId: "10", Name: "ガスト"},
	}
	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}

	return retList, nil

}
