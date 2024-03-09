package handler

import (
	"fmt"
	"music-app/adapter/api/schema"
	"music-app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo"
)

type BuiltinBoardHandler struct {
	BuiltinBoardUsecase interactor.IBuiltinBoardUseCase
}

func NewBuiltinBoardHandler(builtinBoardUsecase interactor.IBuiltinBoardUseCase) *BuiltinBoardHandler {
	return &BuiltinBoardHandler{BuiltinBoardUsecase: builtinBoardUsecase}
}

func (h *BuiltinBoardHandler) Register(c echo.Context) error {
	var req schema.BuiltinBoardRegisterReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	builtinBoard:= h.BuiltinBoardUsecase.Register(interactor.BuiltinBoardRegister{
		ImageUrl:   req.ImageUrl,
		LocationId: req.LocationId,
		Date:       req.Date,
		ArtistId:   req.ArtistId,
	})
	return c.JSON(http.StatusCreated, builtinBoard)
}

func (h *BuiltinBoardHandler) Search(c echo.Context) error {

	var req schema.BuiltinBoardSearchReq
	if err := c.Bind(&req); err != nil {
		fmt.Errorf("Failed to bind request")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.BuiltinBoardUsecase.Search(interactor.BuiltinBoardSearch{
		ArtistId:   req.ArtistId,
		LocationId: req.LocationId,
		Date:       *req.Date,
		Skip:       req.Skip,
		Limit:      req.Limit,
	})
	if err != nil {
		fmt.Errorf("Failed to search company")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
