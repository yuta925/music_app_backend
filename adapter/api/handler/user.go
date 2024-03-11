package handler

import (
	"errors"
	"fmt"
	"music-app/adapter/api/schema"
	"music-app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserUsecase interactor.IUserUseCase
}

func NewUserHandler(userUsecase interactor.IUserUseCase) *UserHandler {
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

func (h *UserHandler) FindById(c echo.Context) error {
	var req schema.ProfileReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res, err := h.UserUsecase.FindByID(interactor.UserSearch{
		UserId: req.UserId,
	})
	if err != nil {
		fmt.Errorf("Failed to search company")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
