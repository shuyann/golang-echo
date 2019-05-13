package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/users", handler)
	e.Logger.Fatal(e.Start(":1323"))
}

// Bind Data
func handler(c echo.Context) (err error) {
	u := new(User)
	// Bind supports decoding application/json,application/xml,application/x-www-form-urlencoded
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}

// Custom Binder
type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// you may use default binder
	binder := new(echo.DefaultBinder)
	if err = binder.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}

	// Define your custom implementation
	return
}

func retrive(c echo.Context) (err error) {
	// Form Data
	name1 := c.FormValue("name")
	// Query Parameters
	name2 := c.QueryParam("name")
	// Path Parameters
	name3 := c.Param("name")
	fmt.Println(name1, name2, name3)
	return
}
