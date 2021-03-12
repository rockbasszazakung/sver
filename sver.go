package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"mwssage": "hello",
		"NAME":    "BAs",
	})
}

//PUT /DELETE
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

func puthello(c echo.Context) error {
	t := Todo{}
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	todos[t.ID] = &t
	return c.JSON(http.StatusCreated, "PUT todo")
}
func puthello2(c echo.Context) error {
	// t := new(todos)
	// if err := c.Bind(t); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	// }
	// id, _ := strconv.Atoi(c.Param("id"))
	// todos[id].Title = t.Title
	// todos[id].Status = t.Status
	return c.JSON(http.StatusCreated, "PUT todo2")
}

func deletehello(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(todos, id)
	return c.JSON(http.StatusCreated, "DELETE todo")
}

func getTodosHandler(c echo.Context) error {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	return c.JSON(http.StatusOK, items)
}

func getTodosHandlerbyid(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	t, ok := todos[id]
	if !ok {
		return c.JSON(http.StatusOK, map[int]string{})
	}
	return c.JSON(http.StatusOK, t)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/hello", helloHandler)
	e.GET("/todo", getTodosHandler)
	e.GET("/todo/:id", getTodosHandlerbyid)
	e.POST("/todo", createTodosHandler)
	e.PUT("/todo", puthello)
	e.PUT("/todo2/:id", puthello2)
	e.DELETE("/todo/:id", deletehello)

	port := os.Getenv("PORT")
	//log.Print("Port", port)
	e.Start(":" + port)
}
