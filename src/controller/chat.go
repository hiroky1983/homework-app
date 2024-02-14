package controller

import (
	"encoding/json"
	"fmt"
	"homework/config"
	"homework/domain/model/chat"
	apperror "homework/error"
	"homework/middleware/token"
	"homework/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"golang.org/x/oauth2"
)

type IChatController interface {
	HandleWebSocket(c echo.Context) error
	ListChat(c echo.Context) error
	DeleteChat(c echo.Context) error
}

type chatController struct {
	cu        usecase.IChatUsecase
	cnf       config.Config
	oauthConf *oauth2.Config
}

// var (
// 	upgrader = websocket.Upgrader{}
// )

func NewChatController(uu usecase.IChatUsecase, cnf config.Config, oauthConf *oauth2.Config) IChatController {
	return &chatController{uu, cnf, oauthConf}
}

func (cc *chatController) HandleWebSocket(c echo.Context) error {
	log.Println("Serving at web socket...")
	userID, err := token.GetUserIDWithTokenCheck(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, apperror.ErrorWrapperWithCode(err, http.StatusUnauthorized))
	}
	// ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	// if err != nil {
	// 	fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
	// 	return err
	// }
	// defer ws.Close()

	// for {
	// 	// Write
	// 	message := []byte("Hello, Client!")
	// 	err := ws.WriteMessage(websocket.TextMessage, message)
	// 	if err != nil {
	// 		c.Logger().Error(err)
	// 	}

	// 	req := chat.Chat{
	// 		Message: "Hello, Server!",
	// 		UserID:  userID.(string),
	// 	}

	// 	res, err := cc.cu.Create(req)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		c.Logger().Error(err)
	// 	}

	// 	r, err := json.Marshal(res)
	// 	fmt.Println(r)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			c.Logger().Error(err)
	// 		}
	// 	// Read
	// 	a, msg, err := ws.ReadMessage()
	// 	fmt.Println(a)
	// 	if err != nil {
	// 		c.Logger().Error(err)
	// 	}
	// 	fmt.Printf("%s\n", msg)
	// }

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Client からのメッセージを読み込む
			msg := ""
			if err := websocket.Message.Receive(ws, &msg); err != nil {
				fmt.Println(err)
				c.Logger().Error(err)
				break // エラー発生時にループを終了
			}

			// string to json
			var chatreq chat.ChatRequest
			if err := json.Unmarshal([]byte(msg), &chatreq); err != nil {
				fmt.Println(err)
				c.Logger().Error(err)
			}
			req := chat.Chat{
				Message: chatreq.Message,
				RoomID:  chatreq.RoomID,
				UserID:  userID,
			}

			res, err := cc.cu.Create(req)
			if err != nil {
				fmt.Println(err)
				c.Logger().Error(err)
			}

			r, err := json.Marshal(res)
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

func (cc *chatController) ListChat(c echo.Context) error {
	userID, err := token.GetUserIDWithTokenCheck(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, apperror.ErrorWrapperWithCode(err, http.StatusUnauthorized))
	}

	roomID := c.Param("room_id")

	res, err := cc.cu.List(userID, roomID)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, res)
}

func (cc *chatController) DeleteChat(c echo.Context) error {
	req := chat.DeleteChatRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(500, err)
	}

	if err := cc.cu.Delete(req.ID); err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, "success")
}
