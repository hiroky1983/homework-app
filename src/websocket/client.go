// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"homework/domain/model/chat"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadMessage() {
	defer func() {
		c.Hub.unregister <- c //Hubからunregisterして
		c.Conn.Close()        //connectionをcloseする
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Printf("error: %v", err)
	}
	c.Conn.SetPongHandler(func(string) error {
		if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			log.Printf("error: %v", err)
		}
		return nil
	}) //何かあればReadDeadlineを延長する
	for {
		_, message, err := c.Conn.ReadMessage() //clientがmessageを送れば，c.conn.ReadMessage()でmessageが受け取れる
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1)) //messageの改行をスペースに変える
		c.Hub.broadcast <- message                                            //Hubに送る
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WriteMessage(e echo.Context) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()  //tickerを止めて
		c.Conn.Close() //connectionをcloseする
	}()
	for {
		select {
		case message, ok := <-c.Send: //messageが送られてきたら，c.sendからmessageを取り出せる
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("error: %v", err)
			} //WriteDeadlineを延長する
			if !ok {
				// The hub closed the channel.
				if err := c.Conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Printf("error: %v", err)
				}
				return
			}
			chat := chat.ChatRequest{}

			if err := json.Unmarshal(message, &chat); err != nil {
				log.Printf("error: %v", err)
			}
			fmt.Println("=================",chat)
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			if _, err := w.Write(message); err != nil {
				log.Printf("error: %v", err)
			}

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ { //さらにc.sendにデータがあればそれも送る
				if _, err := w.Write(newline); err != nil {
					log.Printf("error: %v", err)
				}
				if _, err := w.Write(<-c.Send); err != nil {
					log.Printf("error: %v", err)
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("error: %v", err)
			}
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
