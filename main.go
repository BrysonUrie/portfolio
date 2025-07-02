package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/brysonurie/portfolio/views"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	ContactForm struct {
		Name    string `form:"name" validate:"required"`
		Email   string `form:"email" validate:"required,email"`
		Phone   string `form:"phone" validate:"required"`
		Message string `form:"message"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Static("/static", "assets")
	e.GET("/", homeHandler)

	e.POST("/api/contact", contactFormHandler)
	e.Logger.Fatal(e.Start(":8000"))
}

func homeHandler(c echo.Context) error {
	return render(c, views.Hello("Bryson Urie", "Bryson"))
}

func contactFormHandler(c echo.Context) (err error) {
	f := new(ContactForm)
	if err = c.Bind(f); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(f); err != nil {
		return err
	}
	fmt.Println(f)
	return render(c, views.Hello("Bryson Urie", "Bryson"))
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}
