package port

import (
	"music-app/adapter/database/model"
	"time"
)



type UserRepository interface {
	FindByEmail(email string) (model.User, error)
	Create(user model.User) error
	FindByID(UserId string) (model.User, error)
}
type UserAuth interface {
	Authenticate(token string) (string, error)
	CheckPassword(user model.User, password string) error
	HashPassword(password string) (string, error)
	IssueUserToken(user model.User, issuedAt time.Time) (string, error)
}


