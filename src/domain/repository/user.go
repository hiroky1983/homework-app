package repository

import (
	"context"
	"database/sql"
	"homework/domain/model/user"

	"github.com/uptrace/bun"
)

type IUserRepository interface {
	GetUserByEmail(user *user.User, email string) error
	CreateUser(user *user.User) error
	GetUserByID(user *user.User, userID string) error
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *user.User, email string) error {
	if err := ur.db.NewSelect().Model((user)).Where("email=?", email).Where("password != '' ").Scan(context.Background(), user); err != nil {
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

func (ur *userRepository) GetUserByID(user *user.User, googleID string) error {
	if err := ur.db.NewSelect().Model((user)).Where("google_id=?", googleID).Scan(context.Background(), user); err != nil {
		if sql.ErrNoRows == err {
			return nil
		}
		return err
	}
	return nil
}
