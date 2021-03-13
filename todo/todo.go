package todo

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: {ID: 1, Title: "Pay Com", Status: "active"},
}

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"mwssage": "hello",
		"NAME":    "BAs",
	})
}

func GetTodosHandler(c echo.Context) error {
	// t := Todo{}
	items := []*Todo{}
	// for _, item := range todos {
	// 	items = append(items, item)
	// }
	// return c.JSON(http.StatusOK, items)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error,err")
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}
	for rows.Next() {
		item := &Todo{}
		err := rows.Scan(&item.ID, &item.Title, &item.Status)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}
		// return c.JSON(http.StatusOK, t)
		items = append(items, item)

	}
	return c.JSON(http.StatusOK, items)
	// fmt.Println("query all todos success")
}
