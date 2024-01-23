package repository

import (
	"context"
	"database/sql"
	"homework/domain/model/user"
)

//go:generate moq -pkg fakerepository -out ../../moq/fakerepository/user_test_moq.go . IUserRepository

type IUserRepository interface {
	GetUserByEmail(db DBConn, user *user.User, email string) error
	CreateUser(db DBConn, user *user.User) error
	GetUserByID(db DBConn, user *user.User, userID string) error
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUserByEmail(db DBConn, user *user.User, email string) error {
	if err := db.NewSelect().Model((user)).Where("email=?", email).Where("password != '' ").Scan(context.Background(), user); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(db DBConn, user *user.User) error {
	_, err := db.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByID(db DBConn, user *user.User, googleID string) error {
	if err := db.NewSelect().Model((user)).Where("google_id=?", googleID).Scan(context.Background(), user); err != nil {
		if sql.ErrNoRows == err {
			return nil
		}
		return err
	}
	return nil
}
