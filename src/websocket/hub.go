// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"encoding/json"
	"homework/domain/model/chat"
	"log"
	"net/http"

	"homework/usecase"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	UserID string
	RoomID string
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	Register   chan *Client
	unregister chan *Client
	done       chan struct{}
	upgrader   websocket.Upgrader
	chatUseCase usecase.IChatUsecase
}

func NewHub(chatUseCase usecase.IChatUsecase) *Hub {
	return &Hub{
		broadcast:   make(chan []byte),
		Register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[*Client]bool),
		chatUseCase: chatUseCase,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (h *Hub) Run() {
	defer func() { close(h.done) }()
	for msg := range h.broadcast {
		for client := range h.clients {
			select {
			case client.Send <- msg:
			default:
				close(client.Send)
				delete(h.clients, client)
			}
		}
	}
}

func (h *Hub) readMessage(c *Client) {
	defer func() {
		h.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		// メッセージをChatRequestに変換
		var chatReq chat.ChatRequest
		if err := json.Unmarshal(message, &chatReq); err != nil {
			log.Printf("error: %v", err)
			continue
		}

		// CreateChatRequestを作成
		createReq := chat.Chat{
			UserID:  c.UserID,
			Message: chatReq.Message,
			RoomID:  c.RoomID,
		}

		// DBに保存
		res, err := h.chatUseCase.Create(createReq)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		// レスポンスをブロードキャスト
		resBytes, err := json.Marshal(res)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		h.broadcast <- resBytes
	}
}

func (h *Hub) writeMessage(c *Client) {
	defer c.Conn.Close()
	
	for msg := range c.Send {
		ch := chat.ChatResponse{}
		if err := json.Unmarshal(msg, &ch); err != nil {
			log.Printf("error: %v", err)
			continue
		}

		message, err := json.Marshal(ch)
		if err != nil {
			log.Printf("error: %v", err)
			continue
		}

		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("error: %v", err)
			return
		}
	}
}

func (h *Hub) ServeWS(c echo.Context, userID, roomID string) error {
	conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	client := &Client{
		Hub:    h,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		UserID: userID,
		RoomID: roomID,
	}
	h.Register <- client

	go h.readMessage(client)
	go h.writeMessage(client)

	return nil
}
