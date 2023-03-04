package main

import (
	"password_generator/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.GetAllRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
