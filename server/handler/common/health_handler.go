package common

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/server"
)

var _ server.HTTPHandler = (*HealthHandler)(nil)

type HealthHandler struct{}

func (h *HealthHandler) Method() string {
	return http.MethodGet
}

func (h *HealthHandler) Path() string {
	return "/health"
}

func (h *HealthHandler) HandleFunc() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}
}
