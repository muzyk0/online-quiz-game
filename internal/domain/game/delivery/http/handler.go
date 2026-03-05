package http

import (
	nethttp "net/http"

	"github.com/labstack/echo/v4"

	"github.com/muzyk0/online-quiz-game/internal/domain/game/delivery/http/dto"
	"github.com/muzyk0/online-quiz-game/internal/domain/game/service"
	"github.com/muzyk0/online-quiz-game/internal/pkg/auth"
)

// Handler handles HTTP requests for the quiz pair game.
type Handler struct {
	svc service.GameServiceInterface
}

func NewHandler(svc service.GameServiceInterface) *Handler {
	return &Handler{svc: svc}
}

// Connect godoc
//
//	@Summary		Join or create a quiz game
//	@Description	If a pending game exists the calling user is added as second player and the game starts. Otherwise a new pending game is created.
//	@Tags			Quiz/Game
//	@Produce		json
//	@Success		200	{object}	dto.GameResponse
//	@Failure		403	{object}	dto.ErrorResponse	"Already in an active game"
//	@Security		BearerAuth
//	@Router			/api/pair-game-quiz/pairs/connection [post]
func (h *Handler) Connect(c echo.Context) error {
	playerID := auth.MustGetUserID(c)

	view, err := h.svc.JoinOrCreateGame(c.Request().Context(), playerID)
	if err != nil {
		return mapGameError(err)
	}
	return c.JSON(nethttp.StatusOK, dto.FromServiceView(view))
}

// GetMyCurrent godoc
//
//	@Summary		Get the calling user's current active game
//	@Description	Returns the game in status PendingSecondPlayer or Active. Returns 404 when no active game.
//	@Tags			Quiz/Game
//	@Produce		json
//	@Success		200	{object}	dto.GameResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Security		BearerAuth
//	@Router			/api/pair-game-quiz/pairs/my-current [get]
func (h *Handler) GetMyCurrent(c echo.Context) error {
	playerID := auth.MustGetUserID(c)

	view, err := h.svc.GetMyCurrentGame(c.Request().Context(), playerID)
	if err != nil {
		return mapGameError(err)
	}
	return c.JSON(nethttp.StatusOK, dto.FromServiceView(view))
}

// GetByID godoc
//
//	@Summary		Get a game by ID
//	@Description	Returns the game in any status. The calling user must be a participant.
//	@Tags			Quiz/Game
//	@Produce		json
//	@Param			id	path		string	true	"Game ID"
//	@Success		200	{object}	dto.GameResponse
//	@Failure		403	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Security		BearerAuth
//	@Router			/api/pair-game-quiz/pairs/{id} [get]
func (h *Handler) GetByID(c echo.Context) error {
	playerID := auth.MustGetUserID(c)
	gameID := c.Param("id")

	view, err := h.svc.GetGameByID(c.Request().Context(), gameID, playerID)
	if err != nil {
		return mapGameError(err)
	}
	return c.JSON(nethttp.StatusOK, dto.FromServiceView(view))
}

// SubmitAnswer godoc
//
//	@Summary		Submit an answer for the current question
//	@Description	Submits an answer to the next unanswered question in the current active game. One attempt per question. Sequential order.
//	@Tags			Quiz/Game
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.AnswerRequest	true	"Answer"
//	@Success		200		{object}	dto.GameResponse
//	@Failure		403		{object}	dto.ErrorResponse	"No active game or all questions answered"
//	@Security		BearerAuth
//	@Router			/api/pair-game-quiz/pairs/my-current/answers [post]
func (h *Handler) SubmitAnswer(c echo.Context) error {
	playerID := auth.MustGetUserID(c)

	var req dto.AnswerRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(nethttp.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	result, err := h.svc.SubmitAnswer(c.Request().Context(), playerID, req.Answer)
	if err != nil {
		return mapGameError(err)
	}
	answerStatus := "Incorrect"
	if result.IsCorrect {
		answerStatus = "Correct"
	}
	return c.JSON(nethttp.StatusOK, dto.AnswerSubmitResponse{
		QuestionID:   result.QuestionID,
		AnswerStatus: answerStatus,
		AddedAt:      result.AddedAt,
	})
}
