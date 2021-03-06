package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"mwssage": "hello",
		"NAME":    "BAs",
	})
}

// `
type Todo struct {
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/todo", helloHandler)
	port := os.Getenv("PORT")
	//log.Print("Port", port)
	e.Start(":" + port)
}
