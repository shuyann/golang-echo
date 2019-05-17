package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo World!")
	})
	// basic auth settings
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "foo" && password == "foo" {
			return true, nil
		}
		return false, nil
	}))
	e.Logger.Fatal(e.Start(":8080"))
}
