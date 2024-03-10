package interactor

import (
	"music-app/adapter/database/model"
	"music-app/usecase/port"
)

type MessageUseCase struct {
	ulid        port.ULID
	MessageRepo port.MessageRepository
}

type MessageSearch struct {
	BuiltinBoardId string
	Skip       int
	Limit      int
}

func NewMessageUseCase(
	ulid port.ULID,
	messageRepo port.MessageRepository,
) IMessageUseCase {
	return &MessageUseCase{
		ulid:        ulid,
		MessageRepo: messageRepo,
	}
}

func (u *MessageUseCase) Register(register MessageRegister) model.Message {

	newMessage := model.Message{
		MessageId:  u.ulid.GenerateID(),
	    Text:           register.Text,
    	VoiceUrl:       register.VoiceUrl,
    	UserId:         register.UserId,
    	Time:           register.Time,
    	BuiltinBoardId: register.BuiltinBoardId,
	}

	return newMessage
}

func (u *MessageUseCase) Search(messageSearch MessageSearch) ([]model.Message, error) {
	return u.MessageRepo.Search(port.MessageSearchQuery{
		BuiltinBoardId:   messageSearch.BuiltinBoardId,
		Skip:       messageSearch.Skip,
		Limit:      messageSearch.Limit,
	})
}
