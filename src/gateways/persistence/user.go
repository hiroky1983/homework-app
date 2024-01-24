package persistence

import (
	"context"
	"database/sql"
	"homework/domain/model/user"
	"homework/domain/repository"
)

type User struct{}

func NewUserRepository() *User {
	return &User{}
}

func (ur *User) GetUserByEmail(db repository.DBConn, u *user.User, email string) error {
	if err := db.NewSelect().Model((u)).Where("email=?", email).Where("password != '' ").Scan(context.Background(), u); err != nil {
		return err
	}
	return nil
}

func (ur *User) CreateUser(db repository.DBConn, u *user.User) error {
	_, err := db.NewInsert().Model(u).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (ur *User) GetUserByID(db repository.DBConn, u *user.User, googleID string) error {
	if err := db.NewSelect().Model((u)).Where("google_id=?", googleID).Scan(context.Background(), u); err != nil {
		if sql.ErrNoRows == err {
			return nil
		}
		return err
	}
	return nil
}
