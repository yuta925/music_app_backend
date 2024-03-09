package schema

import "music-app/adapter/database/model"




type UserRegisterReq struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	UserIcon string `json:"userIcon"`
}

type RegisterRes struct {
	AccessToken string     `json:"accessToken"`
	User        model.User `json:"user"`
}
