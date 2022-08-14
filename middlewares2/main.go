package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()
	e.Use(setHeader)

	adminGroup := e.Group("/admin")

	adminGroup.GET("/main", mainAdmin)

	e.Start(":8080")
}

func setHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("Content-Type")
		fmt.Println(contentType)

		if contentType != "application/json" {
			return c.String(http.StatusBadRequest, "Content type must be application/json")
		}

		return next(c)
	}
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Welcome Master!"))
}
