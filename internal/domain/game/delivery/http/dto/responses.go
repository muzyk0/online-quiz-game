package dto

import (
	"time"

	"github.com/muzyk0/online-quiz-game/internal/domain/game/service"
)

// PlayerInfoResponse is the player identity in a game view.
type PlayerInfoResponse struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}

// AnswerViewResponse is one submitted answer as seen in the game view.
type AnswerViewResponse struct {
	QuestionID   string    `json:"questionId"`
	AnswerStatus string    `json:"answerStatus"`
	AddedAt      time.Time `json:"addedAt"`
}

// AnswerSubmitResponse is the spec-compliant response for POST /my-current/answers.
type AnswerSubmitResponse struct {
	QuestionID   string    `json:"questionId"`
	AnswerStatus string    `json:"answerStatus"`
	AddedAt      time.Time `json:"addedAt"`
}

// PlayerProgressResponse holds a player's answers and score.
type PlayerProgressResponse struct {
	Answers []*AnswerViewResponse `json:"answers"`
	Player  PlayerInfoResponse    `json:"player"`
	Score   int                   `json:"score"`
}

// QuestionViewResponse is a question visible during an active game (no correct answers).
type QuestionViewResponse struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

// GameResponse is the full game view returned by all game endpoints.
type GameResponse struct {
	ID                   string                  `json:"id"`
	FirstPlayerProgress  *PlayerProgressResponse `json:"firstPlayerProgress"`
	SecondPlayerProgress *PlayerProgressResponse `json:"secondPlayerProgress"`
	Questions            []*QuestionViewResponse `json:"questions"`
	Status               string                  `json:"status"`
	PairCreatedDate      time.Time               `json:"pairCreatedDate"`
	StartGameDate        *time.Time              `json:"startGameDate"`
	FinishGameDate       *time.Time              `json:"finishGameDate"`
}

// PaginatedGamesResponse is the paginated list of games.
type PaginatedGamesResponse struct {
	PagesCount int             `json:"pagesCount"`
	Page       int             `json:"page"`
	PageSize   int             `json:"pageSize"`
	TotalCount int             `json:"totalCount"`
	Items      []*GameResponse `json:"items"`
}

// StatisticResponse is the current user's statistic across finished games.
type StatisticResponse struct {
	SumScore    int     `json:"sumScore"`
	AvgScores   float64 `json:"avgScores"`
	GamesCount  int     `json:"gamesCount"`
	WinsCount   int     `json:"winsCount"`
	LossesCount int     `json:"lossesCount"`
	DrawsCount  int     `json:"drawsCount"`
}

// ErrorResponse is a generic error body.
type ErrorResponse struct {
	Error string `json:"error"`
}

// FromServiceView converts a service GameView to GameResponse.
func FromServiceView(v *service.GameView) *GameResponse {
	r := &GameResponse{
		ID:             v.ID,
		Status:         v.Status,
		PairCreatedDate: v.PairCreatedDate,
		StartGameDate:  v.StartGameDate,
		FinishGameDate: v.FinishGameDate,
	}

	if v.FirstPlayerProgress != nil {
		r.FirstPlayerProgress = toProgressResponse(v.FirstPlayerProgress)
	}
	if v.SecondPlayerProgress != nil {
		r.SecondPlayerProgress = toProgressResponse(v.SecondPlayerProgress)
	}
	if v.Questions != nil {
		r.Questions = toQuestionResponses(v.Questions)
	}
	return r
}

func toProgressResponse(p *service.PlayerProgress) *PlayerProgressResponse {
	answers := make([]*AnswerViewResponse, len(p.Answers))
	for i, a := range p.Answers {
		answerStatus := "Incorrect"
		if a.IsCorrect {
			answerStatus = "Correct"
		}
		answers[i] = &AnswerViewResponse{
			QuestionID:   a.QuestionID,
			AnswerStatus: answerStatus,
			AddedAt:      a.AddedAt,
		}
	}
	return &PlayerProgressResponse{
		Answers: answers,
		Player:  PlayerInfoResponse{ID: p.Player.ID, Login: p.Player.Login},
		Score:   p.Score,
	}
}

func toQuestionResponses(qs []*service.QuestionView) []*QuestionViewResponse {
	out := make([]*QuestionViewResponse, len(qs))
	for i, q := range qs {
		out[i] = &QuestionViewResponse{ID: q.ID, Body: q.Body}
	}
	return out
}
