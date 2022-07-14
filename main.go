package main

import (
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/host_id", func(c echo.Context) error {
		id := uuid.New()
		return c.JSON(http.StatusOK, struct{ Host_Id string }{Host_Id: id.String()})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
