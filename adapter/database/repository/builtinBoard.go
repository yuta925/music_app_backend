package repository

import (
	"music-app/adapter/database/model"
	"music-app/usecase/port"

	"gorm.io/gorm"
)

type BuiltinBoardRepository struct {
	db   *gorm.DB
	ulid port.ULID
}

func NewBuiltinBoardRepository(
	db *gorm.DB,
	ulid port.ULID,
) port.BuiltinBoardRepository {
	return &BuiltinBoardRepository{db: db, ulid: ulid}
}

func (r *BuiltinBoardRepository) FindByID(BuiltinBoardId string) (model.BuiltinBoard, error) {
	builtinboard := &model.BuiltinBoard{}
	err := r.db.
		Model(&model.User{}).
		Where("builtinboard_id = ?", BuiltinBoardId).
		First(builtinboard).
		Error
	if err != nil {
		return model.BuiltinBoard{}, err
	}
	return model.BuiltinBoard{}, nil
}


