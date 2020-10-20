package adapters

import (
	"github.com/labstack/echo/v4"
	"github.com/leancloud/go-sdk/leancloud/engine"
)

func Echo(e *echo.Echo) {
	e.Group("/1.1", echo.WrapMiddleware(engine.Handler))
}
