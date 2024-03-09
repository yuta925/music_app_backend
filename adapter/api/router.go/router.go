package router

import (
	"music-app/adapter/api/handler"
	"music-app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(
	userUC interactor.IUserUseCase,

) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	authHandler := handler.NewAuthHandler(userUC)
	userHandler := handler.NewUserHandler(userUC)

	api := e.Group("/api")
	api.POST("/auth/access-token", authHandler.Login)

	user := api.Group("/users")
	user.POST("", userHandler.Register)

	return e
}
