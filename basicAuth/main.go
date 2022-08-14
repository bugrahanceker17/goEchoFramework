package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/main", mainHandler)

	adminGroup := e.Group("/admin")

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "123" {
			return true, nil
		}

		return false, nil
	}))

	adminGroup.GET("/main", mainAdmin)

	e.Start(":8080")
}

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main endpoint has get request")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Welcome Master!"))
}
