package handler

import (
	"errors"
	"music-app/adapter/api/schema"
	"music-app/adapter/authentication"
	"music-app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo"
)

type AuthHandler struct {
	UserUC interactor.IUserUseCase
}

func NewAuthHandler(userUC interactor.IUserUseCase) *AuthHandler {
	return &AuthHandler{UserUC: userUC}
}

// Login POST /auth/access-token
func (h *AuthHandler) Login(c echo.Context) error {

	req := &schema.LoginReq{}
	if err := c.Bind(req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, token, err := h.UserUC.Login(req.Email, req.Password)
	if err != nil {

		switch {
		case errors.Is(err, interactor.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case errors.Is(err, authentication.ErrWrongPassword):
			return echo.NewHTTPError(http.StatusUnauthorized, errors.New("Authentication failed").Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	loginUser := &schema.LoginResUser{
		UserId: user.UserID,
		Email:  user.Email,
	}

	return c.JSON(http.StatusOK, &schema.LoginRes{
		AccessToken: token,
		TokenType:   schema.TokenType,
		User:        *loginUser,
	})
}
