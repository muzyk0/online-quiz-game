package http

import (
	"errors"

	"github.com/muzyk0/online-quiz-game/internal/domain/question/service"
	"github.com/muzyk0/online-quiz-game/internal/platform/http/apperrors"
)

func mapQuestionError(err error) error {
	switch {
	case errors.Is(err, service.ErrQuestionNotFound):
		return apperrors.NotFound("Question not found")
	case errors.Is(err, service.ErrBodyRequired):
		return apperrors.BadRequest("Body is required")
	case errors.Is(err, service.ErrCorrectAnswersRequired):
		return apperrors.BadRequest("Cannot publish question without correct answers")
	default:
		return apperrors.Internal("Failed to process request").Wrap(err)
	}
}
