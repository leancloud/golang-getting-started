package adapters

import (
	"github.com/labstack/echo/v4"
	"github.com/leancloud/go-sdk/leancloud/engine"
)

func Echo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		engine.Handler(c.Echo().Server.Handler).ServeHTTP(c.Response().Writer, c.Request())
		return next(c)
	}
}
