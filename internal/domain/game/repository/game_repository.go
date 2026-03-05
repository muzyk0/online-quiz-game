package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/muzyk0/online-quiz-game/internal/app/database"
	"github.com/muzyk0/online-quiz-game/internal/domain/game/models"
)

var (
	ErrGameNotFound    = errors.New("game not found")
	ErrAnswerDuplicate = errors.New("answer already submitted for this question")
)

//go:generate go run github.com/matryer/moq@latest -out ../service/mock_game_repository_test.go -pkg service . GameRepositoryInterface

// GameRepositoryInterface defines all DB operations needed by the game service.
type GameRepositoryInterface interface {
	// Game lifecycle
	CreatePending(ctx context.Context, firstPlayerID pgtype.UUID) (*models.QuizGame, error)
	FindPending(ctx context.Context) (*models.QuizGame, error)
	ActivateGame(ctx context.Context, gameID, secondPlayerID pgtype.UUID) (*models.QuizGame, error)
	GetByID(ctx context.Context, id pgtype.UUID) (*models.QuizGame, error)
	GetActiveByPlayerID(ctx context.Context, playerID pgtype.UUID) (*models.QuizGame, error)
	IsPlayerInActiveGame(ctx context.Context, playerID pgtype.UUID) (bool, error)
	UpdateScoresAndFinish(ctx context.Context, g *models.QuizGame) error

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

	_, err := r.db.ExecContext(ctx, query,
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
	return nil
}

func (r *GameRepository) AssignQuestions(ctx context.Context, gameID pgtype.UUID, questionIDs []pgtype.UUID) error {
	for i, qID := range questionIDs {
		_, err := r.db.ExecContext(ctx,
			`INSERT INTO quiz_game_questions (game_id, question_id, order_index) VALUES ($1, $2, $3)`,
			gameID, qID, i,
		)
		if err != nil {
			return fmt.Errorf("assign question %d: %w", i, err)
		}
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
