package port

import (
	"music-app/adapter/database/model"
	"time"
)

type BuiltinBoardRepository interface {
	FindByID(BuiltinBoardId string) (model.BuiltinBoard, error)
	Search(query BuiltinBoardSearchQuery) ([]model.BuiltinBoard, error)
}

type BuiltinBoardSearchQuery struct {
	ArtistId   string
	LocationId string
	Date       *time.Time
	Skip       int
	Limit      int
}
