package main

import (
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// sets the maximum allowed size for a request body.
	e.Use(middleware.BodyLimit("2M"))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
