package router

import (
	"github.com/labstack/echo/v4"
	"url-shortener/app/interface/controller"
)

func CreateURLRedirectRouter(g *echo.Group, c controller.URLRedirectController) {
	g.GET("/:short-code", c.Redirect)
}
