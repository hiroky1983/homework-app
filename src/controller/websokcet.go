package controller

import (
	"homework/domain/model/room"
	"homework/domain/repository"
	"homework/middleware/token"
	"homework/websocket"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type IWebSocketController interface {
	ServeRoomWs(c echo.Context) error
}

type webSocketController struct {
	Hub *websocket.Hub
	rr  repository.IRoomRepository
	db  *bun.DB
}

func NewWebSocketController(hub *websocket.Hub, rr repository.IRoomRepository, db *bun.DB) IWebSocketController {
	return &webSocketController{hub, rr, db}
}

func (w *webSocketController) ServeRoomWs(c echo.Context) error {
	roomID := c.Param("room_id")
	userID, err := token.GetUserIDWithTokenCheck(c)
	if err != nil {
		return c.JSON(401, err)
	}
	hub := w.Hub
	go hub.Run()
	r := []*room.RoomMap{}
	userIDs, err := w.rr.GetUserIDByRoomID(w.db, r, roomID)
	if err != nil {
		return c.JSON(500, err)
	}
	room.RoomToHub[roomID] = hub
	serveWs(hub, c, userID, userIDs)
	return nil
}

func serveWs(hub *websocket.Hub, c echo.Context, userID string, userIDs []string) {
	conn, err := websocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &websocket.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client //Hubにregisterする

	go client.WriteMessage(userID, userIDs)
	go client.ReadMessage()
}
