package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/hello", hello)
	e.Any("/any", func(c echo.Context) error {
		return c.String(http.StatusOK, "handle for all HTTP methods")
	})
	e.Match([]string{http.MethodGet, http.MethodPost}, "/match", func(c echo.Context) error {
		return c.String(http.StatusOK, "handle for GET/POST HTTP methods")
	})

	// Path Matching Order
	e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})

	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	})

	e.GET("/users/1/files/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/1/files/*")
	})
	// Above routes would resolve /users/new -> users/:id -> users/1/files/*

	// Group
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "foo" && password == "foo" {
			return true, nil
		}
		return false, nil
	}))

	// List Routes
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	// write route list in json file
	ioutil.WriteFile("routes.json", data, 0644)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
