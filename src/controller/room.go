package controller

import (
	"homework/config"
	"homework/domain/model/room"
	apperror "homework/error"
	"homework/middleware/token"
	"homework/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type IRoomController interface {
	CreateRoom(c echo.Context) error
}

type roomController struct {
	ru  usecase.IRoomUsecase
	cnf config.Config
	db  *bun.DB
}

func NewRoomController(ru usecase.IRoomUsecase, cnf config.Config, db *bun.DB) IRoomController {
	return &roomController{ru, cnf, db}
}

// CreateRoom godoc
//
// @Summary      ルーム作成API
// @Description  ルーム作成
// @Accept       json
// @Produce      json
// @Param        body body room.RoomRequest  false  "ルームID"
// @Success      200 {object} room.RoomResponse
// @Router       /room/create [post]
func (rc *roomController) CreateRoom(c echo.Context) error {
	userID, err := token.GetUserIDWithTokenCheck(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, apperror.ErrorWrapperWithCode(err, http.StatusUnauthorized))
	}
	otherUser := room.RoomRequest{}
	if err := c.Bind(&otherUser); err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}

	Room, err := rc.ru.Create(userID, room.RoomMap{
		UserID: otherUser.UserID,
		RoomID: otherUser.RoomID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, Room)
}
