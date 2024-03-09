package interactor

import (
	"errors"
	"fmt"
	"music-app/adapter/database/model"
	"music-app/usecase/port"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyUsed  = errors.New("email already used")
)

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

	if _, err := u.userRepo.FindByEmail(register.Email); err == nil {
		return model.User{}, "", fmt.Errorf("Email is already existed")
	} else if !errors.Is(err, ErrUserNotFound) {
		return model.User{}, "", err
	}

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

	ret, err := u.userRepo.FindByID(newUser.UserID)
	return ret, token, err
}

func (u *UserUseCase) FindByID(userID string) (model.User, error) {
	return u.userRepo.FindByID(userID)
}

func (u *UserUseCase) Authenticate(token string) (string, error) {
	return u.userAuth.Authenticate(token)
}
