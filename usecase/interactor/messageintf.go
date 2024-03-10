package interactor

import (
	"music-app/adapter/database/model"
	"time"
)

type MessageRegister struct {
	VoiceUrl       string
	UserId         string
	Time           time.Time
	BuiltinBoardId string
}


type IMessageUseCase interface {
	Register(MessageRegister) model.Message
	Search(messageSearch MessageSearch) ([]model.Message, error)
}
