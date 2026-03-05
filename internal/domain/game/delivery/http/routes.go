package http

import "github.com/labstack/echo/v4"

// RegisterRoutes registers public quiz game routes.
// All routes require JWT authentication via authMiddleware.
func RegisterRoutes(e *echo.Echo, h *Handler, authMiddleware echo.MiddlewareFunc) {
	g := e.Group("/api/pair-game-quiz/pairs", authMiddleware)
	g.POST("/connection", h.Connect)
	g.GET("/my-current", h.GetMyCurrent)
	g.POST("/my-current/answers", h.SubmitAnswer)
	g.GET("/:id", h.GetByID)
}
