package routes

import (
	"github.com/labstack/echo/v4"
	"ofspace-be/factory"
)

func New() *echo.Echo {
	presenter := factory.New()
	e := echo.New()
	// accessibility
	e.POST("/accessibility", presenter.AccessibilityPresentation.CreateAccessibility)
	e.GET("/accessibility/:accessibility_id", presenter.AccessibilityPresentation.GetAccessibility)
	e.POST("/accessibility/:accessibility_id", presenter.AccessibilityPresentation.RequestAccessibility)
	e.GET("/accessibility/search", presenter.AccessibilityPresentation.SearchAccessibility)
	e.PUT("/accessibility/:accessibility_id", presenter.AccessibilityPresentation.UpdateAccessibility)
	return e
}
