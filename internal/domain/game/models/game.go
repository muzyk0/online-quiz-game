package models

import "github.com/jackc/pgx/v5/pgtype"

// GameStatus represents the lifecycle state of a quiz game.
type GameStatus string

const (
	GameStatusPending  GameStatus = "PendingSecondPlayer"
	GameStatusActive   GameStatus = "Active"
	GameStatusFinished GameStatus = "Finished"
)

// QuizGame is a pair session between two registered users.
type QuizGame struct {
	ID                       pgtype.UUID        `db:"id"`
	FirstPlayerID            pgtype.UUID        `db:"first_player_id"`
	SecondPlayerID           pgtype.UUID        `db:"second_player_id"`
	FirstPlayerScore         int                `db:"first_player_score"`
	SecondPlayerScore        int                `db:"second_player_score"`
	FirstPlayerFinishedAt    pgtype.Timestamptz `db:"first_player_finished_at"`
	SecondPlayerFinishedAt   pgtype.Timestamptz `db:"second_player_finished_at"`
	Status                   GameStatus         `db:"status"`
	CreatedAt                pgtype.Timestamptz `db:"created_at"`
	StartedAt                pgtype.Timestamptz `db:"started_at"`
	FinishedAt               pgtype.Timestamptz `db:"finished_at"`
}

// QuizGameQuestion links a question to a game at a specific position (0-4).
type QuizGameQuestion struct {
	ID         pgtype.UUID `db:"id"`
	GameID     pgtype.UUID `db:"game_id"`
	QuestionID pgtype.UUID `db:"question_id"`
	OrderIndex int         `db:"order_index"`
	// Joined from quiz_questions
	Body pgtype.Text `db:"body"`
}

// QuizGameAnswer is one answer submitted by a player for one question.
type QuizGameAnswer struct {
	ID         pgtype.UUID        `db:"id"`
	GameID     pgtype.UUID        `db:"game_id"`
	PlayerID   pgtype.UUID        `db:"player_id"`
	QuestionID pgtype.UUID        `db:"question_id"`
	Answer     string             `db:"answer"`
	IsCorrect  bool               `db:"is_correct"`
	AnsweredAt pgtype.Timestamptz `db:"answered_at"`
}
