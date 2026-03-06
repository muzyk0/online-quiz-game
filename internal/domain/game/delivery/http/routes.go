package http

import "github.com/labstack/echo/v4"

// RegisterRoutes registers public quiz game routes.
// All routes require JWT authentication via authMiddleware.
func RegisterRoutes(e *echo.Echo, h *Handler, authMiddleware echo.MiddlewareFunc) {
	g := e.Group("/api/pair-game-quiz/pairs", authMiddleware)
	g.POST("/connection", h.Connect)
	g.GET("/my-current", h.GetMyCurrent)
	g.GET("/my", h.GetMyGames) // must be before /:id
	g.POST("/my-current/answers", h.SubmitAnswer)
	g.GET("/:id", h.GetByID)

	u := e.Group("/api/pair-game-quiz/users", authMiddleware)
	u.GET("/my-statistic", h.GetMyStatistic)

	// Public routes (no auth)
	pub := e.Group("/api/pair-game-quiz/users")
	pub.GET("/top", h.GetTopPlayers)
}
