package interceptor

import "github.com/labstack/echo/v4"

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Set(key string, val interface{}) {
	c.Context.Set(key, val)
}
