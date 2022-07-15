package main

import (
	"fmt"
	"net/http"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MyUUID struct {
	uuid string `json:"uuid"`
}

func main() {

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/", func(c echo.Context) error {
		id := uuid.New()
		return c.JSON(http.StatusOK, struct{ Host_Id string }{Host_Id: id.String()})
	})

	e.GET("/db-uuid", func(c echo.Context) error {
		db, err := sql.Open("mysql", "root:password@tcp(db:3306)/test")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		fmt.Println("Success!")

		results, err := db.Query("SELECT uuid FROM uuid LIMIT 1")
		if err != nil {
			panic(err.Error())
		}

		var myUuid MyUUID
		if results.Next() {
			err = results.Scan(&myUuid.uuid)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(myUuid.uuid)
		}

		return c.JSON(http.StatusOK, struct{ Host_Id string }{Host_Id: myUuid.uuid})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
