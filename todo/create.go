package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateTodosHandler(e echo.Context) error {

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
