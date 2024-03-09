package handler

import (
	"errors"
	"music-app/adapter/api/schema"
	"music-app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserUsecase interactor.IUserUseCase
}

func NewWorkerUserHandler(userUsecase interactor.IUserUseCase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) Register(c echo.Context) error {
	var req schema.UserRegisterReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, token, err := h.UserUsecase.Register(interactor.UserRegister{
		UserName: req.Name,
		UserIcon: req.UserIcon,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		switch {
		case errors.Is(err, interactor.ErrUserAlreadyExists):
			return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(400))
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	registerRes := schema.RegisterRes{
		AccessToken: token,
		User:        user,
	}
	return c.JSON(http.StatusCreated, registerRes)
}
