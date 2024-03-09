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

type ProfileRes struct {
	Name     string `json:"name"`
	UserIcon string `json:"userIcon"`
}

func ProfileResFromModel(User model.User) ProfileRes {
	return ProfileRes{
		Name:     User.Name,
		UserIcon: User.UserIcon,
	}
}
