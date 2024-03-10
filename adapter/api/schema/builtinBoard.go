package schema

import "time"

type BuiltinBoardRegisterReq struct {
	ImageUrl   string `json:"live_image"`
	Date       string `json:"live_date"`
	LocationId int    `json:"live_venue_id"`
	ArtistId   int    `json:"artist_id"`
}

type BuiltinBoardRegisterRes struct {
	BuiltinBoardId string      `json:"builtinboardid"`
	ImageUrl       string      `json:"live_image"`
	Date           time.Time   `json:"live_date"`
	Location       LocationRes `json:"live_venue_id"`
	Artist         ArtistRes   `json:"artist"`
}

type BuiltinBoardSearchReq struct {
	Date       string `json:"live_date"`
	LocationId string    `json:"locationid"`
	ArtistId   string    `json:"artistid"`
	Skip       int       `query:"skip"`
	Limit      int       `query:"limit"`
}
