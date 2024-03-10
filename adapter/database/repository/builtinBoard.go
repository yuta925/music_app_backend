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
		Model(&model.BuiltinBoard{}).
		Where("builtin_board_id = ?", BuiltinBoardId).
		First(builtinboard).
		Error
	if err != nil {
		return model.BuiltinBoard{}, err
	}
	return *builtinboard, nil
}

func (r *BuiltinBoardRepository) Search(query port.BuiltinBoardSearchQuery) ([]model.BuiltinBoard, error) {
	sql := r.db.Model(&model.BuiltinBoard{})

	if query.ArtistId != "" {
		sql = sql.Where("artist_id = ?", query.ArtistId)
	}
	if query.LocationId != "" {
		sql = sql.Where("location_id = ?", query.LocationId)
	}
	if query.Date != nil {
		sql = sql.Where("date = ?", query.Date)
	}


	var builtinBoards []model.BuiltinBoard
	if err := sql.
		Offset(query.Skip).
		Limit(query.Limit).
		Find(&builtinBoards).
		Error; err != nil {
		return nil, err
	}

	return builtinBoards, nil
}


