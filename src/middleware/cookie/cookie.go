package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetCookie(tokenString string, APIDomain string, c echo.Context, t time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = "token"                   // cookieにセットするkey名を定義
	cookie.Value = tokenString              // cookieにセットするvalueを定義
	cookie.Expires = t                      // cookieの有効期限を定義(24時間に設定)
	cookie.Path = "/"                       // cookieの有効パスを定義
	cookie.Domain = APIDomain               //  cookieの有効ドメインを定義
	cookie.Secure = true                    // cookieのHTTPS通信のみ有効にする（postmanやlocalhostで試す場合はfalseにする）
	cookie.HttpOnly = true                  // cookieをHTTP通信のみ有効にする（JSからのアクセスを禁止する）
	cookie.SameSite = http.SameSiteNoneMode // cookieをサイト間で共有する（クロスサイトリクエストを許可する）
	c.SetCookie(cookie)                     // cookieをセット
}
