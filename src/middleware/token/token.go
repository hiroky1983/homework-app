package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetUserIDWithTokenCheck(c echo.Context) (string, error) {
	u := c.Get("user").(*jwt.Token)
	claims, ok := u.Claims.(jwt.MapClaims)
	if !ok || !u.Valid {
		return "", fmt.Errorf("invalid token")
	}

	exp := claims.VerifyExpiresAt(time.Now().Unix(), true)
	if !exp {
		return "", fmt.Errorf("token is expired")
	}

	return claims["user_id"].(string), nil
}

func QueryTokenCheck(c echo.Context) error {
	t := c.QueryParam("token")
	if t == "" {
		return fmt.Errorf("token is required")
	}
	u := c.Get("user").(*jwt.Token)
	if u.Raw != t {
		return fmt.Errorf("invalid token")
	}
	return nil
}
