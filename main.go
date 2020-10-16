package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leancloud/golang-getting-started/adapters"
	_ "github.com/leancloud/golang-getting-started/functions"
	"github.com/leancloud/golang-getting-started/routes"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	port := os.Getenv("LEANCLOUD_APP_PORT")
	if port == "" {
		port = "3000"
	}

	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = t
	e.Static("/assets", "./assets")

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if herr, ok := err.(*echo.HTTPError); ok {
			code = herr.Code
		}

		c.Logger().Error(err)
		c.Render(http.StatusInternalServerError, "error", struct {
			Message string
			Status  int
			Error   string
		}{
			Message: err.Error(),
			Status:  code,
			Error:   err.Error(),
		})
	}

	e.GET("/", routes.Index)
	e.GET("/todos", routes.GetTodos)
	e.POST("/todos", routes.PostTodos)
	adapters.Echo(e)

	e.Logger.Fatal(e.Start(":" + port))
}
