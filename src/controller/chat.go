package controller

import (
	"fmt"
	"homework/config"
	"homework/usecase"
	"log"

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


func (cc *chatController)HandleWebSocket(c echo.Context) error {
	log.Println("Serving at web socket...")
	websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			// 初回のメッセージを送信
			err := websocket.Message.Send(ws, "Server: Hello, Next.js!")
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Println("====================2===========================")
			for {
				// Client からのメッセージを読み込む
				fmt.Println("==================3=============================")
				msg := ""
				err := websocket.Message.Receive(ws, &msg)
				if err != nil {
					if err.Error() == "EOF" {
						fmt.Println("==================4=============================")
						c.Logger().Error(err)
						break
					}
					log.Println(fmt.Errorf("read %s", err))	
					c.Logger().Error(err)
				}

			// 	// Client からのメッセージを元に返すメッセージを作成し送信する
				err = websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))
				if err != nil {

					fmt.Println("====================5===========================")
					fmt.Println(err)
					c.Logger().Error(err)
				}
			}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
