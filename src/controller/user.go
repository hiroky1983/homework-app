package controller

import (
	"context"
	"fmt"
	"homework/config"
	"homework/domain/model/user"
	"homework/domain/repository"
	apperror "homework/error"
	"homework/middleware/cookie"
	"homework/usecase"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	GetUser(c echo.Context) error
	GoogleAuth(c echo.Context) error
	GoogleAuthCallback(c echo.Context) error
	CreateProfile(c echo.Context) error
	SignUpCallback(c echo.Context) error
}

type userController struct {
	uu        usecase.IUserUsecase
	ur        repository.IUserRepository
	cnf       config.Config
	oauthConf *oauth2.Config
	db *bun.DB
}

func NewUserController(uu usecase.IUserUsecase,ur repository.IUserRepository, cnf config.Config, oauthConf *oauth2.Config, 	db *bun.DB) IUserController {
	return &userController{uu, ur, cnf, oauthConf, db}
}

func (uc *userController) SignUp(c echo.Context) error {
	u := user.User{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, apperror.ErrorWrapperWithCode(err, http.StatusBadRequest))
	}
	userRes,tokenString, err := uc.uu.SignUp(u, uc.cnf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}
	cookie.SetCookie(tokenString, uc.cnf.APIDomain, c, time.Now().Add(24*time.Hour))
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	u := user.User{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, apperror.ErrorWrapperWithCode(err, http.StatusBadRequest))
	}
	tokenString, err := uc.uu.Login(u, uc.cnf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}
	cookie.SetCookie(tokenString, uc.cnf.APIDomain, c, time.Now().Add(24*time.Hour))
	return c.JSON(http.StatusOK, user.LonginResponse{
		Code:    http.StatusOK,
		Message: "success",
	})
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie.SetCookie("", uc.cnf.APIDomain, c, time.Now())
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

func (uc *userController) GoogleAuth(c echo.Context) error {
	token := c.Get("csrf").(string)
	url := uc.oauthConf.AuthCodeURL(token, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return c.JSON(http.StatusOK, url)
}

func (uc *userController) GoogleAuthCallback(c echo.Context) error {
	Code := c.QueryParam("code")
	ctx := context.Background()
	tok, err := uc.oauthConf.Exchange(ctx, Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}

	s, err := v2.NewService(ctx, option.WithTokenSource(uc.oauthConf.TokenSource(ctx, tok)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}

	info, err := s.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}
	u := user.User{}
	u.Email = info.Email
	u.GoogleID = info.UserId
	u.IsVerified = true
	var url string
	tokenString, err := uc.uu.LoginWithGoogle(u, uc.cnf)
	if err != nil {
		url = fmt.Sprintf("%s/not-found", uc.cnf.AppURL)
		return c.Redirect(http.StatusFound, url)
	}

	cookie.SetCookie(tokenString, uc.cnf.APIDomain, c, time.Now().Add(24*time.Hour))

	url = fmt.Sprintf("%s/top", uc.cnf.AppURL)
	return c.Redirect(http.StatusFound, url)
}

func (uc *userController) GetUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	return c.JSON(http.StatusOK, userID)
}

func (uc *userController) CreateProfile(c echo.Context) error {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	us := user.UserProfileRequest{}
	if err := c.Bind(&us); err != nil {
		return c.JSON(http.StatusBadRequest, apperror.ErrorWrapperWithCode(err, http.StatusBadRequest))
	}
	user := user.User{}
	user.ID = userID.(string)
	user.UserName = us.UserName
	if err := uc.uu.CreateProfile(user); err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}
	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "success",
	})
}

func (uc *userController) SignUpCallback(c echo.Context) error {
	u := c.Get("user").(*jwt.Token)
	claims, ok := u.Claims.(jwt.MapClaims)
	if !ok || !u.Valid {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(fmt.Errorf("invalid token"), http.StatusInternalServerError))
	}
	userID := claims["user_id"]
	user := user.User{}
	user.ID = userID.(string)

	if err := uc.ur.UpdateIsVerifiedUser(uc.db, user.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}

	return c.Redirect(http.StatusFound, "http://localhost:3000/profile")
}
