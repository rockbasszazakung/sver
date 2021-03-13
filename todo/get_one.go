package todo

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTodosHandlerbyid(c echo.Context) error {
	t := Todo{}
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error,err")
	}
	defer db.Close()

	id, _ := strconv.Atoi(c.Param("id"))
	stmt, err := db.Prepare("SELECT id,title,status FROM todos WHERE id=$1")
	if err != nil {
		// log.Fatal("can'tprepare query one row statment", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	row := stmt.QueryRow(id)
	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		log.Fatal("can't Scan row into ", err)

	}
	return c.JSON(http.StatusOK, t)
}
