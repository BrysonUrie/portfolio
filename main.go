package main

import (
	"github.com/a-h/templ"
	"github.com/brysonurie/portfolio/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/static", "assets")
	e.GET("/", homeHandler)
	e.Logger.Fatal(e.Start(":8000"))
}

func homeHandler(c echo.Context) error {
	return render(c, views.Hello("Bryson Urie", "Bryson"))
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}
