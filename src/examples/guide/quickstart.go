package main

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

func main() {
	// create echo instance
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo World!")
	})
	// Routing
	e.POST("/users", saveUser)
	// Path Parameters
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	// Query Parameters
	e.GET("/show", show)
	// Form application/x-www-form-urlencoded
	e.POST("/save", save)
	// Form multipart/form-data
	e.POST("/upload", upload)
	// Handling Request
	e.POST("/bind", bindUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func saveUser(c echo.Context) error {
	return c.String(http.StatusOK, "Request saveUser")
}
func getUser(c echo.Context) error {
	id := c.Param("id") // get path parameter
	msg := fmt.Sprintf("Request getUser by id:%s", id)
	return c.String(http.StatusOK, msg)
}
func updateUser(c echo.Context) error {
	id := c.Param("id")
	msg := fmt.Sprintf("Request updateUser by id:%s", id)
	return c.String(http.StatusOK, msg)
}
func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "Request deleteUser")
}

// e.g. /show?team=x-men&member=wolverine
func show(c echo.Context) error {
	team := c.QueryParam("team") // get query parameter
	member := c.QueryParam("member")
	msg := fmt.Sprintf("team:%s, member:%s", team, member)
	return c.String(http.StatusOK, msg)
}

func save(c echo.Context) error {
	// Get name and email from form
	name := c.FormValue("name")
	email := c.FormValue("email")
	msg := fmt.Sprintf("name:%s, email:%s", name, email)
	return c.String(http.StatusOK, msg)
}

func upload(c echo.Context) error {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar") // get form file
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	markup := fmt.Sprintf("<b>Thank you! %s</b>", name)
	return c.HTML(http.StatusOK, markup)
}

func bindUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	// Static Content
	// Serve any file from static directory for path /static/*.
	// e.Static("/static", "static")
	return c.JSON(http.StatusCreated, u)
}
