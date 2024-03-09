package authentication

import (
	"music-app/adapter/database/model"
	"music-app/usecase/port"
	"time"
)

type UserAuth struct{}

func NewUserAuth() port.UserAuth {
	return &UserAuth{}
}

func (a *UserAuth) CheckPassword(u model.User, password string) error {
	return CheckBcryptPassword(u.HashedPassword, password)
}

func (a *UserAuth) IssueUserToken(user model.User, issuedAt time.Time) (string, error) {
	return IssueUserToken(user.UserID)
}

func (a *UserAuth) Authenticate(token string) (string, error) {
	return VerifyUserToken(token)
}

func (a *UserAuth) HashPassword(password string) (string, error) {
	hp, err := HashBcryptPassword(password)
	if err != nil {
		return "", err
	}
	return hp, nil
}
