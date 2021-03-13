package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Puthello(c echo.Context) error {
	t := Todo{}
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	todos[t.ID] = &t
	return c.JSON(http.StatusCreated, "PUT todo")
}
