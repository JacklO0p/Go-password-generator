package routes

import (
	"password_generator/controllers"

	"github.com/labstack/echo/v4"
)

func GetAllRoutes(e *echo.Echo) {
	e.Static("/", "static")

	e.GET("/generate", controllers.Home)
}
