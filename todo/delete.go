package todo

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Deletehello(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(todos, id)
	return c.JSON(http.StatusCreated, "DELETE todo")
}
