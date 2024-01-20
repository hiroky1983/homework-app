package controller

import (
	"encoding/json"
	"fmt"
	"homework/config"
	"homework/usecase"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"golang.org/x/oauth2"
)

type IChatController interface {
	HandleWebSocket(c echo.Context) error
}

type chatController struct {
	cu        usecase.IChatUsecase
	cnf       config.Config
	oauthConf *oauth2.Config
}

func NewChatController(uu usecase.IChatUsecase, cnf config.Config, oauthConf *oauth2.Config) IChatController {
	return &chatController{uu, cnf, oauthConf}
}

type Chat struct {
	ID  int    `json:"id"`
	Message string `json:"message"`
	Sender string `json:"sender"`
	CreatedAt time.Time `json:"created_at"`
}

func (cc *chatController)HandleWebSocket(c echo.Context) error {
	log.Println("Serving at web socket...")
	websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			for {
				// Client からのメッセージを読み込む				
				msg := ""
				err := websocket.Message.Receive(ws, &msg)
				if msg == "" {
					return
				}
				res := &Chat{
					ID: 4,
					Message: msg,
					Sender: "me",
					CreatedAt: time.Now(),
				}
				if err != nil {
					if err.Error() == "EOF" {
						c.Logger().Error(err)
						break
					}
					log.Println(fmt.Errorf("read %s", err))	
					c.Logger().Error(err)
				}

				r ,err := json.Marshal(res)
				if err != nil {
					fmt.Println(err)
					c.Logger().Error(err)
				}

			// 	// Client からのメッセージを元に返すメッセージを作成し送信する
				err = websocket.Message.Send(ws, string(r))
				if err != nil {
					fmt.Println(err)
					c.Logger().Error(err)
				}
			}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
