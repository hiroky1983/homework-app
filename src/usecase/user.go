package usecase

import (
	"fmt"
	"homework/config"
	userModel "homework/domain/model/user"
	"homework/domain/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user userModel.User) (userModel.UserResponse, error)
	Login(user userModel.User, conf config.Config) (string, error)
	LoginWithGoogle(user userModel.User, cnf config.Config) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user userModel.User) (userModel.UserResponse, error) {
	if err := user.Validate(); err != nil {
		fmt.Println(err)
		return userModel.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return userModel.UserResponse{}, err
	}
	newUser := userModel.User{
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
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
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
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
	if err := uu.ur.GetUserByID(&storedUser, user.GoogleID); err != nil {
		return "", err
	}
	if storedUser.GoogleID == "" {
		if err := uu.ur.CreateUser(&user); err != nil {
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
		return "", err
	}
	return tokenString, nil
}
