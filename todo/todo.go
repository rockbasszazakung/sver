package todo

import (
	"net/http"

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
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	return c.JSON(http.StatusOK, items)
}

// func GetTodosHandlerbyid(c echo.Context) error {
// 	t := Todo{}
// 	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		log.Fatal("Connect to database error,err")
// 	}
// 	defer db.Close()

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	stmt, err := db.Prepare("SELECT id,title,status FROM todos WHERE id=$1")
// 	if err != nil {
// 		log.Fatal("can'tprepare query one row statment", err)
// 	}

// 	row := stmt.QueryRow(id)
// 	err = row.Scan(&t.ID, &t.Title, &t.Status)
// 	if err != nil {
// 		log.Fatal("can't Scan row into ", err)

// 	}
// 	return c.JSON(http.StatusOK, t)
// }
