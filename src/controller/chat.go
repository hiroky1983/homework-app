package controller

import (
	"homework/config"
	"homework/domain/model/chat"
	apperror "homework/error"
	"homework/middleware/token"
	"homework/usecase"
	"net/http"

	"github.com/labstack/echo/v4"

	"golang.org/x/oauth2"
)

type IChatController interface {
	CreateChat(c echo.Context) error
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

// CreateChat godoc
//
// @Summary      チャット保存API
// @Description  チャット保存
// @Accept       json
// @Produce      json
// @Param        body  body    chat.CreateChatRequest  false  "チャット保存"
// @Success      200  {object}  chat.ChatResponse
// @Router       /chat/create [post]
func (cc *chatController) CreateChat(c echo.Context) error {
	userID, err := token.GetUserIDWithTokenCheck(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, apperror.ErrorWrapperWithCode(err, http.StatusUnauthorized))
	}

	req := chat.CreateChatRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(500, err)
	}

	chat := chat.Chat{
		Message: req.Message,
		RoomID:  req.RoomID,
		UserID:  userID,
	}

	res, err := cc.cu.Create(chat)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, res)
}

// ListChat godoc
//
// @Summary      チャット取得API
// @Description  ルーム内のチャットを取得
// @Accept       json
// @Produce      json
// @Param        room_id path string  false  "ルームID"
// @Success      200  {array}  chat.ChatResponse
// @Router       /chat/get/{room_id} [get]
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

// DeleteChat godoc
//
// @Summary      チャット削除API
// @Description  ルーム内のチャットを削除
// @Accept       json
// @Produce      json
// @Param        body body chat.DeleteChatRequest  false  "ルームID"
// @Success      200 {object} string 
// @Router       /chat/delete [put]
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
