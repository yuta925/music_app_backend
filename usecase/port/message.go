package port

import "music-app/adapter/database/model"

type MessageRepository interface {
	Search(query MessageSearchQuery) ([]model.Message, error)
	Create(messageCreate model.Message) error
}

type MessageSearchQuery struct {
	BuiltinBoardId string
	Skip           int
	Limit          int
}
