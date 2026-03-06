package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSMiddleware sets up CORS headers with credentials support for httpOnly cookies
// Required for cross-domain authentication where refresh tokens are stored in httpOnly cookies
func CORSMiddleware(allowedOrigins []string) echo.MiddlewareFunc {
	// Validate CORS configuration: wildcard origins are incompatible with credentials
	hasWildcard := false
	for _, origin := range allowedOrigins {
		if origin == "*" {
			hasWildcard = true
			break
		}
	}

	if hasWildcard {
		log.Println("WARNING: CORS wildcard '*' origin is incompatible with AllowCredentials=true. " +
			"Credentials will be ignored. Please specify explicit origins in CORS_ALLOWED_ORIGINS.")
	}

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		ExposeHeaders:    []string{echo.HeaderAuthorization},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	})
}
