package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/muzyk0/online-quiz-game/internal/domain/auth/delivery/http/dto"
	userservice "github.com/muzyk0/online-quiz-game/internal/domain/user/service"
	"github.com/muzyk0/online-quiz-game/internal/pkg/apperrors"
	"github.com/muzyk0/online-quiz-game/internal/pkg/auth"
	"github.com/muzyk0/online-quiz-game/internal/pkg/helpers"
)

// UserServiceInterface defines what the auth handler needs from the user domain.
type UserServiceInterface interface {
	GetUser(ctx context.Context, userID string) (*userservice.UserOutput, error)
	ChangeEmail(ctx context.Context, userID, currentPassword, newEmail string) error
	ChangePassword(ctx context.Context, userID, currentPassword, newPassword string) error
}

// Handler handles cross-domain authentication endpoints
type Handler struct {
	userService  UserServiceInterface
	tokenManager *auth.TokenManager
}

// NewHandler creates a new auth Handler instance
func NewHandler(
	userService UserServiceInterface,
	tokenManager *auth.TokenManager,
) *Handler {
	return &Handler{
		userService:  userService,
		tokenManager: tokenManager,
	}
}

// GetMe godoc
//
//	@Summary		Get current user info
//	@Description	Returns email, login and userId of the authenticated user
//	@Tags			Authentication
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	dto.MeResponse
//	@Failure		401	{object}	map[string]string	"Unauthorized"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/auth/me [get]
func (h *Handler) GetMe(c echo.Context) error {
	userID := auth.MustGetUserID(c)
	user, err := h.userService.GetUser(c.Request().Context(), userID)
	if err != nil {
		return mapAuthServiceError(err)
	}
	return c.JSON(http.StatusOK, dto.MeResponse{
		Email:  user.Email,
		Login:  user.Login,
		UserID: user.ID,
	})
}

// Refresh godoc
//
//	@Summary		Refresh access token
//	@Description	Exchange refresh token for a new access token. Accepts refresh token via httpOnly cookie (web clients) or Authorization Bearer header (mobile clients). Implements token rotation - returns new refresh token on success.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.RefreshResponse		"Token refreshed successfully"
//	@Failure		401	{object}	map[string]string	"Invalid or expired refresh token"
//	@Router			/auth/refresh [post]
func (h *Handler) Refresh(c echo.Context) error {
	var refreshToken string

	// Try to get refresh token from cookie first (web clients)
	cookie, err := c.Cookie("refreshToken")
	if err == nil && cookie.Value != "" {
		refreshToken = cookie.Value
	} else {
		// Try Authorization header (mobile clients)
		authHeader := c.Request().Header.Get("Authorization")
		if token, found := strings.CutPrefix(authHeader, "Bearer "); found {
			refreshToken = token
		} else {
			// Try request body (alternative for mobile)
			var req dto.RefreshRequest
			if err := c.Bind(&req); err == nil && req.RefreshToken != "" {
				refreshToken = req.RefreshToken
			}
		}
	}

	if refreshToken == "" {
		return apperrors.Unauthorized("No refresh token provided")
	}

	// Validate refresh token
	claims, err := h.tokenManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return apperrors.Unauthorized("Invalid or expired refresh token")
	}

	// Generate new access token
	newAccessToken, err := h.tokenManager.GenerateAccessToken(claims.UserID, claims.Email, claims.UserType)
	if err != nil {
		return apperrors.Internal("Failed to generate access token").Wrap(err)
	}

	// Generate new refresh token (rotation)
	newTokenID := uuid.New().String()
	newRefreshToken, err := h.tokenManager.GenerateRefreshToken(claims.UserID, claims.Email, claims.UserType, newTokenID)
	if err != nil {
		return apperrors.Internal("Failed to generate refresh token").Wrap(err)
	}

	// Set new refresh token cookie for web clients
	c.SetCookie(auth.NewRefreshTokenCookie(newRefreshToken))

	// Return both tokens in response for mobile clients
	return c.JSON(http.StatusOK, dto.RefreshResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
}

// Logout godoc
//
//	@Summary		Logout user
//	@Description	Clear refresh token cookie and invalidate session
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	dto.MessageResponse		"Logout successful"
//	@Router			/auth/logout [post]
func (h *Handler) Logout(c echo.Context) error {
	// Clear refresh token cookie
	c.SetCookie(auth.ClearRefreshTokenCookie())

	return c.JSON(http.StatusOK, dto.MessageResponse{
		Message: "Logged out successfully",
	})
}

// ChangeEmail godoc
//
//	@Summary		Change user email
//	@Description	Change the authenticated user's email address with password verification. Requires current password to prevent unauthorized changes.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		dto.ChangeEmailRequest	true	"Email change request"
//	@Success		200		{object}	dto.MessageResponse		"Email changed successfully"
//	@Failure		400		{object}	map[string]string	"Invalid request body or validation error"
//	@Failure		401		{object}	map[string]string	"Unauthorized or incorrect password"
//	@Failure		409		{object}	map[string]string	"Email already in use"
//	@Failure		500		{object}	map[string]string	"Internal server error"
//	@Router			/auth/change-email [post]
func (h *Handler) ChangeEmail(c echo.Context) error {
	userID := auth.MustGetUserID(c)

	var req dto.ChangeEmailRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	err := h.userService.ChangeEmail(ctx, userID, req.CurrentPassword, req.NewEmail)
	if err != nil {
		return mapAuthServiceError(err)
	}

	return c.JSON(http.StatusOK, dto.MessageResponse{
		Message: "Email changed successfully",
	})
}

// ChangePassword godoc
//
//	@Summary		Change user password
//	@Description	Change the authenticated user's password with current password verification. This will invalidate all existing sessions except the current one.
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		dto.ChangePasswordRequest	true	"Password change request"
//	@Success		200		{object}	dto.MessageResponse			"Password changed successfully"
//	@Failure		400		{object}	map[string]string		"Invalid request body or validation error"
//	@Failure		401		{object}	map[string]string		"Unauthorized or incorrect password"
//	@Failure		500		{object}	map[string]string		"Internal server error"
//	@Router			/auth/change-password [post]
func (h *Handler) ChangePassword(c echo.Context) error {
	userID := auth.MustGetUserID(c)

	var req dto.ChangePasswordRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	err := h.userService.ChangePassword(ctx, userID, req.CurrentPassword, req.NewPassword)
	if err != nil {
		return mapAuthServiceError(err)
	}

	return c.JSON(http.StatusOK, dto.MessageResponse{
		Message: "Password changed successfully",
	})
}
