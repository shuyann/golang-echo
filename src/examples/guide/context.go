package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	// Create a middleware to extend default context
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return h(cc)
		}
	})
	e.GET("/", func(c echo.Context) error {
		cc := c.(*CustomContext)
		cc.Foo()
		cc.Bar()
		return c.String(http.StatusOK, "custom context")
	})
	e.Logger.Fatal(e.Start(":8080"))
}

// Extending Context
// define custom context
type CustomContext struct {
	echo.Context
}

// implement custom context method
func (c *CustomContext) Foo() {
	fmt.Println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}
