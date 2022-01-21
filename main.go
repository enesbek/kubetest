package main

import (
	"httpServer/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handler.GetUser)

	e.Logger.Fatal(e.Start(":3001"))
}
