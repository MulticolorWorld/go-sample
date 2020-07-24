package main

import (
	"github.com/labstack/echo/v4"
	"go-sample/controller"
	"go-sample/service"
)

func main() {
	e := echo.New()
	g := e.Group("")
	{
		s := service.NewMainService()
		c := controller.NewMainController(*s)
		c.Handle(g)
	}
	e.Logger.Fatal(e.Start(":1323"))
}
