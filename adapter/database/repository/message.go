package repository

import (
	"music-app/adapter/database/model"
	"music-app/usecase/port"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db   *gorm.DB
	ulid port.ULID
}

func NewMessageRepository(
	db *gorm.DB,
	ulid port.ULID,
) port.MessageRepository {
	return &MessageRepository{db: db, ulid: ulid}
}


func (r *MessageRepository) Search(query port.MessageSearchQuery) ([]model.Message, error) {
	sql := r.db.Model(&model.Message{})

	if query.BuiltinBoardId != "" {
		sql = sql.Where("builtin_board_id = ?", query.BuiltinBoardId)
	}


	var messages []model.Message
	if err := sql.
		Offset(query.Skip).
		Limit(query.Limit).
		Find(&messages).
		Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) Create(messageCreate model.Message) error {
	if err := r.db.Create(&messageCreate).Error; err != nil {
		return err
	}
	return nil
}
