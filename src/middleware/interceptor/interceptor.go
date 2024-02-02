package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTAuthMiddleware は、JWTによる認証を行うEchoのミドルウェアです。
func JWTAuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// トークンをヘッダーから取得
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, "Authorization header is missing")
			}

			// "Bearer "のプレフィックスをチェック
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, "Invalid token format")
			}

			// トークンを検証
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid signing method")
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, "Invalid or expired token")
			}

			// トークンからユーザー情報を取得してリクエストコンテキストに設定
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("user", claims["user_id"])
			} else {
				return c.JSON(http.StatusUnauthorized, "Invalid token claims")
			}

			return next(c)
		}
	}
}
