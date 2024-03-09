package repository

import (
	"errors"
	"fmt"
	"music-app/adapter/database/model"
	"music-app/usecase/interactor"
	"music-app/usecase/port"

	"gorm.io/gorm"
)

type UserRepository struct {
	db   *gorm.DB
	ulid port.ULID
}

func NewUserRepository(
	db *gorm.DB,
	ulid port.ULID,
) port.UserRepository {
	return &UserRepository{db: db, ulid: ulid}
}


func (r *UserRepository) FindByEmail(email string) (model.User, error) {
	ret := &model.User{}
	err := r.db.
		Where("email = ?", email).
		First(ret).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, interactor.ErrUserNotFound
	}
	if err != nil {
		return model.User{}, err
	}
	return *ret, nil
}

func (r *UserRepository) FindByID(UserId string) (model.User, error) {
	user := &model.User{}
	err := r.db.
		Model(&model.User{}).
		Where("user_id = ?", UserId).
		First(user).
		Error
	if err != nil {
		return model.User{}, err
	}
	fmt.Println(*user)
	return *user, nil
}
