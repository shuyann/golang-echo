package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/string", sendString)
	e.GET("/html", sendHTML)
	e.GET("/json", sendJSON)
	e.GET("/streamjson", sendStreamJSON)
	e.GET("/prettyjson", sendPrettyJSON)
	e.GET("/blobjson", sendBlobJSON)
	e.GET("/file", sendFile)
	e.GET("/nocontent", noContent)
	e.GET("/redirect", redirect)
	e.GET("/hooks", hooks)
	e.Logger.Fatal(e.Start(":8080"))
}

func sendString(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func sendHTML(c echo.Context) error {
	markup := "<strong>Hello, World!</strong>"
	return c.HTML(http.StatusOK, markup)
}

func sendJSON(c echo.Context) error {
	u := &User{
		Name:  "foo",
		Email: "foo@com",
	}
	return c.JSON(http.StatusOK, u)
}

func sendStreamJSON(c echo.Context) error {
	u := &User{
		Name:  "foo",
		Email: "foo@com",
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(u)
}

func sendPrettyJSON(c echo.Context) error {
	u := &User{
		Name:  "foo",
		Email: "foo@com",
	}
	return c.JSONPretty(http.StatusOK, u, " ")
}

func sendBlobJSON(c echo.Context) error {
	encodedJSON := []byte{} // Encoded JSON from external source
	return c.JSONBlob(http.StatusOK, encodedJSON)
}

func sendFile(c echo.Context) error {
	return c.File("<PATH_TO_FILE>")
}

func noContent(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func redirect(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "<URL>")
}

func hooks(c echo.Context) error {
	c.Response().Before(func() {
		log.Println("before response")
	})
	c.Response().After(func() {
		log.Println("after response")
	})
	return c.NoContent(http.StatusNoContent)
}
