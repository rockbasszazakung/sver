package main

import (
	"net/http"
	"os"

	"github.com/bas/sver/todo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"mwssage": "hello",
		"NAME":    "BAs",
	})
}

//PUT /DELETE
// `

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/hello", helloHandler)
	e.GET("/todo", todo.GetTodosHandler)
	e.GET("/todo/:id", todo.GetTodosHandlerbyid)
	e.POST("/todo", todo.CreateTodosHandler)
	e.PUT("/todo", todo.Puthello)
	e.PUT("/todo2/:id", todo.Puthello2)
	e.DELETE("/todo/:id", todo.Deletehello)

	port := os.Getenv("PORT")
	//log.Print("Port", port)
	e.Start(":" + port)
}
