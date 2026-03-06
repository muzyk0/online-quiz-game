package http

import (
	"errors"

	"github.com/muzyk0/online-quiz-game/internal/domain/game/service"
	"github.com/muzyk0/online-quiz-game/internal/platform/http/apperrors"
)

func mapGameError(err error) error {
	switch {
	case errors.Is(err, service.ErrInvalidGameID):
		return apperrors.BadRequest("Invalid game id format")
	case errors.Is(err, service.ErrGameNotFound):
		return apperrors.NotFound("Game not found")
	case errors.Is(err, service.ErrAlreadyInGame):
		return apperrors.Forbidden("You are already participating in an active game")
	case errors.Is(err, service.ErrGameNotActive):
		return apperrors.Forbidden("No active game found")
	case errors.Is(err, service.ErrAllQuestionsAnswered):
		return apperrors.Forbidden("You have already answered all questions in this game")
	case errors.Is(err, service.ErrNotEnoughQuestions):
		return apperrors.BadRequest("Not enough published questions to start a game")
	case errors.Is(err, service.ErrAccessDenied):
		return apperrors.Forbidden("Access denied")
	default:
		return apperrors.Internal("Failed to process request").Wrap(err)
	}
}
