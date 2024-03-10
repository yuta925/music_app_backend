package schema

import "time"

type MessageRegisterReq struct {
	VoiceUrl       string    `json:"voice_url"`
	UserId         string    `json:"user_id"`
	Time           time.Time `json:"time"`
	BuiltinBoardId string    `json:"builtin_board_id"`
}

type MessageSearchrReq struct {
	BuiltinBoardId string `json:"builtin_board_id"`
	Skip           int    `query:"skip"`
	Limit          int    `query:"limit"`
}
