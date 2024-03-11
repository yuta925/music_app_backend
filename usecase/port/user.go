package port

import (
	"music-app/adapter/database/model"
	"time"
)

type UserRepository interface {
	FindByEmail(email string) (model.User, error)
	FindByID(query UserSearchQuery) (model.User, error)
	Create(userCreate model.User) error
}
type UserAuth interface {
	Authenticate(token string) (string, error)
	CheckPassword(user model.User, password string) error
	HashPassword(password string) (string, error)
	IssueUserToken(user model.User, issuedAt time.Time) (string, error)
}

type UserSearchQuery struct {
	UserId string
}
