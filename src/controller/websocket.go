package controller

import (
	"github.com/labstack/echo/v4"
	"homework/middleware/token"
	"homework/websocket"
	"homework/usecase"
	apperror "homework/error"
	"net/http"
)

type IWebSocketController interface {
	ServeRoomWs(ctx echo.Context) error
}

type WebSocketController struct {
	hub *websocket.Hub
	chatUseCase usecase.IChatUsecase
}

func NewWebSocketController(hub *websocket.Hub, chatUseCase usecase.IChatUsecase) IWebSocketController {
	return &WebSocketController{hub, chatUseCase}
}

func (c *WebSocketController) ServeRoomWs(ctx echo.Context) error {
	userID, err := token.GetUserIDWithTokenCheck(ctx)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, apperror.ErrorWrapperWithCode(err, http.StatusUnauthorized))
	}

	roomID := ctx.Param("room_id")
	if roomID == "" {
		return ctx.JSON(http.StatusBadRequest, apperror.ErrorWrapperWithCode(err, http.StatusBadRequest))
	}

	return c.hub.ServeWS(ctx, userID, roomID)
} 