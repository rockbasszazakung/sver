package todo

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Puthello2(c echo.Context) error {
	t := new(Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	todos[id].Title = t.Title
	todos[id].Status = t.Status
	return c.JSON(http.StatusCreated, "PUT todo2")
}
