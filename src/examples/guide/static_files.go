package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()

	// Using Echo#Static()
	// register a new route with path prefix to serve static files
	// Usage1. /static/js/main.js will fetch and serve /assets/js/main.js file
	e.Static("/static", "assets")
	// Usage2. /js/main.js will fetch and serve /assets/js/main.js
	e.Static("/", "assets")

	// Using Echo#File()
	// register a new route with path to serve static file
	e.File("/", "/public/index.html")
	e.File("/favicon.ico", "/images/favicon.ico")

	e.Logger.Fatal(e.Start(":8080"))
}
