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

func (rc *roomController) CreateRoom(c echo.Context) error {
	userID, err := token.GetUserIDWithTokenCheck(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, apperror.ErrorWrapperWithCode(err, http.StatusUnauthorized))
	}
	otherUser := room.RoomMap{}
	if err := c.Bind(&otherUser); err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}

	RoomID, err := rc.ru.Create(userID, otherUser)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]string{"roomId": RoomID})
}
