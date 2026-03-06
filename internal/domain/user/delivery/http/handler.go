package http

import (
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/muzyk0/online-quiz-game/internal/domain/user/delivery/http/dto"
	userservice "github.com/muzyk0/online-quiz-game/internal/domain/user/service"
	"github.com/muzyk0/online-quiz-game/internal/pkg/apperrors"
	"github.com/muzyk0/online-quiz-game/internal/pkg/auth"
	"github.com/muzyk0/online-quiz-game/internal/pkg/helpers"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service      userservice.UserServiceInterface
	tokenManager *auth.TokenManager
}

func NewHandler(service userservice.UserServiceInterface, tokenManager *auth.TokenManager) *Handler {
	return &Handler{
		service:      service,
		tokenManager: tokenManager,
	}
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Create a new user account and return access and refresh tokens
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RegisterRequest	true	"Registration request"
//	@Success		201		{object}	dto.AuthResponse	"User registered successfully"
//	@Failure		400		{object}	map[string]string	"Validation error"
//	@Failure		409		{object}	map[string]string	"Email already in use"
//	@Router			/api/auth/register [post]
func (h *Handler) Register(c echo.Context) error {
	var req dto.RegisterRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := h.service.Register(ctx, req.ToDomain())
	if err != nil {
		return mapUserServiceError(err)
	}

	accessToken, err := h.tokenManager.GenerateAccessToken(user.ID, user.Email, "user")
	if err != nil {
		return apperrors.Internal("Could not generate access token").Wrap(err)
	}

	tokenID := uuid.New().String()
	refreshToken, err := h.tokenManager.GenerateRefreshToken(user.ID, user.Email, "user", tokenID)
	if err != nil {
		return apperrors.Internal("Could not generate refresh token").Wrap(err)
	}

	c.SetCookie(auth.NewRefreshTokenCookie(refreshToken))

	return c.JSON(nethttp.StatusCreated, dto.AuthResponse{
		User:         dto.UserResponseFromDomain(user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Authenticate with email and password, return access and refresh tokens
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequest	true	"Login request"
//	@Success		200		{object}	dto.AuthResponse	"Login successful"
//	@Failure		400		{object}	map[string]string	"Validation error"
//	@Failure		401		{object}	map[string]string	"Invalid credentials"
//	@Router			/api/auth/login [post]
func (h *Handler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := h.service.LoginByLoginOrEmail(ctx, req.LoginOrEmail, req.Password)
	if err != nil {
		return apperrors.Unauthorized("Invalid credentials")
	}

	accessToken, err := h.tokenManager.GenerateAccessToken(user.ID, user.Email, "user")
	if err != nil {
		return apperrors.Internal("Could not generate access token").Wrap(err)
	}

	tokenID := uuid.New().String()
	refreshToken, err := h.tokenManager.GenerateRefreshToken(user.ID, user.Email, "user", tokenID)
	if err != nil {
		return apperrors.Internal("Could not generate refresh token").Wrap(err)
	}

	c.SetCookie(auth.NewRefreshTokenCookie(refreshToken))

	return c.JSON(nethttp.StatusOK, dto.LoginSuccessResponse{
		AccessToken: accessToken,
	})
}

// GetProfile godoc
//
//	@Summary		Get current user profile
//	@Description	Return the authenticated user's profile information
//	@Tags			Users
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	dto.UserResponse	"User profile"
//	@Failure		401	{object}	map[string]string	"Unauthorized"
//	@Failure		404	{object}	map[string]string	"User not found"
//	@Router			/api/protected/profile [get]
func (h *Handler) GetProfile(c echo.Context) error {
	userID := auth.MustGetUserID(c)
	user, err := h.service.GetUser(c.Request().Context(), userID)
	if err != nil {
		return mapUserServiceError(err)
	}
	return c.JSON(nethttp.StatusOK, dto.UserResponseFromDomain(user))
}

// UpdateProfile godoc
//
//	@Summary		Update current user profile
//	@Description	Update the authenticated user's first name, last name, or avatar URL
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		dto.UpdateProfileRequest	true	"Profile update request"
//	@Success		200		{object}	dto.UserResponse			"Updated user profile"
//	@Failure		400		{object}	map[string]string			"Validation error"
//	@Failure		401		{object}	map[string]string			"Unauthorized"
//	@Failure		404		{object}	map[string]string			"User not found"
//	@Router			/api/protected/profile [put]
func (h *Handler) UpdateProfile(c echo.Context) error {
	userID := auth.MustGetUserID(c)

	var req dto.UpdateProfileRequest
	if err := helpers.BindAndValidate(c, &req); err != nil {
		return err
	}

	user, err := h.service.UpdateProfile(c.Request().Context(), userID, req.ToDomain())
	if err != nil {
		return mapUserServiceError(err)
	}
	return c.JSON(nethttp.StatusOK, dto.UserResponseFromDomain(user))
}

// DeleteAccount godoc
//
//	@Summary		Delete current user account
//	@Description	Permanently delete the authenticated user's account
//	@Tags			Users
//	@Security		BearerAuth
//	@Success		204	"Account deleted"
//	@Failure		401	{object}	map[string]string	"Unauthorized"
//	@Failure		500	{object}	map[string]string	"Internal server error"
//	@Router			/api/protected/account [delete]
func (h *Handler) DeleteAccount(c echo.Context) error {
	userID := auth.MustGetUserID(c)
	if err := h.service.DeleteUser(c.Request().Context(), userID); err != nil {
		return apperrors.Internal("Failed to delete account").Wrap(err)
	}
	return c.NoContent(nethttp.StatusNoContent)
}
