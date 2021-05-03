package router

import (
	"github.com/labstack/echo/v4"
	"url-shortener/app/interface/controller"
)

func CreateURLClientRouter(g *echo.Group, c controller.URLClientController) {
	g.POST("/", c.CreateURL)
	g.GET("/short-code", c.GetShortCodeFromFullURL)
}
