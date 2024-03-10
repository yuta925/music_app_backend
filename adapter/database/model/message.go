package model
import "time"

type Message struct {
	MessageId      string
	VoiceUrl       string
	UserId         string
	Time           time.Time
	BuiltinBoardId string
}
