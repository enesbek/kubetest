package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name 	string 
	Surname string
}

func GetUser(c echo.Context) error {
	a := User{
		Name: "Ahmet", 
		Surname: "Mehmet",
	}

	return c.JSON(http.StatusOK, a)
}
