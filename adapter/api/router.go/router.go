package router

import (
	"music-app/adapter/api/handler"
	apiMiddleware "music-app/adapter/api/middleware"
	"music-app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(
	userUC interactor.IUserUseCase,
	builtinBoardUC interactor.IBuiltinBoardUseCase,

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

	api := e.Group("/api")
	api.POST("/auth/access-token", authHandler.Login)

	auth := api.Group("", apiMiddleware.NewAuthMiddleware(userUC).Authenticate(true))

	user := auth.Group("/users")
	user.POST("", userHandler.Register)
	user.GET("/me", userHandler.FindMe)

	builtinboard := auth.Group("/builtinboards")
	builtinboard.POST("", builtinBoardHandler.Register)
	builtinboard.GET("", builtinBoardHandler.Search)

	return e
}
