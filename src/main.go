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
	userRepository := persistence.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository, db)
	userController := controller.NewUserController(userUsecase, *cfg, googleOauthConfig)
	chatRepository := persistence.NewChatRepository()
	chatUseCase := usecase.NewChatUsecase(chatRepository, db)
	chatController := controller.NewChatController(chatUseCase, *cfg, googleOauthConfig)
	e := router.NewRouter(userController, chatController,*cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
