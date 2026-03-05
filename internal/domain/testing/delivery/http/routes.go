package http

import (
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers testing utility routes on the Echo instance.
func RegisterRoutes(e *echo.Echo, h *TestingHandler) {
	e.DELETE("/api/testing/all-data", h.DeleteAllData)
}
