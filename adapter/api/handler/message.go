package handler

import (
	"fmt"
	"music-app/adapter/api/schema"
	"music-app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo"
)

type MessageHandler struct {
	MessageUsecase interactor.IMessageUseCase
}

func NewMessageHandler(messageUsecase interactor.IMessageUseCase) *MessageHandler {
	return &MessageHandler{MessageUsecase: messageUsecase}
}

func (h *MessageHandler) Register(c echo.Context) error {
	var req schema.MessageRegisterReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message := h.MessageUsecase.Register(interactor.MessageRegister{
		VoiceUrl:       req.UserId,
		UserId:         req.UserId,
		Time:           req.Time,
		BuiltinBoardId: req.BuiltinBoardId,
	})
	return c.JSON(http.StatusCreated, message)
}

func (h *MessageHandler) Search(c echo.Context) error {

	var req schema.MessageSearchrReq
	if err := c.Bind(&req); err != nil {
		fmt.Errorf("Failed to bind request")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.MessageUsecase.Search(interactor.MessageSearch{
		BuiltinBoardId: req.BuiltinBoardId,
		Skip:           req.Skip,
		Limit:          req.Limit,
	})
	if err != nil {
		fmt.Errorf("Failed to search company")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
