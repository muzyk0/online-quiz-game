package http

import (
	"github.com/labstack/echo/v4"

	"github.com/muzyk0/online-quiz-game/internal/app/middleware"
)

// RegisterRoutes registers auth domain HTTP routes on the /api/auth group.
func RegisterRoutes(e *echo.Echo, h *Handler, authMiddleware echo.MiddlewareFunc) {
	authGroup := e.Group("/api/auth")

	// Refresh endpoint
	refreshLimiter := middleware.NewRefreshRateLimiter()
	authGroup.POST("/refresh", h.Refresh,
		middleware.AuthRateLimitMiddleware(refreshLimiter, middleware.IPIdentifier))

	// Protected auth endpoints
	authGroup.GET("/me", h.GetMe, authMiddleware)
	authGroup.POST("/logout", h.Logout, authMiddleware)
	authGroup.POST("/change-email", h.ChangeEmail, authMiddleware)
	authGroup.POST("/change-password", h.ChangePassword, authMiddleware)
}
