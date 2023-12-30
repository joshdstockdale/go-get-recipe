package main

import (
	"get-recipe-inator/handler"
	"get-recipe-inator/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/public", "public")
	urlHandler := handler.UrlHandler{}
	app.Use(middleware.WithUser)
	app.GET("/", urlHandler.HandleHome)
	app.GET("/list", urlHandler.HandleList)
	app.GET("/detail", urlHandler.HandleDetail)

	app.Start(":3000")
}
