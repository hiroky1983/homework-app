package controller

import (
	"encoding/json"
	"fmt"
	"homework/config"
	"homework/domain/model/chat"
	"homework/usecase"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func NewChatController(uu usecase.IChatUsecase, cnf config.Config, oauthConf *oauth2.Config) IChatController {
	return &chatController{uu, cnf, oauthConf}
}

type Chat struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	Sender    string    `json:"sender"`
	CreatedAt time.Time `json:"created_at"`
}

func (cc *chatController) HandleWebSocket(c echo.Context) error {
	log.Println("Serving at web socket...")
	websocket.Handler(func(ws *websocket.Conn) {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["user_id"]
		defer ws.Close()
		for {
			// Client からのメッセージを読み込む
			msg := ""
			if err := websocket.Message.Receive(ws, &msg); err != nil {
				fmt.Println(err)
				c.Logger().Error(err)
				break // エラー発生時にループを終了
			}

			req := chat.Chat{
				Message: msg,
				UserID:  userID.(string),
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]

	res, err := cc.cu.List(userID.(string))
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
	fmt.Println(req)
	if err := cc.cu.Delete(req.ID); err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, "success")
}
