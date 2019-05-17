package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	// download a file
	e.GET("/file", func(c echo.Context) error {
		return c.File("echo.svg")
	})
	// download a file as inline
	e.GET("/inline", func(c echo.Context) error {
		return c.Inline("inline.txt", "inline.txt")
	})
	// download a file as attachment
	e.GET("/attachment", func(c echo.Context) error {
		return c.Attachment("attachment.txt", "attachment.txt")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
