package auth

import (
	"crypto/subtle"

	"github.com/muzyk0/online-quiz-game/internal/pkg/apperrors"

	"github.com/labstack/echo/v4"
)

// BasicAuthMiddleware creates a middleware that validates HTTP Basic auth
// against a single admin login/password pair.
//
// Usage:
//
//	saMiddleware := auth.BasicAuthMiddleware(cfg.SAAdminLogin, cfg.SAAdminPassword)
//	e.Group("/api/sa", saMiddleware)
func BasicAuthMiddleware(login, password string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			gotLogin, gotPassword, ok := c.Request().BasicAuth()
			if !ok {
				c.Response().Header().Set("WWW-Authenticate", `Basic realm="SA API"`)
				return apperrors.Unauthorized("Basic authentication required")
			}

			loginMatch    := subtle.ConstantTimeCompare([]byte(gotLogin), []byte(login))
			passwordMatch := subtle.ConstantTimeCompare([]byte(gotPassword), []byte(password))

			if loginMatch != 1 || passwordMatch != 1 {
				return apperrors.Unauthorized("Invalid credentials")
			}

			return next(c)
		}
	}
}
