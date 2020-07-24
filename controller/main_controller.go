package controller

import (
	"github.com/labstack/echo/v4"
	"go-sample/service"
	"net/http"
)

type MainController struct {
	service service.MainService
}

func NewMainController(ms service.MainService) *MainController {
	return &MainController{service: ms}
}

func (c MainController) Handle(g *echo.Group) {
	g.GET("/", index(c))
}

func index(c MainController) echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.String(http.StatusOK, c.service.Hello())
	}
}
