package controller

import (
	"homework/domain/model/room"
	"homework/websocket"
	"log"

	"github.com/labstack/echo/v4"
)

type IWebSocketController interface {
	ServeRoomWs(c echo.Context) error
}

type webSocketController struct {
	Hub *websocket.Hub
}

func NewWebSocketController(hub *websocket.Hub) IWebSocketController {
	return &webSocketController{hub}
}

func (w *webSocketController)ServeRoomWs(c echo.Context) error {
	roomID := c.Param("room_id")
	hub := w.Hub
	go hub.Run()
	room.RoomToHub[roomID] = hub
	serveWs(hub, c)
	return nil
}

func serveWs(hub *websocket.Hub, c echo.Context) {
	conn, err := websocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &websocket.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client                                                   //Hubにregisterする

	go client.WriteMessage()
	go client.ReadMessage()
}
