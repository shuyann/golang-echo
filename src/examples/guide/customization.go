package main

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
	"time"
)

func main() {
	e := echo.New()
	// Custom Logging
	e.Debug = true // use debug mode
	e.Logger.SetHeader("${time_rfc3339} ${level}")
	e.Logger.SetOutput(os.Stdout)
	e.GET("/", func(c echo.Context) error {
		e.Logger.Debug("foo")
		return c.String(http.StatusOK, "customization example")
	})

	// Custom Server
	s := &http.Server{
		Addr: ":8080",
		// Read timeout
		ReadTimeout: 20 * time.Minute,
		// Write timeout
		WriteTimeout: 20 * time.Minute,
	}
	// Startup Banner
	e.HideBanner = true

	e.Logger.Fatal(e.StartServer(s))
}
