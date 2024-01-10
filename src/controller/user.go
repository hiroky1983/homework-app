package controller

import (
	"homework/config"
	"homework/domain/model/user"
	apperror "homework/error"
	"homework/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu  usecase.IUserUsecase
	cnf config.Config
}

func NewUserController(uu usecase.IUserUsecase, cnf config.Config) IUserController {
	return &userController{uu, cnf}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := user.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, apperror.ErrorWrapperWithCode(err, http.StatusBadRequest))
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := user.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, apperror.ErrorWrapperWithCode(err, http.StatusBadRequest))
	}
	tokenString, err := uc.uu.Login(user, uc.cnf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, apperror.ErrorWrapperWithCode(err, http.StatusInternalServerError))
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"                           // cookieにセットするkey名を定義
	cookie.Value = tokenString                      // cookieにセットするvalueを定義
	cookie.Expires = time.Now().Add(24 * time.Hour) // cookieの有効期限を定義(24時間に設定)
	cookie.Path = "/"                               // cookieの有効パスを定義
	cookie.Domain = uc.cnf.APIDomain                //  cookieの有効ドメインを定義
	cookie.Secure = true                            // cookieのHTTPS通信のみ有効にする（postmanやlocalhostで試す場合はfalseにする）
	cookie.HttpOnly = true                          // cookieをHTTP通信のみ有効にする（JSからのアクセスを禁止する）
	cookie.SameSite = http.SameSiteNoneMode         // cookieをサイト間で共有する（クロスサイトリクエストを許可する）
	c.SetCookie(cookie)                             // cookieをセット
	return c.JSON(http.StatusOK, "{ status: 200 , message: success}")
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token" // cookie名を定義
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = uc.cnf.APIDomain
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
