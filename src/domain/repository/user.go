package repository

import (
	"context"
	"homework/domain/model/user"

	"github.com/uptrace/bun"
)

type IUserRepository interface {
	GetUserByEmail(user *user.User, email string) error
	CreateUser(user *user.User) error
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *user.User, email string) error {
	if err := ur.db.NewSelect().Model((user)).Where("email=?", email).Scan(context.Background(), user); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *user.User) error {
	_, err := ur.db.NewInsert().Model(user).Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}
