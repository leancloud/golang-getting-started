package adapters

import (
	"github.com/labstack/echo/v4"
	"github.com/leancloud/go-sdk/leancloud"
)

func Echo(e *echo.Echo) {
	e.Any("/1/*", echo.WrapHandler(leancloud.Handler(nil)))
	e.Any("/1.1/*", echo.WrapHandler(leancloud.Handler(nil)))
	e.Any("/__engine/*", echo.WrapHandler(leancloud.Handler(nil)))
}
