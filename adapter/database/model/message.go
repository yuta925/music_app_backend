package model
import "time"

type Message struct {
	MessageId      string
	Text           string
	VoiceUrl       string
	UserId         string
	Time           time.Time
	BuiltinBoardId string
}
