package repository

import (
	"fmt"
	"music-app/adapter/database/model"
	"music-app/usecase/port"
	"time"

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

func (r *BuiltinBoardRepository) Create(buitinBoardCreate model.BuiltinBoard) error {
	if err := r.db.Create(&buitinBoardCreate).Error; err != nil {
		return err
	}
	return nil
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
	fmt.Println(query)
	sql := r.db.Model(&model.BuiltinBoard{})

	if query.ArtistId != "" {
		sql = sql.Where("artist_id = ?", query.ArtistId)
	}
	if query.LocationId != "" {
		sql = sql.Where("location_id = ?", query.LocationId)
	}
	dateEnd := time.Date(query.Date.Year(), query.Date.Month(), query.Date.Day(), 23, 59, 59, 0, query.Date.Location())
	if query.Date != (time.Time{}) {
		sql = sql.Where("date BETWEEN ? AND ?", query.Date, dateEnd)
	}
	fmt.Println(sql)

	var builtinBoards []model.BuiltinBoard
	if err := sql.
		Offset(query.Skip).
		Limit(query.Limit).
		Find(&builtinBoards).
		Error; err != nil {
		return nil, err
	}
	fmt.Println(builtinBoards)
	return builtinBoards, nil
}
