package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main endpoint'i get isteği yapıldı.")
}

func main() {
	fmt.Printf("Hello world")

	e := echo.New()

	e.GET("/main", mainHandler)

	e.Start(":8080")
}
