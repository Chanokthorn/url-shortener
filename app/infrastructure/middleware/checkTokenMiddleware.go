package middleware

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AdminTokenMiddleware struct {
	adminToken string
}

func NewAdminTokenMiddleware(adminToken string) *AdminTokenMiddleware {
	return &AdminTokenMiddleware{adminToken: adminToken}
}

func (m *AdminTokenMiddleware) extractTokenFromHeader(e echo.Context) (string, error) {
	token := e.Request().Header.Get("admin-token")
	if token == "" {
		return "", errors.New("empty admin-token header")
	}
	return token, nil
}

func (m *AdminTokenMiddleware) CheckTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		token, err := m.extractTokenFromHeader(e)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("authorization error: %v", err))
		}
		if token != m.adminToken {
			return echo.NewHTTPError(http.StatusForbidden, "invalid admin token")
		}
		return next(e)
	}
}
