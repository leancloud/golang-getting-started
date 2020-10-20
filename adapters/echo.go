package adapters

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leancloud/go-sdk/leancloud/engine"
)

func Echo(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		MaxAge:       86400,
		AllowHeaders: []string{
			"Content-Type",
			"X-AVOSCloud-Application-Id",
			"X-AVOSCloud-Application-Key",
			"X-AVOSCloud-Application-Production",
			"X-AVOSCloud-Client-Version",
			"X-AVOSCloud-Request-Sign",
			"X-AVOSCloud-Session-Token",
			"X-AVOSCloud-Super-Key",
			"X-LC-Hook-Key",
			"X-LC-Id",
			"X-LC-Key",
			"X-LC-Prod",
			"X-LC-Session",
			"X-LC-Sign",
			"X-LC-UA",
			"X-Requested-With",
			"X-Uluru-Application-Id",
			"X-Uluru-Application-Key",
			"X-Uluru-Application-Production",
			"X-Uluru-Client-Version",
			"X-Uluru-Session-Token",
		},
	}), func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("origin"))
			return next(c)
		}
	})

	e.Group("/1.1", echo.WrapMiddleware(engine.Handler))
}
