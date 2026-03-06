package http

import (
	"errors"

	userservice "github.com/muzyk0/online-quiz-game/internal/domain/user/service"
	"github.com/muzyk0/online-quiz-game/internal/platform/http/apperrors"
)

// mapAuthServiceError converts auth-related service errors to AppErrors
func mapAuthServiceError(err error) error {
	switch {
	case errors.Is(err, userservice.ErrUserNotFound):
		return apperrors.Unauthorized("User not found")
	case errors.Is(err, userservice.ErrInvalidPassword):
		return apperrors.Unauthorized("Current password is incorrect")
	case errors.Is(err, userservice.ErrUserAlreadyExists):
		return apperrors.Conflict("Email already in use")
	default:
		return apperrors.Internal("Failed to process request").Wrap(err)
	}
}
