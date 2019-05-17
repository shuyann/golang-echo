package main

import (
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// redirect middleware redirects http requests to https
	// http://localhost -> https://localhost
	e.Pre(middleware.HTTPSRedirect())
	// http://localhost -> https://www.localhost
	e.Pre(middleware.HTTPSWWWRedirect())
	// http://www.localhost -> https://localhost
	e.Pre(middleware.HTTPSNonWWWRedirect())
	// http://localhost -> http://www.localhost
	e.Pre(middleware.WWWRedirect())
	// http://www.localhost -> http://localhost
	e.Pre(middleware.NonWWWRedirect())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
