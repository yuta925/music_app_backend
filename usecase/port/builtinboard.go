package port

import "music-app/adapter/database/model"

type BuiltinBoardRepository interface {
	FindByID(BuiltinBoardId string) (model.BuiltinBoard, error)
}
