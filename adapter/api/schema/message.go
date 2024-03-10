package schema

import "time"

type MessageRegisterReq struct {
	Text           string    `json:"text"`
	VoiceUrl       string    `json:"voice_url"`
	UserId         string    `json:"user_id"`
	Time           time.Time `json:"time"`
	BuiltinBoardId string    `json:"builtin_board_id"`
}
