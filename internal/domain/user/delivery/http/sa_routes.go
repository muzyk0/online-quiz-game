package http

import "github.com/labstack/echo/v4"

// RegisterSARoutes registers super-admin user management routes.
// All routes require Basic auth via saMiddleware.
func RegisterSARoutes(e *echo.Echo, h *SAHandler, saMiddleware echo.MiddlewareFunc) {
	sa := e.Group("/api/sa/users", saMiddleware)
	sa.GET("", h.GetUsers)
	sa.POST("", h.CreateUser)
	sa.DELETE("/:id", h.DeleteUser)
}
