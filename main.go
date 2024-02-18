package main

//https://dzone.com/articles/building-a-concurrent-chat-app-with-go-and-websock Made by following this as shell

import (
	
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/Graybb/MessageSystem/templates"
)
func main() {
	e := echo.New()
	component := templates.Hello("John")
	component.Render(context.Background(), os.Stdout)
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(),c.Response().Writer)
	})
	e.Static("/static","static")
	e.Logger.Fatal(e.Start(":1323"))
}

