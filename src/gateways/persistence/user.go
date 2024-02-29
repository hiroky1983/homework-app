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

func (ur *User) ListUser(db repository.DBConn, userID string) (user.Users, error) {
	u := &user.User{}
	var users user.Users
	if err := db.NewSelect().Model(u).
		Column("u.id", "u.user_name", "u.email", "u.image_path").
		ColumnExpr("rm.room_id AS room_id").
		Join("LEFT JOIN room_map AS rm ON u.id = rm.user_id").
		Where("u.id != ?", userID).
		Scan(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *User) IsExistUser(db repository.DBConn, u *user.User, UserID string) (bool, error) {
	if err := db.NewSelect().Model((u)).Where("id=?", UserID).Scan(context.Background(), u); err != nil {
		if sql.ErrNoRows == err {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
