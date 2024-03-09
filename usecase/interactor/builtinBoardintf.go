package interactor

import (
	"music-app/adapter/database/model"
	"time"
)

type BuiltinBoardRegister struct {
	BuiltinBoardId string
	ImageUrl       string
	Date           *time.Time
	LocationId     string
	ArtistId       string
}

type IBuiltinBoardUseCase interface {
	Register(BuiltinBoardRegister) (model.BuiltinBoard, error)
}
