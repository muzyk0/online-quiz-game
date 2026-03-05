package http

import "github.com/labstack/echo/v4"

// RegisterRoutes registers SA quiz question routes.
// All routes require Basic auth via saMiddleware.
func RegisterRoutes(e *echo.Echo, h *Handler, saMiddleware echo.MiddlewareFunc) {
	sa := e.Group("/api/sa/quiz/questions", saMiddleware)
	sa.GET("", h.GetQuestions)
	sa.POST("", h.CreateQuestion)
	sa.PUT("/:id", h.UpdateQuestion)
	sa.DELETE("/:id", h.DeleteQuestion)
	sa.PUT("/:id/publish", h.PublishQuestion)
}
