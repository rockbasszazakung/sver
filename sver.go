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
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: {ID: 1, Title: "Pay Com", Status: "active"},
}

func createTodosHandler(e echo.Context) error {

	t := Todo{}
	if err := e.Bind(&t); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	id := len(todos)
	id++
	t.ID = id
	todos[t.ID] = &t
	return e.JSON(http.StatusCreated, "created todo")

}

func getTodosHandler(c echo.Context) error {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	return c.JSON(http.StatusOK, items)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/hello", helloHandler)
	e.GET("/todo", getTodosHandler)
	e.POST("/todo", createTodosHandler)

	port := os.Getenv("PORT")
	//log.Print("Port", port)
	e.Start(":" + port)
}
