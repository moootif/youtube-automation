package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "web/static")

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "YouTube Automation Server is Running!")
	})

	// Start server
	// 0.0.0.0を指定することで、同じWi-Fi内のiPhoneからアクセス可能になります
	e.Logger.Fatal(e.Start(":8080"))
}
