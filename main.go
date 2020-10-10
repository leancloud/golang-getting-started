package main

import (
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leancloud/go-sdk/leancloud/engine"
	"github.com/leancloud/golang-getting-started/middlewares"
	"github.com/leancloud/golang-getting-started/routes"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echo.WrapMiddleware(engine.Handler))
	e.Use(middlewares.Echo)
	e.Renderer = t
	e.Static("/assets", "./assets")

	e.GET("/", routes.Index)
	e.GET("/todos", routes.GetTodos)
	e.POST("/todos", routes.PostTodos)

	e.Logger.Fatal(e.Start(":" + port))

}
