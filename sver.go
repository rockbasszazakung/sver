package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"mwssage": "hello",
	})
}

func main() {
	e := echo.New()
	e.GET("/hello", helloHandler)
	port := os.Getenv("PORT")
	//log.Print("Port", port)
	e.Start(":" + port)
}
