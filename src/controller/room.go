package controller

import (
	"homework/config"
	"homework/usecase"

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
	return nil
}
