package schema

const TokenType = "Bearer"

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResUser struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
}

type LoginRes struct {
	AccessToken string       `json:"accessToken"`
	TokenType   string       `json:"tokenType"`
	User        LoginResUser `json:"user"`
}
