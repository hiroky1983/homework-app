package main

import (
	"context"
	"fmt"
	"homework/config"
	"homework/controller"
	"homework/db"
	"homework/domain/repository"
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
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase, *cfg)
	e := router.NewRouter(userController, *cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
