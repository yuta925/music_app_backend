package repository

import (
	"gorm.io/gorm"
	"music-app/usecase/port"
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
