package controller

import (
	"homework/config"
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
	userID ,err := token.GetUserIDWithTokenCheck(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, apperror.ErrorWrapperWithCode(err, http.StatusUnauthorized))
	}

	if err := rc.ru.Create(userID);err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, nil)
}
