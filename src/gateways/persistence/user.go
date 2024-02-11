package persistence

import (
	"context"
	"database/sql"
	"homework/domain/model/user"
	"homework/domain/repository"
)

type User struct{}

func NewUser() *User {
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

func (ur *User) UpdateUser(db repository.DBConn, u *user.User) error {
	_, err := db.NewUpdate().Model(u).Set("user_name = ?", u.UserName).Set("updated_at = NOW()").Where("id=?", u.ID).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (ur *User) UpdateIsVerifiedUser(db repository.DBConn, userID string) error {
	_, err := db.NewUpdate().Model(&user.User{}).Set("is_verified = ?", true).Where("id = ?", userID).Exec(context.Background())
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

func (ur *User) GetProfile(db repository.DBConn, u *user.User, userID string) error {
	if err := db.NewSelect().Model((u)).Where("id=?", userID).Scan(context.Background()); err != nil {
		return err
	}
	return nil
}

func (ur *User) ListUser(db repository.DBConn, u *user.Users, userID string) error {
	if err := db.NewSelect().Model(u).Where("id!=?", userID).Scan(context.Background()); err != nil {
		return err
	}
	return nil
}
