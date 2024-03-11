package schema

import "music-app/adapter/database/model"

type UserRegisterReq struct {
	Email    string `json:"email"`
	UserIcon string `json:"icon"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterRes struct {
	AccessToken string     `json:"accessToken"`
	User        model.User `json:"user"`
}

type ProfileRes struct {
	Name     string `json:"name"`
	UserIcon string `json:"userIcon"`
}

type ProfileReq struct {
	UserId     string `query:"user_name"`
}

func ProfileResFromModel(User model.User) ProfileRes {
	return ProfileRes{
		Name:     User.Name,
		UserIcon: User.UserIcon,
	}
}
