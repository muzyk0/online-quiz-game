package dto

import (
	userservice "github.com/muzyk0/online-quiz-game/internal/domain/user/service"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"Error message" validate:"required"`
}

// UserResponse is the handler-level DTO for user data
type UserResponse struct {
	ID        string `json:"id" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarUrl string `json:"avatar_url"`
}

// UserResponseFromDomain maps service layer UserOutput to handler layer UserResponse
func UserResponseFromDomain(user *userservice.UserOutput) *UserResponse {
	if user == nil {
		return nil
	}
	return &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		AvatarUrl: user.AvatarUrl,
	}
}

// AuthResponse contains user info with access and refresh tokens
type AuthResponse struct {
	// User information
	User *UserResponse `json:"user" validate:"required"`
	// Access token (short-lived, 15 minutes)
	AccessToken string `json:"accessToken" validate:"required"` //nolint:gosec // API field name for auth response
	// Refresh token (long-lived, 7 days) - also set as httpOnly cookie
	RefreshToken string `json:"refreshToken" validate:"required"` //nolint:gosec // API field name for auth response
}

// LoginSuccessResponse is the spec-compliant login response body.
type LoginSuccessResponse struct {
	AccessToken string `json:"accessToken"` //nolint:gosec // API field name for login response
}

