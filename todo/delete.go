package todo

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Deletehello(c echo.Context) error {
	// id, _ := strconv.Atoi(c.Param("id"))
	// delete(todos, id)
	// return c.JSON(http.StatusCreated, "DELETE todo")

	id, _ := strconv.Atoi(c.Param("id"))
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		// log.Fatal("Connect to database error,err")
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM todos WHERE  id=$1;")

	if err != nil {
		// log.Fatal("can't prepare statment update", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	if _, err := stmt.Exec(id); err != nil {
		// log.Fatal("error execute update", err)
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusCreated, "DELETE todo")
}
