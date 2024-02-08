package main

import (
	"context"
	"fmt"
	"homework/config"
	"homework/controller"
	"homework/db"
	"homework/gateways/persistence"
	"homework/router"
	"homework/usecase"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		fmt.Println(err)
	}
	db := db.NewDB(*cfg)
	googleOauthConfig := cfg.NewGoogleOauthConfig()
	userRepository := persistence.NewUser()
	mailRepository := persistence.NewMail()
	chatRepository := persistence.NewChatRepository()
	roomRepository := persistence.NewRoom()
	userUsecase := usecase.NewUserUsecase(userRepository, mailRepository, db)
	chatUseCase := usecase.NewChatUsecase(chatRepository, db)
	roomUseCase := usecase.NewRoomUsecase(roomRepository, db)
	userController := controller.NewUserController(userUsecase, userRepository, *cfg, googleOauthConfig, db)
	chatController := controller.NewChatController(chatUseCase, *cfg, googleOauthConfig)
	roomController := controller.NewRoomController(roomUseCase, *cfg, db)
	e := router.NewRouter(userController, chatController, roomController, *cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
