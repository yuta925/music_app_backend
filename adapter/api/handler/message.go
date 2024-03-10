package handler

import (
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
		Text:           req.Text,
		VoiceUrl:       req.UserId,
		UserId:         req.UserId,
		Time:           req.Time,
		BuiltinBoardId: req.BuiltinBoardId,
	})
	return c.JSON(http.StatusCreated, message)
}
