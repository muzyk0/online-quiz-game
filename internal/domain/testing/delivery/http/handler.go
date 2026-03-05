package http

import (
	nethttp "net/http"

	"github.com/labstack/echo/v4"
	"github.com/muzyk0/online-quiz-game/internal/app/database"
	"github.com/muzyk0/online-quiz-game/internal/pkg/apperrors"
)

// TestingHandler handles testing utility endpoints.
type TestingHandler struct {
	db *database.DB
}

// NewHandler creates a new TestingHandler.
func NewHandler(db *database.DB) *TestingHandler {
	return &TestingHandler{db: db}
}

// DeleteAllData godoc
//
//	@Summary		Delete all data
//	@Description	Truncates all application tables. For testing purposes only.
//	@Tags			Testing
//	@Success		204	"All data deleted"
//	@Failure		500	{object}	apperrors.AppError
//	@Router			/api/testing/all-data [delete]
//
// DeleteAllData truncates all application tables and returns 204.
func (h *TestingHandler) DeleteAllData(c echo.Context) error {
	_, err := h.db.ExecContext(c.Request().Context(),
		`TRUNCATE TABLE quiz_game_answers, quiz_game_questions, quiz_games, quiz_questions, users CASCADE`)
	if err != nil {
		return apperrors.New(nethttp.StatusInternalServerError, "failed to clear data").Wrap(err)
	}
	return c.NoContent(nethttp.StatusNoContent)
}
