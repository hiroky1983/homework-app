package controller

import (
	"homework/config"
	"homework/domain/model/chat"
	apperror "homework/error"
	"homework/middleware/token"
	"homework/usecase"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

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

var (
	upgrader = websocket.Upgrader{	
		CheckOrigin: func(r *http.Request) bool {
		// ここでオリジンの検証を行う場合があります。
		// セキュリティ上の理由からオリジンを検証することが推奨されますが、
		// 開発中や内部ツールの場合は true を返してすべて許可することがあります。
		return true
		},
	}	
)

func NewChatController(uu usecase.IChatUsecase, cnf config.Config, oauthConf *oauth2.Config) IChatController {
	return &chatController{uu, cnf, oauthConf}
}

func (cc *chatController) HandleWebSocket(c echo.Context) error {
	log.Println("Serving at web socket...")

	userID, err := token.GetUserIDWithTokenCheck(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, apperror.ErrorWrapperWithCode(err, http.StatusUnauthorized))
	}

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return err
	}
	defer ws.Close()

	for {
		// Client からのメッセージを読み込む
		var chatreq chat.ChatRequest
		if err := ws.ReadJSON(&chatreq); err != nil {
			log.Println("Read error:", err)
			break // エラー発生時にループを終了
		}

		req := chat.Chat{
			Message: chatreq.Message,
			RoomID:  chatreq.RoomID,
			UserID:  userID,
		}

		res, err := cc.cu.Create(req)
		if err != nil {
			log.Println("Create error:", err)
			continue // エラーが発生した場合は次のメッセージを待つ
		}

		// クライアントにレスポンスを送信する
		if err := ws.WriteJSON(res); err != nil {
			log.Println("Write error:", err)
			break
		}
	}

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
