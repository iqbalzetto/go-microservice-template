package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	//SET UP CONFIG

	//SET UP DATABASE CONNECTION

	//SET UP SERVER
	e := echo.New()
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Users!")

	})
	e.Logger.Fatal(e.Start(":4001"))
}
