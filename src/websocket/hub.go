// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool //Clientのpointerがkeyでvalueがbool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	//外部から停止通知を送るためのchannel
	stop chan struct{}

	//goroutineの終了時に外部へ完了通知を送るためのchannel
	done chan struct{}
}

func NewHub() *Hub { //新たにHubを作ってそのpointerを返す
	return &Hub{
		broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		stop:       make(chan struct{}),
		done:       make(chan struct{}),
	}
}

func (h *Hub) Run() {
	defer func() { close(h.done) }()
	for {
		select {
		case client := <-h.Register: //Hubのregisterというchannelに*Clientが入っているとき
			h.clients[client] = true //clientを登録する
		case client := <-h.unregister: //Hubのunregisterというchannelに*Clientが入っているとき
			if _, ok := h.clients[client]; ok { //そのclientが登録されていれば
				delete(h.clients, client) //削除する
				close(client.Send)        //そのclientのchannelをcloseする
			}
		case message := <-h.broadcast: //Hubのbroadcastというchannelにmessage(byte)が入っているとき
			for client := range h.clients { //登録されているclient全員に対して
				select {
				case client.Send <- message: //messageを送ることができれば送る
				default: //送ることができなければ
					close(client.Send)        //channelをcloseして
					delete(h.clients, client) //Hubからdeleteする
				}
			}
		case <-h.stop: //stopがcloseした場合，forループを抜ける
			log.Print("stop recieved")
			return
		}
	}
}

func (h *Hub) Stop() {
	close(h.stop)
	<-h.done
}
