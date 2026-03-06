package http

import (
	nethttp "net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/muzyk0/online-quiz-game/internal/domain/game/delivery/http/dto"
	"github.com/muzyk0/online-quiz-game/internal/domain/game/service"
	"github.com/muzyk0/online-quiz-game/internal/platform/http/auth"
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

// GetMyGames godoc
//
//	@Summary		Get all games of the current user (paginated)
//	@Tags			Quiz/Game
//	@Produce		json
//	@Param			sortBy			query	string	false	"Field to sort by (default: pairCreatedDate)"
//	@Param			sortDirection	query	string	false	"asc | desc (default: desc)"
//	@Param			pageNumber		query	int		false	"Page number (default: 1)"
//	@Param			pageSize		query	int		false	"Page size 1-20 (default: 10)"
//	@Success		200	{object}	dto.PaginatedGamesResponse
//	@Security		BearerAuth
//	@Router			/api/pair-game-quiz/pairs/my [get]
func (h *Handler) GetMyGames(c echo.Context) error {
	playerID := auth.MustGetUserID(c)

	input := service.MyGamesInput{
		SortBy:        queryParamOrDefault(c, "sortBy", "pairCreatedDate"),
		SortDirection: queryParamOrDefault(c, "sortDirection", "desc"),
		PageNumber:    queryIntOrDefault(c, "pageNumber", 1),
		PageSize:      queryIntOrDefault(c, "pageSize", 10),
	}

	result, err := h.svc.GetMyGames(c.Request().Context(), playerID, input)
	if err != nil {
		return mapGameError(err)
	}

	items := make([]*dto.GameResponse, len(result.Items))
	for i, v := range result.Items {
		items[i] = dto.FromServiceView(v)
	}
	return c.JSON(nethttp.StatusOK, dto.PaginatedGamesResponse{
		PagesCount: result.PagesCount,
		Page:       result.Page,
		PageSize:   result.PageSize,
		TotalCount: result.TotalCount,
		Items:      items,
	})
}

// GetMyStatistic godoc
//
//	@Summary		Get current user's game statistics
//	@Tags			Quiz/Game
//	@Produce		json
//	@Success		200	{object}	dto.StatisticResponse
//	@Security		BearerAuth
//	@Router			/api/pair-game-quiz/users/my-statistic [get]
func (h *Handler) GetMyStatistic(c echo.Context) error {
	playerID := auth.MustGetUserID(c)

	stats, err := h.svc.GetMyStatistic(c.Request().Context(), playerID)
	if err != nil {
		return mapGameError(err)
	}
	return c.JSON(nethttp.StatusOK, dto.StatisticResponse{
		SumScore:    stats.SumScore,
		AvgScores:   stats.AvgScores,
		GamesCount:  stats.GamesCount,
		WinsCount:   stats.WinsCount,
		LossesCount: stats.LossesCount,
		DrawsCount:  stats.DrawsCount,
	})
}

// GetTopPlayers godoc
//
//	@Summary		Get top players leaderboard
//	@Description	Returns a paginated list of players sorted by their game statistics.
//	@Tags			Quiz/Game
//	@Produce		json
//	@Param			sort		query	[]string	false	"Sort criteria (e.g. avgScores desc)"
//	@Param			pageNumber	query	int			false	"Page number (default: 1)"
//	@Param			pageSize	query	int			false	"Page size 1-20 (default: 10)"
//	@Success		200	{object}	dto.PaginatedTopPlayersResponse
//	@Router			/api/pair-game-quiz/users/top [get]
func (h *Handler) GetTopPlayers(c echo.Context) error {
	sort := c.QueryParams()["sort"]

	input := service.TopPlayersInput{
		Sort:       sort,
		PageNumber: queryIntOrDefault(c, "pageNumber", 1),
		PageSize:   queryIntOrDefault(c, "pageSize", 10),
	}

	result, err := h.svc.GetTopPlayers(c.Request().Context(), input)
	if err != nil {
		return mapGameError(err)
	}

	items := make([]*dto.TopPlayerResponse, len(result.Items))
	for i, v := range result.Items {
		items[i] = &dto.TopPlayerResponse{
			SumScore:    v.SumScore,
			AvgScores:   v.AvgScores,
			GamesCount:  v.GamesCount,
			WinsCount:   v.WinsCount,
			LossesCount: v.LossesCount,
			DrawsCount:  v.DrawsCount,
			Player:      dto.PlayerInfoResponse{ID: v.Player.ID, Login: v.Player.Login},
		}
	}
	return c.JSON(nethttp.StatusOK, dto.PaginatedTopPlayersResponse{
		PagesCount: result.PagesCount,
		Page:       result.Page,
		PageSize:   result.PageSize,
		TotalCount: result.TotalCount,
		Items:      items,
	})
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

func queryParamOrDefault(c echo.Context, key, def string) string {
	v := c.QueryParam(key)
	if v == "" {
		return def
	}
	return v
}

func queryIntOrDefault(c echo.Context, key string, def int) int {
	v := c.QueryParam(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}
