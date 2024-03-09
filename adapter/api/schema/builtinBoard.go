package schema

import "time"

type BuiltinBoardRegisterReq struct {
	ImageUrl       string     `json:"imageurl "`
	Date           *time.Time `json:"date"`
	LocationId     string     `json:"locationid"`
	ArtistId       string     `json:"artistid"`
}

type BuiltinBoardRegisterRes struct {
	BuiltinBoardId string      `json:"builtinboardid"`
	ImageUrl       string      `json:"imageurl "`
	Date           *time.Time  `json:"date"`
	Location       LocationRes `json:"location"`
	Artist         ArtistRes   `json:"artist"`
}
