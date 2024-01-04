package main

import (
	"homework/controller"
	"homework/db"
	"homework/domain/repository"
	"homework/router"
	"homework/usecase"
	"homework/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}