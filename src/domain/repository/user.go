package repository

import (
	"homework/domain/model/user"
)

//go:generate moq -pkg fakerepository -out ../../moq/fakerepository/user_test_moq.go . IUserRepository
type IUserRepository interface {
	GetUserByEmail(db DBConn, user *user.User, email string) error
	CreateUser(db DBConn, user *user.User) error
	GetUserByID(db DBConn, user *user.User, userID string) error
}
