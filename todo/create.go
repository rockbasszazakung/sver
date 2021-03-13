package todo

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func CreateTodosHandler(e echo.Context) error {

	t := Todo{}
	if err := e.Bind(&t); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// id := len(todos)
	// id++
	// t.ID = id
	// todos[t.ID] = &t
	// return e.JSON(http.StatusCreated, "created todo")
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO todos (title, status) values ($1, $2)  RETURNING id", t.Title, t.Status)
	err = row.Scan(&t.ID)
	str := fmt.Sprintf("insert todo success id : ", t.ID)
	return e.JSON(http.StatusCreated, str)
}
