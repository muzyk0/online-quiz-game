package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/muzyk0/online-quiz-game/internal/app/database"
	"github.com/muzyk0/online-quiz-game/internal/domain/game/models"
)

var (
	ErrGameNotFound    = errors.New("game not found")
	ErrAnswerDuplicate = errors.New("answer already submitted for this question")
)

// GameListFilter holds pagination and sorting options for listing games.
type GameListFilter struct {
	SortBy        string // "pairCreatedDate" | "status" | "finishGameDate" | "startGameDate"
	SortDirection string // "asc" | "desc"
	PageNumber    int
	PageSize      int
}

// PlayerStats holds aggregated statistics for a player across finished games.
type PlayerStats struct {
	GamesCount  int `db:"games_count"`
	SumScore    int `db:"sum_score"`
	WinsCount   int `db:"wins_count"`
	LossesCount int `db:"losses_count"`
	DrawsCount  int `db:"draws_count"`
}

//go:generate go run github.com/matryer/moq@latest -out ../service/mock_game_repository_test.go -pkg service . GameRepositoryInterface

// GameRepositoryInterface defines all DB operations needed by the game service.
type GameRepositoryInterface interface {
	// Game lifecycle
	CreatePending(ctx context.Context, firstPlayerID pgtype.UUID) (*models.QuizGame, error)
	// FindPendingAndActivate atomically finds the oldest pending game (with FOR UPDATE SKIP LOCKED),
	// activates it, and assigns the given questions — all in one transaction.
	// Returns nil, nil when no pending game is available.
	FindPendingAndActivate(ctx context.Context, secondPlayerID pgtype.UUID, questionIDs []pgtype.UUID) (*models.QuizGame, error)
	GetByID(ctx context.Context, id pgtype.UUID) (*models.QuizGame, error)
	GetActiveByPlayerID(ctx context.Context, playerID pgtype.UUID) (*models.QuizGame, error)
	IsPlayerInActiveGame(ctx context.Context, playerID pgtype.UUID) (bool, error)
	UpdateScoresAndFinish(ctx context.Context, g *models.QuizGame) error

	// History & stats
	GetAllByPlayerID(ctx context.Context, playerID pgtype.UUID, filter GameListFilter) ([]*models.QuizGame, int, error)
	GetStatsByPlayerID(ctx context.Context, playerID pgtype.UUID) (*PlayerStats, error)

	// Questions
	AssignQuestions(ctx context.Context, gameID pgtype.UUID, questionIDs []pgtype.UUID) error
	GetGameQuestions(ctx context.Context, gameID pgtype.UUID) ([]*models.QuizGameQuestion, error)

	// Answers
	SaveAnswer(ctx context.Context, a models.QuizGameAnswer) (*models.QuizGameAnswer, error)
	GetPlayerAnswers(ctx context.Context, gameID, playerID pgtype.UUID) ([]*models.QuizGameAnswer, error)
	CountPlayerAnswers(ctx context.Context, gameID, playerID pgtype.UUID) (int, error)
}

// GameRepository implements GameRepositoryInterface.
type GameRepository struct {
	db *database.DB
}

func NewGameRepository(db *database.DB) GameRepositoryInterface {
	return &GameRepository{db: db}
}

const gameColumns = `id, first_player_id, second_player_id, first_player_score, second_player_score,
	first_player_finished_at, second_player_finished_at, status, created_at, started_at, finished_at`

func (r *GameRepository) CreatePending(ctx context.Context, firstPlayerID pgtype.UUID) (*models.QuizGame, error) {
	query := `
		INSERT INTO quiz_games (first_player_id, status)
		VALUES ($1, 'PendingSecondPlayer')
		RETURNING ` + gameColumns

	var g models.QuizGame
	if err := r.db.GetContext(ctx, &g, query, firstPlayerID); err != nil {
		return nil, fmt.Errorf("create pending game: %w", err)
	}
	return &g, nil
}

func (r *GameRepository) FindPending(ctx context.Context) (*models.QuizGame, error) {
	query := `SELECT ` + gameColumns + ` FROM quiz_games WHERE status = 'PendingSecondPlayer' ORDER BY created_at ASC LIMIT 1`

	var g models.QuizGame
	if err := r.db.GetContext(ctx, &g, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // no pending game
		}
		return nil, fmt.Errorf("find pending game: %w", err)
	}
	return &g, nil
}

func (r *GameRepository) ActivateGame(ctx context.Context, gameID, secondPlayerID pgtype.UUID) (*models.QuizGame, error) {
	query := `
		UPDATE quiz_games
		SET second_player_id = $2, status = 'Active', started_at = NOW()
		WHERE id = $1 AND status = 'PendingSecondPlayer'
		RETURNING ` + gameColumns

	var g models.QuizGame
	if err := r.db.GetContext(ctx, &g, query, gameID, secondPlayerID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrGameNotFound
		}
		return nil, fmt.Errorf("activate game: %w", err)
	}
	return &g, nil
}

func (r *GameRepository) ActivateGameWithQuestions(
	ctx context.Context,
	gameID, secondPlayerID pgtype.UUID,
	questionIDs []pgtype.UUID,
) (*models.QuizGame, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin activate game transaction: %w", err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	query := `
		UPDATE quiz_games
		SET second_player_id = $2, status = 'Active', started_at = NOW()
		WHERE id = $1 AND status = 'PendingSecondPlayer'
		RETURNING ` + gameColumns

	var g models.QuizGame
	if err := tx.GetContext(ctx, &g, query, gameID, secondPlayerID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrGameNotFound
		}
		return nil, fmt.Errorf("activate game: %w", err)
	}

	for i, qID := range questionIDs {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO quiz_game_questions (game_id, question_id, order_index) VALUES ($1, $2, $3)`,
			gameID, qID, i,
		); err != nil {
			return nil, fmt.Errorf("assign question %d: %w", i, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit activate game transaction: %w", err)
	}

	return &g, nil
}

func (r *GameRepository) FindPendingAndActivate(
	ctx context.Context,
	secondPlayerID pgtype.UUID,
	questionIDs []pgtype.UUID,
) (*models.QuizGame, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin find-and-activate transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	// SKIP LOCKED: if another transaction is already activating a pending game,
	// we skip it rather than waiting, so concurrent joiners don't queue up.
	const findQuery = `SELECT ` + gameColumns + `
		FROM quiz_games
		WHERE status = 'PendingSecondPlayer'
		ORDER BY created_at ASC
		LIMIT 1
		FOR UPDATE SKIP LOCKED`

	var pending models.QuizGame
	if err := tx.GetContext(ctx, &pending, findQuery); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // no available pending game
		}
		return nil, fmt.Errorf("find pending game: %w", err)
	}

	// Activate within the same transaction while the row lock is still held.
	const activateQuery = `
		UPDATE quiz_games
		SET second_player_id = $2, status = 'Active', started_at = NOW()
		WHERE id = $1 AND status = 'PendingSecondPlayer'
		RETURNING ` + gameColumns

	var activated models.QuizGame
	if err := tx.GetContext(ctx, &activated, activateQuery, pending.ID, secondPlayerID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrGameNotFound
		}
		return nil, fmt.Errorf("activate game: %w", err)
	}

	for i, qID := range questionIDs {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO quiz_game_questions (game_id, question_id, order_index) VALUES ($1, $2, $3)`,
			pending.ID, qID, i,
		); err != nil {
			return nil, fmt.Errorf("assign question %d: %w", i, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("commit find-and-activate transaction: %w", err)
	}
	return &activated, nil
}

func (r *GameRepository) GetByID(ctx context.Context, id pgtype.UUID) (*models.QuizGame, error) {
	query := `SELECT ` + gameColumns + ` FROM quiz_games WHERE id = $1`

	var g models.QuizGame
	if err := r.db.GetContext(ctx, &g, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrGameNotFound
		}
		return nil, fmt.Errorf("get game by id: %w", err)
	}
	return &g, nil
}

func (r *GameRepository) GetActiveByPlayerID(ctx context.Context, playerID pgtype.UUID) (*models.QuizGame, error) {
	query := `
		SELECT ` + gameColumns + `
		FROM quiz_games
		WHERE status IN ('PendingSecondPlayer', 'Active')
		  AND (first_player_id = $1 OR second_player_id = $1)
		ORDER BY created_at DESC
		LIMIT 1`

	var g models.QuizGame
	if err := r.db.GetContext(ctx, &g, query, playerID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrGameNotFound
		}
		return nil, fmt.Errorf("get active game by player: %w", err)
	}
	return &g, nil
}

func (r *GameRepository) IsPlayerInActiveGame(ctx context.Context, playerID pgtype.UUID) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM quiz_games
		WHERE status IN ('PendingSecondPlayer', 'Active')
		  AND (first_player_id = $1 OR second_player_id = $1)`

	var count int
	if err := r.db.GetContext(ctx, &count, query, playerID); err != nil {
		return false, fmt.Errorf("check player in active game: %w", err)
	}
	return count > 0, nil
}

func (r *GameRepository) UpdateScoresAndFinish(ctx context.Context, g *models.QuizGame) error {
	query := `
		UPDATE quiz_games
		SET first_player_score          = $2,
		    second_player_score         = $3,
		    first_player_finished_at    = $4,
		    second_player_finished_at   = $5,
		    status                      = $6,
		    finished_at                 = $7
		WHERE id = $1`

	res, err := r.db.ExecContext(ctx, query,
		g.ID,
		g.FirstPlayerScore,
		g.SecondPlayerScore,
		g.FirstPlayerFinishedAt,
		g.SecondPlayerFinishedAt,
		g.Status,
		g.FinishedAt,
	)
	if err != nil {
		return fmt.Errorf("update game scores: %w", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("update game scores rows affected: %w", err)
	}
	if affected == 0 {
		return ErrGameNotFound
	}
	return nil
}

func (r *GameRepository) AssignQuestions(ctx context.Context, gameID pgtype.UUID, questionIDs []pgtype.UUID) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin assign questions transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	for i, qID := range questionIDs {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO quiz_game_questions (game_id, question_id, order_index) VALUES ($1, $2, $3)`,
			gameID, qID, i,
		); err != nil {
			return fmt.Errorf("assign question %d: %w", i, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit assign questions transaction: %w", err)
	}
	return nil
}

func (r *GameRepository) GetGameQuestions(ctx context.Context, gameID pgtype.UUID) ([]*models.QuizGameQuestion, error) {
	query := `
		SELECT gq.id, gq.game_id, gq.question_id, gq.order_index, q.body
		FROM quiz_game_questions gq
		JOIN quiz_questions q ON q.id = gq.question_id
		WHERE gq.game_id = $1
		ORDER BY gq.order_index`

	var questions []*models.QuizGameQuestion
	if err := r.db.SelectContext(ctx, &questions, query, gameID); err != nil {
		return nil, fmt.Errorf("get game questions: %w", err)
	}
	return questions, nil
}

func (r *GameRepository) SaveAnswer(ctx context.Context, a models.QuizGameAnswer) (*models.QuizGameAnswer, error) {
	query := `
		INSERT INTO quiz_game_answers (game_id, player_id, question_id, answer, is_correct)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, game_id, player_id, question_id, answer, is_correct, answered_at`

	var out models.QuizGameAnswer
	if err := r.db.GetContext(ctx, &out, query, a.GameID, a.PlayerID, a.QuestionID, a.Answer, a.IsCorrect); err != nil {
		return nil, fmt.Errorf("save answer: %w", err)
	}
	return &out, nil
}

func (r *GameRepository) GetPlayerAnswers(ctx context.Context, gameID, playerID pgtype.UUID) ([]*models.QuizGameAnswer, error) {
	query := `
		SELECT id, game_id, player_id, question_id, answer, is_correct, answered_at
		FROM quiz_game_answers
		WHERE game_id = $1 AND player_id = $2
		ORDER BY answered_at`

	var answers []*models.QuizGameAnswer
	if err := r.db.SelectContext(ctx, &answers, query, gameID, playerID); err != nil {
		return nil, fmt.Errorf("get player answers: %w", err)
	}
	return answers, nil
}

func (r *GameRepository) CountPlayerAnswers(ctx context.Context, gameID, playerID pgtype.UUID) (int, error) {
	var count int
	if err := r.db.GetContext(ctx, &count,
		`SELECT COUNT(*) FROM quiz_game_answers WHERE game_id = $1 AND player_id = $2`,
		gameID, playerID,
	); err != nil {
		return 0, fmt.Errorf("count player answers: %w", err)
	}
	return count, nil
}

func (r *GameRepository) GetAllByPlayerID(ctx context.Context, playerID pgtype.UUID, filter GameListFilter) ([]*models.QuizGame, int, error) {
	const where = `WHERE first_player_id = $1 OR second_player_id = $1`

	var total int
	if err := r.db.GetContext(ctx, &total, `SELECT COUNT(*) FROM quiz_games `+where, playerID); err != nil {
		return nil, 0, fmt.Errorf("count player games: %w", err)
	}

	pageSize := filter.PageSize
	if pageSize < 1 || pageSize > 20 {
		pageSize = 10
	}
	pageNumber := filter.PageNumber
	if pageNumber < 1 {
		pageNumber = 1
	}
	offset := (pageNumber - 1) * pageSize

	sortCol := validGameSortColumn(filter.SortBy)
	sortDir := validSortDir(filter.SortDirection)

	query := fmt.Sprintf(
		`SELECT %s FROM quiz_games %s ORDER BY %s %s, created_at DESC LIMIT $2 OFFSET $3`,
		gameColumns, where, sortCol, sortDir,
	)

	var games []*models.QuizGame
	if err := r.db.SelectContext(ctx, &games, query, playerID, pageSize, offset); err != nil {
		return nil, 0, fmt.Errorf("list player games: %w", err)
	}
	return games, total, nil
}

func (r *GameRepository) GetStatsByPlayerID(ctx context.Context, playerID pgtype.UUID) (*PlayerStats, error) {
	query := `
		SELECT
			COUNT(*) AS games_count,
			COALESCE(SUM(CASE WHEN first_player_id = $1 THEN first_player_score ELSE second_player_score END), 0) AS sum_score,
			COALESCE(SUM(CASE WHEN
				(first_player_id = $1 AND first_player_score > second_player_score) OR
				(second_player_id = $1 AND second_player_score > first_player_score)
			THEN 1 ELSE 0 END), 0) AS wins_count,
			COALESCE(SUM(CASE WHEN
				(first_player_id = $1 AND first_player_score < second_player_score) OR
				(second_player_id = $1 AND second_player_score < first_player_score)
			THEN 1 ELSE 0 END), 0) AS losses_count,
			COALESCE(SUM(CASE WHEN first_player_score = second_player_score THEN 1 ELSE 0 END), 0) AS draws_count
		FROM quiz_games
		WHERE (first_player_id = $1 OR second_player_id = $1)
		  AND status = 'Finished'`

	var stats PlayerStats
	if err := r.db.GetContext(ctx, &stats, query, playerID); err != nil {
		return nil, fmt.Errorf("get player stats: %w", err)
	}
	return &stats, nil
}

func validGameSortColumn(s string) string {
	switch s {
	case "status":
		return "status"
	case "finishGameDate":
		return "finished_at"
	case "startGameDate":
		return "started_at"
	default: // "pairCreatedDate" and fallback
		return "created_at"
	}
}

func validSortDir(s string) string {
	if strings.ToLower(s) == "asc" {
		return "ASC"
	}
	return "DESC"
}
