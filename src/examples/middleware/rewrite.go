package main

import (
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// rewrites the URL path based on provided rules
	e.Pre(middleware.Rewrite(map[string]string{
		"/old":   "/new",
		"/api/*": "/$1",
		"/js":    "/public/javascript/$1",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo World!")
	})
	e.GET("/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "new")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
