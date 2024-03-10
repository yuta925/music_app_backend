package interactor

import (
	"music-app/adapter/database/model"
	"time"
)

type BuiltinBoardRegister struct {
	ImageUrl       string
	Date           time.Time
	LocationId     string
	ArtistId       string
}

type IBuiltinBoardUseCase interface {
	Register(BuiltinBoardRegister) (model.BuiltinBoard)
	Search(builtinBoardSearch BuiltinBoardSearch) ([]model.BuiltinBoard, error)
}
