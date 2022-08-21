package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

func main() {

	e := echo.New()

	adminGroup := e.Group("/admin")
	adminGroup.GET("/login", loginAdmin, checkCookie)

	e.Start(":8080")
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("userData")

		if err != nil {
			fmt.Println(err)
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "Unauthorized!")
			}

			return err
		}

		if cookie.Value == "admin" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Unauthorized!")
	}
}

func loginAdmin(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "admin" && password == "590" {
		cookie := http.Cookie{
			Name:    "userData",
			Value:   username,
			Expires: time.Now().Add(48 * time.Hour),
		}

		c.SetCookie(&cookie)
		return c.String(http.StatusOK, "Logged in! Welcome!")
	}

	return c.String(http.StatusOK, "Login failed!")
}
