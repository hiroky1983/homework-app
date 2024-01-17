package router

import (
	"fmt"
	"homework/config"
	"homework/controller"
	apperror "homework/error"
	"log"
	"net/http"
	"runtime/debug"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func NewRouter(uc controller.IUserController, cnf config.Config) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", cnf.AppURL},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   cnf.APIDomain,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	}))
	e.Use(middleware.Logger())
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	e.POST("/google", uc.GoogleAuth)
	e.GET("/google/callback", uc.GoogleAuthCallback)
	user := e.Group("/user")
	user.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(cnf.Seclet),
		TokenLookup: "cookie:token",
	}))
	user.GET("", uc.GetUser)
	e.GET("/socket", handleWebSocket)

	e.HTTPErrorHandler = customHTTPErrorHandler
	return e
}

// エラーハンドリング関数
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "Internal Server Error"

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message.(string)
	}

	// スタックトレースを取得
	trace := string(debug.Stack())

	// コンテナログにエラー情報を出力
	logrus.WithFields(logrus.Fields{
		"error": err,
		"trace": trace,
	}).Error("An error occurred")

	// レスポンスにはスタックトレースを含めない
	customErr := &apperror.ErrorResponse{
		Code:    code,
		Message: msg,
	}

	// レスポンスを送信
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD { // HEADリクエストの場合はボディを含めない
			c.NoContent(code)
		} else {
			c.JSON(code, customErr)
		}
	}
}

func handleWebSocket(c echo.Context) error {
	log.Println("Serving at localhost:8000...")
	websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()
			fmt.Println("===============================================")
			// 初回のメッセージを送信
			err := websocket.Message.Send(ws, "Server: Hello, Next.js!")
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Println("===============================================")
			for {
				// Client からのメッセージを読み込む
				fmt.Println("===============================================")
				msg := ""
				err := websocket.Message.Receive(ws, &msg)
				if err != nil {
					if err.Error() == "EOF" {
						fmt.Println("===============================================")
						fmt.Println(err)
						c.Logger().Error(err)
						break
					}
					log.Println(fmt.Errorf("read %s", err))	
					c.Logger().Error(err)
				}

				// Client からのメッセージを元に返すメッセージを作成し送信する
				err = websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))
				if err != nil {

					fmt.Println("===============================================")
					fmt.Println(err)
					c.Logger().Error(err)
				}
			}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
