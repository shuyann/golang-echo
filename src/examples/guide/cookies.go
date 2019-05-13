package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/write", writeCookie)
	e.GET("/read/:name", readCookie)
	e.GET("/readall", readAllCookies)
	e.Logger.Fatal(e.Start(":8080"))
}

// Create a Cookie
func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "foo"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// adds a Set-Cookie header in HTTP response.
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

// Read a Cookie
func readCookie(c echo.Context) error {
	// get path cookieName
	cookieName := c.Param("name")
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return err
	}
	fmt.Printf("cookie (name:%s, value:%s)", cookie.Name, cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}

// Read All Cookies
func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Printf("cookie (name:%s, value:%s)", cookie.Name, cookie.Value)
	}
	return c.String(http.StatusOK, "read all cookies")
}
