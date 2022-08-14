package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func addUser(c echo.Context) error {
	user := User{}

	req, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	json.Unmarshal(req, &user)

	return c.String(http.StatusCreated, fmt.Sprintf("Eklendi => %s", user.Username))
}

func main() {
	e := echo.New()

	e.POST("/user", addUser)

	e.Start(":8080")
}
