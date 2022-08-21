package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main endpoint has get request")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("Welcome Master!"))
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "statusCode=${status}, method=${method}, uri=${uri}\n",
	}))

	e.GET("/main", mainHandler)

	adminGroup := e.Group("/admin", middleware.Logger())

	adminGroup.GET("/main", mainAdmin)

	e.Start(":8080")
}

// help endpoint'i "/admin/main" şeklinde kullanılabilir. Çünkü group'da tanımlanmıştır.
// help endpoint'i aslında alt endpoint'tir.
// g.Use() sadece group endpoint'leri için middleware tanımlar
// e.User() ise bütün endpointler için tanımlar
// Skipping middleware; belirli koşullarda middleware i çalıştırmamak. Örn : localhost'ta middleware'i çalıştırma
