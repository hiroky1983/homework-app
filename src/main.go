package main

import (
	"context"
	"flag"
	"homework/config"
	"homework/controller"
	"homework/db"
	_ "homework/docs"
	"homework/gateways/persistence"
	"homework/router"
	"homework/usecase"
	"homework/websocket"
	"log"
)

// @title         homework API
// @version       1.0
// @license.name  ymd
// @BasePath      /
func main() {
	ctx := context.Background()
	flag.Parse()
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Println(err)
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
	hub := websocket.NewHub(chatUseCase)
	go hub.Run()
	webSocketController := controller.NewWebSocketController(hub, chatUseCase)
	e := router.NewRouter(userController, chatController, roomController, webSocketController, *cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
