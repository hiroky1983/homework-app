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
	userUsecase := usecase.NewUserUsecase(userRepository, mailRepository, db)
	userController := controller.NewUserController(userUsecase, userRepository, *cfg, googleOauthConfig, db)
	chatRepository := persistence.NewChatRepository()
	chatUseCase := usecase.NewChatUsecase(chatRepository, db)
	chatController := controller.NewChatController(chatUseCase, *cfg, googleOauthConfig)
	e := router.NewRouter(userController, chatController, *cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
