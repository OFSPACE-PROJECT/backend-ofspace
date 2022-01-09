package routes

import (
	"ofspace_be/factory"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	presenter := factory.New()
	e := echo.New()

	// users
	e.POST("/register", presenter.UserPresentation.RegisterUser)
	e.POST("/login", presenter.UserPresentation.LoginUser)
	e.GET("/user/:id", presenter.UserPresentation.GetUserByID)
	return e
}
