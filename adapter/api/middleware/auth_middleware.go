package middleware

import (
	"context"
	"errors"
	"fmt"
	"music-app/adapter/api/schema"
	"music-app/adapter/database/model"
	"music-app/usecase/interactor"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type AuthMiddleware struct {
	userUC interactor.IUserUseCase
}

func NewAuthMiddleware(userUC interactor.IUserUseCase) *AuthMiddleware {
	return &AuthMiddleware{userUC}
}

// Authenticate
// tokenを取得して、認証するmiddlewareを返す関数. isRequireActivatedがtrueの場合、仮登録ユーザーは認証できない
func (m *AuthMiddleware) Authenticate(isRequireActivated bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get JWT Token From Header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, schema.TokenType+" ") {
				fmt.Errorf("Failed to authenticate")
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			token := strings.TrimPrefix(authHeader, schema.TokenType+" ")

			// Authenticate
			userID, err := m.userUC.Authenticate(token)
			if err != nil {
				fmt.Errorf("Failed to authenticate")
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			user, err := m.userUC.FindByID(userID)
			if err != nil {
				fmt.Errorf("Failed to find me")
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			c = SetToContext(c, user)
			return next(c)
		}
	}
}
func SetToContext(c echo.Context, user model.User) echo.Context {
	ctx := c.Request().Context()
	ctx = SetUserToContext(ctx, user)
	c.SetRequest(c.Request().WithContext(ctx))
	return c
}
func SetUserToContext(ctx context.Context, user model.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

type ContextKey string

var (
	userKey ContextKey = "userKey"
)

func GetUserFromContext(ctx context.Context) (model.User, error) {
	v := ctx.Value(userKey)
	user, ok := v.(model.User)
	if !ok {
		return model.User{}, errors.New("no user found in context")
	}
	return user, nil
}
