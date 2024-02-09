package controller

import (
	"homework/config"
	"homework/usecase"

	"github.com/golang-jwt/jwt/v4"
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
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	userID := claims["user_id"]

	err := rc.ru.Create(userID.(string))
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, nil)
}
