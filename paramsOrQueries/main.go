package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main endpoint has get request")
}

func userHandler(c echo.Context) error {
	dataType := c.Param("data")

	username := c.QueryParam("username")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Username : %s, Name: %s, Surname: %s", username, name, surname))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
			"surname":  surname,
		})
	}

	//fmt.Println(username)
	//fmt.Println(name)
	//fmt.Println(surname)

	//return c.String(http.StatusOK, "User endpoint has get request")
	return c.String(http.StatusBadRequest, "Only string or JSON format is accepted")
}

func main() {
	e := echo.New()

	e.GET("/main", mainHandler)
	e.GET("/user/:data", userHandler)

	e.Start(":8080")
}
