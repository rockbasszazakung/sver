package todo

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Puthello2(c echo.Context) error {
	t := new(Todo)
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	id, _ := strconv.Atoi(c.Param("id"))
	// todos[id].Title = t.Title
	// todos[id].Status = t.Status
	// return c.JSON(http.StatusCreated, "PUT todo2")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		// log.Fatal("Connect to database error,err")
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE todos SET title=$2 , status=$3 WHERE id=$1;")

	if err != nil {
		// log.Fatal("can't prepare statment update", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	if _, err := stmt.Exec(id, t.Title, t.Status); err != nil {
		// log.Fatal("error execute update", err)
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusCreated, "PUT todo2")

}
