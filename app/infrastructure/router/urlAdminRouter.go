package router

import (
	"github.com/labstack/echo/v4"
	"url-shortener/app/interface/controller"
)

func CreateURLAdminRouter(g *echo.Group, c controller.URLAdminController) {
	g.GET("/url", c.ListURL)
	g.DELETE("/url/:short-code", c.DeleteURL)
}
