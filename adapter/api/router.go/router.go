package router

import (
	"music-app/adapter/api/handler"
	// apiMiddleware "music-app/adapter/api/middleware"
	"music-app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(
	userUC interactor.IUserUseCase,
	builtinBoardUC interactor.IBuiltinBoardUseCase,
	messageUC interactor.IMessageUseCase,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowCredentials: true,
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	authHandler := handler.NewAuthHandler(userUC)
	userHandler := handler.NewUserHandler(userUC)
	builtinBoardHandler := handler.NewBuiltinBoardHandler(builtinBoardUC)
	messageHandler := handler.NewMessageHandler(messageUC)

	api := e.Group("/api")
	api.POST("/auth/access-token", authHandler.Login)

	user := api.Group("/users")
	user.POST("", userHandler.Register)
	user.GET("/me", userHandler.FindById)

	builtinboard := api.Group("/bulletin-board")
	builtinboard.POST("", builtinBoardHandler.Register)
	builtinboard.GET("", builtinBoardHandler.Search)

	message := api.Group("/message")
	message.POST("", messageHandler.Register)
	message.GET("", messageHandler.Search)

	return e
}
