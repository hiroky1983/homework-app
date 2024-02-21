package router

import (
	"homework/config"
	"homework/controller"
	apperror "homework/error"
	"net/http"
	"os"
	"runtime/debug"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func NewRouter(uc controller.IUserController, cc controller.IChatController, rc controller.IRoomController, wc controller.IWebSocketController, cnf config.Config) *echo.Echo {
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
	e.Use(middleware.Recover())
	// =============== user ===============
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
	user.GET("/confirm", uc.SignUpCallback)
	user.GET("/get_profile", uc.GetUser)
	user.POST("/create_profile", uc.CreateProfile)
	user.GET("/users", uc.ListUser)
	// =============== chat ===============
	chat := e.Group("/chat")
	chat.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	chat.POST("/create", cc.CreateChat)
	chat.GET("/get/:room_id", cc.ListChat)
	chat.PUT("/delete", cc.DeleteChat)
	// =============== room ===============
	room := e.Group("/room")
	room.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	room.POST("/create", rc.CreateRoom)

	// ================ websocket ================
	ws := e.Group("/socket")
	ws.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	ws.GET("/:room_id", wc.ServeRoomWs)

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
