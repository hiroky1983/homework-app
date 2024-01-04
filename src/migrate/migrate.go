package main

import (
	"fmt"
	"homework/db"
	"homework/domain/model/user"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	dbConn.AutoMigrate(&user.User{})
}