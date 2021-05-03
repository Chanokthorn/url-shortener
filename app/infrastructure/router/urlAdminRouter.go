package router

import (
	"github.com/labstack/echo/v4"
	"url-shortener/app/infrastructure/middleware"
	"url-shortener/app/interface/controller"
)

func CreateURLAdminRouter(g *echo.Group, c controller.URLAdminController, m *middleware.AdminTokenMiddleware) {
	g.Use(m.CheckTokenMiddleware)
	g.GET("/url", c.ListURL)
	g.DELETE("/url/:short-code", c.DeleteURL)
}
