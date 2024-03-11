package interactor

import (
	"errors"
	"fmt"
	"log"
	"music-app/adapter/database/model"
	"music-app/usecase/port"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyUsed  = errors.New("email already used")
)

type UserSearch struct {
	UserId string
}
type UserUseCase struct {
	clock    port.Clock
	ulid     port.ULID
	userAuth port.UserAuth
	userRepo port.UserRepository
}

func NewUserUseCase(
	clock port.Clock,
	ulid port.ULID,
	userAuth port.UserAuth,
	userRepo port.UserRepository,
) IUserUseCase {
	return &UserUseCase{
		clock:    clock,
		ulid:     ulid,
		userAuth: userAuth,
		userRepo: userRepo,
	}
}

func (u *UserUseCase) Login(email, password string) (model.User, string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return model.User{}, "", err
	}

	err = u.userAuth.CheckPassword(user, password)
	if err != nil {
		return model.User{}, "", err
	}

	token, err := u.userAuth.IssueUserToken(user, u.clock.Now())
	if err != nil {
		return model.User{}, "", err
	}
	return user, token, nil
}

func (u *UserUseCase) Register(register UserRegister) (model.User, string, error) {

	hashedPassword, err := u.userAuth.HashPassword(register.Password)
	if err != nil {
		return model.User{}, "", err
	}

	newUser := model.User{
		UserID:         u.ulid.GenerateID(),
		Name:           register.UserName,
		UserIcon:       register.UserIcon,
		HashedPassword: hashedPassword,
		Email:          register.Email,
	}
	token, err := u.userAuth.IssueUserToken(newUser, u.clock.Now())
	if err != nil {
		return model.User{}, "", err
	}
	fmt.Println(newUser)

	e := u.userRepo.Create(newUser)
	if e != nil {
		log.Println("Error:", err)
	}

	return newUser, token, err
}

func (u *UserUseCase) FindByID(userSearch UserSearch) (model.User, error) {
	return u.userRepo.FindByID(port.UserSearchQuery{
		UserId: userSearch.UserId,
	})
}

func (u *UserUseCase) Authenticate(token string) (string, error) {
	return u.userAuth.Authenticate(token)
}
