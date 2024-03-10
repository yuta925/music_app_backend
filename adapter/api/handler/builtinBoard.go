package handler

import (
	"fmt"
	"music-app/adapter/api/schema"
	"music-app/usecase/interactor"
	"net/http"
	"time"

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
	t, _ := time.Parse(time.RFC3339Nano, req.Date)
	dateOnly := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	builtinBoard := h.BuiltinBoardUsecase.Register(interactor.BuiltinBoardRegister{
		ImageUrl:   req.ImageUrl,
		LocationId: req.LocationId,
		Date:       dateOnly,
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
	fmt.Println(req)
	// req.Dateをtimeに直す
	t, _ := time.Parse(time.RFC3339Nano, req.LiveDate)
	dateOnly := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	res, err := h.BuiltinBoardUsecase.Search(interactor.BuiltinBoardSearch{
		ArtistId:   req.ArtistId,
		LocationId: req.LocationId,
		Date:       dateOnly,
		Skip:       req.Skip,
		Limit:      req.Limit,
	})
	if err != nil {
		fmt.Errorf("Failed to search company")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
