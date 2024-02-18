package main

//https://dzone.com/articles/building-a-concurrent-chat-app-with-go-and-websock Made by following this as shell

import (
	"ChatApp/templates"
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)
func main() {
	e := echo.New()
	component := templates.Hello("John")
	component.Render(context.Background(), os.Stdout)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Static("/static","static")
	e.Logger.Fatal(e.Start(":1323"))
}

