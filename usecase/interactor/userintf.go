package interactor

import "music-app/adapter/database/model"

type UserRegister struct {
	UserName string
	UserIcon string
	Password string
	Email    string
}

type IUserUseCase interface {
	Login(email, password string) (model.User, string, error)
	Register(UserRegister) (model.User, string, error)
	FindByID(userSearch UserSearch) (model.User, error)
	Authenticate(token string) (string, error)
}
