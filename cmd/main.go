package main

import (
	"get-recipe-inator/handler"
	"get-recipe-inator/middleware"

	"github.com/labstack/echo/v4"
	// "get-recipe-inator/handler"
)

func main() {
	app := echo.New()
	app.Static("/public", "public")
	pageHandler := handler.PageHandler{}
	app.Use(middleware.WithUser)
	app.GET("/", pageHandler.HandleDashShow)

	app.Start(":3000")
}
