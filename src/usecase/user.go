package usecase

import (
	"database/sql"
	"errors"
	"homework/config"
	userModel "homework/domain/model/user"
	"homework/domain/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user userModel.User, conf config.Config) (userModel.UserResponse, error)
	Login(user userModel.User, conf config.Config) (string, error)
	LoginWithGoogle(user userModel.User, cnf config.Config) (string, error)
	CreateProfile(user userModel.User) error
}

type userUsecase struct {
	ur repository.IUserRepository
	mu			repository.Mail
	db *bun.DB
}

func NewUserUsecase(ur repository.IUserRepository, mr repository.Mail ,db *bun.DB) IUserUsecase {
	return &userUsecase{ur,mr , db}
}

func (uu *userUsecase) SignUp(user userModel.User , cnf config.Config) (userModel.UserResponse, error) {
	if err := user.Validate(); err != nil {
		return userModel.UserResponse{}, err
	}
	storedUser := userModel.User{}
	if err := uu.ur.GetUserByEmail(uu.db, &storedUser, user.Email); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return userModel.UserResponse{}, err
		}
	}
	if storedUser.Email != "" {
		return userModel.UserResponse{}, errors.New("email already exists")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return userModel.UserResponse{}, err
	}
	newUser := userModel.User{
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uu.ur.CreateUser(uu.db, &newUser); err != nil {
		return userModel.UserResponse{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(cnf.Seclet))
	if err := uu.mu.SendMail(newUser.Email, tokenString, cnf); err != nil {
		return userModel.UserResponse{}, err
	}
	resUser := userModel.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user userModel.User, cnf config.Config) (string, error) {
	if err := user.Validate(); err != nil {
		return "", err
	}
	storedUser := userModel.User{}
	if err := uu.ur.GetUserByEmail(uu.db, &storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(cnf.Seclet))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *userUsecase) LoginWithGoogle(user userModel.User, cnf config.Config) (string, error) {
	storedUser := userModel.User{}
	if err := uu.ur.GetUserByID(uu.db, &storedUser, user.GoogleID); err != nil {
		return "", err
	}
	tx, err := uu.db.Begin()
	if storedUser.GoogleID == "" {
		if err != nil {
			return "", err
		}
		if err := uu.ur.CreateUser(tx, &user); err != nil {
			tx.Rollback()
			return "", err
		}
		storedUser.ID = user.ID
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(cnf.Seclet))
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return tokenString, nil
}

func (uu *userUsecase) CreateProfile(user userModel.User) error {
	if err := uu.ur.UpdateUser(uu.db, &user); err != nil {
		return err
	}
	return nil
}
