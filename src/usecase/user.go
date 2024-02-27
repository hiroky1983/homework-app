package usecase

import (
	"database/sql"
	"errors"
	"homework/config"
	userModel "homework/domain/model/user"
	"homework/domain/repository"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user userModel.User, conf config.Config) (userModel.UserResponse, string, error)
	Login(user userModel.User, conf config.Config) (string, error)
	LoginWithGoogle(user userModel.User, cnf config.Config) (string, error)
	CreateProfile(user userModel.User) error
	GetProfile(userID string) (userModel.UserProfileResponse, error)
	List(userID string) (userModel.Users, error)
	Authorization(userID string) (userModel.AuthorizationResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	mu repository.Mail
	db *bun.DB
}

func NewUserUsecase(ur repository.IUserRepository, mr repository.Mail, db *bun.DB) IUserUsecase {
	return &userUsecase{ur, mr, db}
}

func (uu *userUsecase) SignUp(user userModel.User, cnf config.Config) (userModel.UserResponse, string, error) {
	if err := user.ValidateUser(); err != nil {
		return userModel.UserResponse{}, "", err
	}
	storedUser := userModel.User{}
	if err := uu.ur.GetUserByEmail(uu.db, &storedUser, user.Email); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return userModel.UserResponse{}, "", err
		}
	}
	if storedUser.Email != "" {
		return userModel.UserResponse{}, "", errors.New("email already exists")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return userModel.UserResponse{}, "", err
	}
	newUser := userModel.User{
		Email:    user.Email,
		Password: string(hash),
	}
	tx, err := uu.db.Begin()
	if err != nil {
		return userModel.UserResponse{}, "", err
	}

	if err := uu.ur.CreateUser(tx, &newUser); err != nil {
		tx.Rollback()
		return userModel.UserResponse{}, "", err
	}

	tokenString, err := newUser.GenerateToken(cnf)
	if err != nil {
		tx.Rollback()
		return userModel.UserResponse{}, "", err
	}

	if err := uu.mu.SendMail(newUser.Email, tokenString, cnf); err != nil {
		tx.Rollback()
		return userModel.UserResponse{}, "", err
	}
	tx.Commit()
	res := userModel.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return res, tokenString, nil
}

func (uu *userUsecase) Login(user userModel.User, cnf config.Config) (string, error) {
	if err := user.ValidateUser(); err != nil {
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

	tokenString, err := storedUser.GenerateToken(cnf)
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

	tokenString, err := storedUser.GenerateToken(cnf)
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

func (uu *userUsecase) GetProfile(userID string) (userModel.UserProfileResponse, error) {
	user := userModel.User{}
	if err := uu.ur.GetProfile(uu.db, &user, userID); err != nil {
		return userModel.UserProfileResponse{}, err
	}

	res := user.NewUserProfileResponse()

	return res, nil
}

func (uu *userUsecase) List(userID string) (userModel.Users, error) {
	users, err := uu.ur.ListUser(uu.db, userID)
	if err != nil {
		return nil, err
	}

	res := users.NewUserListResponse()

	return res, nil
}

func (uu *userUsecase) Authorization(userID string) (userModel.AuthorizationResponse, error) {
	user := userModel.User{}
	ok, err := uu.ur.IsExistUser(uu.db, &user, userID)
	if err != nil {
		return userModel.AuthorizationResponse{}, err
	}
	if !ok {
		return userModel.AuthorizationResponse{}, errors.New("user not found")
	}

	res := user.NewAuthResponse()
	if !res.IsVerified {
		return userModel.AuthorizationResponse{}, errors.New("still user not signup verified")
	}
	if res.IsDeleted {
		return userModel.AuthorizationResponse{}, errors.New("user is deleted")
	}

	return res, nil
}
