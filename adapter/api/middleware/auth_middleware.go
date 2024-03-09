package middleware

import (
	"context"
	"errors"
	"music-app/adapter/database/model"
)

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
