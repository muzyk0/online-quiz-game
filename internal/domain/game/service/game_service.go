package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	gamemodels "github.com/muzyk0/online-quiz-game/internal/domain/game/models"
	gamerepo "github.com/muzyk0/online-quiz-game/internal/domain/game/repository"
	questionrepo "github.com/muzyk0/online-quiz-game/internal/domain/question/repository"
)

const questionsPerGame = 5

// Sentinel errors
var (
	ErrGameNotFound         = errors.New("game not found")
	ErrInvalidGameID        = errors.New("invalid game id format")
	ErrAlreadyInGame        = errors.New("player is already in an active or pending game")
	ErrGameNotActive        = errors.New("game is not active")
	ErrAllQuestionsAnswered = errors.New("player has already answered all questions")
	ErrNotEnoughQuestions   = errors.New("not enough published questions to start a game")
	ErrAccessDenied         = errors.New("player is not a participant of this game")
)

// AnswerSubmitResult is returned by SubmitAnswer with the spec-required fields.
type AnswerSubmitResult struct {
	QuestionID string
	IsCorrect  bool
	AddedAt    time.Time
}

// MyGamesInput holds pagination/sorting params for GetMyGames.
type MyGamesInput struct {
	SortBy        string
	SortDirection string
	PageNumber    int
	PageSize      int
}

// PaginatedGamesOutput is the paginated result of GetMyGames.
type PaginatedGamesOutput struct {
	PagesCount int
	Page       int
	PageSize   int
	TotalCount int
	Items      []*GameView
}

// StatisticView holds aggregated stats for a player.
type StatisticView struct {
	SumScore    int
	AvgScores   float64
	GamesCount  int
	WinsCount   int
	LossesCount int
	DrawsCount  int
}

// GameServiceInterface defines public operations for quiz game.
type GameServiceInterface interface {
	JoinOrCreateGame(ctx context.Context, playerID string) (*GameView, error)
	GetMyCurrentGame(ctx context.Context, playerID string) (*GameView, error)
	GetGameByID(ctx context.Context, gameID, playerID string) (*GameView, error)
	SubmitAnswer(ctx context.Context, playerID, answer string) (*AnswerSubmitResult, error)
	GetMyGames(ctx context.Context, playerID string, input MyGamesInput) (*PaginatedGamesOutput, error)
	GetMyStatistic(ctx context.Context, playerID string) (*StatisticView, error)
}

// --- Output types ---

type PlayerInfo struct {
	ID    string
	Login string // email used as login
}

type AnswerView struct {
	QuestionID string
	IsCorrect  bool
	AddedAt    time.Time
}

type PlayerProgress struct {
	Answers []*AnswerView
	Player  PlayerInfo
	Score   int
}

type QuestionView struct {
	ID   string
	Body string
}

type GameView struct {
	ID                   string
	FirstPlayerProgress  *PlayerProgress
	SecondPlayerProgress *PlayerProgress // nil when PendingSecondPlayer
	Questions            []*QuestionView // nil when PendingSecondPlayer
	Status               string
	PairCreatedDate      time.Time
	StartGameDate        *time.Time
	FinishGameDate       *time.Time
}

// UserLookup is a minimal user data source injected to avoid circular domain deps.
type UserLookup interface {
	GetLoginByID(ctx context.Context, userID pgtype.UUID) (string, error)
}

// GameService implements GameServiceInterface.
type GameService struct {
	gameRepo     gamerepo.GameRepositoryInterface
	questionRepo questionrepo.QuestionRepositoryInterface
	users        UserLookup
}

func NewGameService(
	gameRepo gamerepo.GameRepositoryInterface,
	questionRepo questionrepo.QuestionRepositoryInterface,
	users UserLookup,
) GameServiceInterface {
	return &GameService{
		gameRepo:     gameRepo,
		questionRepo: questionRepo,
		users:        users,
	}
}

// JoinOrCreateGame handles the "connect" button logic.
// If a pending game exists: join it and activate; otherwise create a new pending game.
func (s *GameService) JoinOrCreateGame(ctx context.Context, playerID string) (*GameView, error) {
	pid, err := parseUUID(playerID)
	if err != nil {
		return nil, ErrAccessDenied
	}

	// Reject if already in an active/pending game
	inGame, err := s.gameRepo.IsPlayerInActiveGame(ctx, pid)
	if err != nil {
		return nil, fmt.Errorf("check active game: %w", err)
	}
	if inGame {
		return nil, ErrAlreadyInGame
	}

	// Load questions upfront so activation can be fully atomic.
	questionIDs, err := s.loadRandomQuestionIDs(ctx)
	if err != nil {
		return nil, err
	}

	// Atomically find a pending game, lock it (SKIP LOCKED), activate it,
	// and assign questions in one transaction. Returns nil if no pending game
	// is available (another player grabbed it or none exists).
	game, err := s.gameRepo.FindPendingAndActivate(ctx, pid, questionIDs)
	if err != nil {
		return nil, fmt.Errorf("find and activate game: %w", err)
	}

	if game != nil {
		return s.buildGameView(ctx, game, pid)
	}

	// No pending game available: create one and wait for a second player.
	game, err = s.gameRepo.CreatePending(ctx, pid)
	if err != nil {
		return nil, fmt.Errorf("create pending game: %w", err)
	}
	return s.buildGameView(ctx, game, pid)
}

// GetMyCurrentGame returns the active/pending game for the calling player.
func (s *GameService) GetMyCurrentGame(ctx context.Context, playerID string) (*GameView, error) {
	pid, err := parseUUID(playerID)
	if err != nil {
		return nil, ErrGameNotFound
	}

	game, err := s.gameRepo.GetActiveByPlayerID(ctx, pid)
	if err != nil {
		if errors.Is(err, gamerepo.ErrGameNotFound) {
			return nil, ErrGameNotFound
		}
		return nil, fmt.Errorf("get active game: %w", err)
	}
	return s.buildGameView(ctx, game, pid)
}

// GetGameByID returns any game by ID, only if the calling player is a participant.
func (s *GameService) GetGameByID(ctx context.Context, gameID, playerID string) (*GameView, error) {
	gid, err := parseUUID(gameID)
	if err != nil {
		return nil, ErrInvalidGameID
	}
	pid, err := parseUUID(playerID)
	if err != nil {
		return nil, ErrAccessDenied
	}

	game, err := s.gameRepo.GetByID(ctx, gid)
	if err != nil {
		if errors.Is(err, gamerepo.ErrGameNotFound) {
			return nil, ErrGameNotFound
		}
		return nil, fmt.Errorf("get game by id: %w", err)
	}

	// Verify player is a participant
	if !uuidEqual(game.FirstPlayerID, pid) && !uuidEqual(game.SecondPlayerID, pid) {
		return nil, ErrAccessDenied
	}

	return s.buildGameView(ctx, game, pid)
}

// SubmitAnswer processes a player's answer to the next question in their current game.
func (s *GameService) SubmitAnswer(ctx context.Context, playerID, answer string) (*AnswerSubmitResult, error) {
	pid, err := parseUUID(playerID)
	if err != nil {
		return nil, ErrAccessDenied
	}

	// Get current active game
	game, err := s.gameRepo.GetActiveByPlayerID(ctx, pid)
	if err != nil {
		if errors.Is(err, gamerepo.ErrGameNotFound) {
			return nil, ErrGameNotActive // player is not in any active game → 403
		}
		return nil, fmt.Errorf("get active game: %w", err)
	}
	if game.Status != gamemodels.GameStatusActive {
		return nil, ErrGameNotActive
	}

	// How many questions has this player answered?
	answerCount, err := s.gameRepo.CountPlayerAnswers(ctx, game.ID, pid)
	if err != nil {
		return nil, fmt.Errorf("count answers: %w", err)
	}
	if answerCount >= questionsPerGame {
		return nil, ErrAllQuestionsAnswered
	}

	// Get the question at position = answerCount
	questions, err := s.gameRepo.GetGameQuestions(ctx, game.ID)
	if err != nil {
		return nil, fmt.Errorf("get questions: %w", err)
	}
	if answerCount >= len(questions) {
		return nil, fmt.Errorf("invalid game state: answered=%d, questions=%d", answerCount, len(questions))
	}
	currentQuestion := questions[answerCount]

	// Fetch correct answers for this question
	qFull, err := s.questionRepo.GetByID(ctx, currentQuestion.QuestionID)
	if err != nil {
		return nil, fmt.Errorf("get question details: %w", err)
	}

	isCorrect := checkAnswer(answer, qFull.CorrectAnswers)

	// Save the answer
	savedAnswer, err := s.gameRepo.SaveAnswer(ctx, gamemodels.QuizGameAnswer{
		GameID:     game.ID,
		PlayerID:   pid,
		QuestionID: currentQuestion.QuestionID,
		Answer:     answer,
		IsCorrect:  isCorrect,
	})
	if err != nil {
		return nil, fmt.Errorf("save answer: %w", err)
	}

	// Determine if this player finished
	newAnswerCount := answerCount + 1
	playerFinished := newAnswerCount == questionsPerGame

	isFirstPlayer := uuidEqual(game.FirstPlayerID, pid)

	if playerFinished {
		now := pgtype.Timestamptz{Time: time.Now(), Valid: true}
		if isFirstPlayer {
			game.FirstPlayerFinishedAt = now
		} else {
			game.SecondPlayerFinishedAt = now
		}
	}

	// Recalculate scores
	game, err = s.recalculateScores(ctx, game)
	if err != nil {
		return nil, err
	}

	// Check if game is over (both players finished)
	if game.FirstPlayerFinishedAt.Valid && game.SecondPlayerFinishedAt.Valid {
		game.Status = gamemodels.GameStatusFinished
		game.FinishedAt = pgtype.Timestamptz{Time: time.Now(), Valid: true}
	}

	// Persist score/status changes
	if err := s.gameRepo.UpdateScoresAndFinish(ctx, game); err != nil {
		return nil, fmt.Errorf("update game: %w", err)
	}

	qidVal, _ := currentQuestion.QuestionID.Value()
	return &AnswerSubmitResult{
		QuestionID: fmt.Sprintf("%v", qidVal),
		IsCorrect:  isCorrect,
		AddedAt:    savedAnswer.AnsweredAt.Time,
	}, nil
}

// GetMyGames returns a paginated list of all games the calling player participated in.
func (s *GameService) GetMyGames(ctx context.Context, playerID string, input MyGamesInput) (*PaginatedGamesOutput, error) {
	pid, err := parseUUID(playerID)
	if err != nil {
		return nil, ErrAccessDenied
	}

	pageSize := input.PageSize
	if pageSize < 1 || pageSize > 20 {
		pageSize = 10
	}
	pageNumber := input.PageNumber
	if pageNumber < 1 {
		pageNumber = 1
	}

	filter := gamerepo.GameListFilter{
		SortBy:        input.SortBy,
		SortDirection: input.SortDirection,
		PageNumber:    pageNumber,
		PageSize:      pageSize,
	}

	games, total, err := s.gameRepo.GetAllByPlayerID(ctx, pid, filter)
	if err != nil {
		return nil, fmt.Errorf("get player games: %w", err)
	}

	items := make([]*GameView, len(games))
	for i, g := range games {
		view, err := s.buildGameView(ctx, g, pid)
		if err != nil {
			return nil, err
		}
		items[i] = view
	}

	pagesCount := 0
	if total > 0 {
		pagesCount = (total + pageSize - 1) / pageSize
	}

	return &PaginatedGamesOutput{
		PagesCount: pagesCount,
		Page:       pageNumber,
		PageSize:   pageSize,
		TotalCount: total,
		Items:      items,
	}, nil
}

// GetMyStatistic returns aggregated stats for the calling player across all finished games.
func (s *GameService) GetMyStatistic(ctx context.Context, playerID string) (*StatisticView, error) {
	pid, err := parseUUID(playerID)
	if err != nil {
		return nil, ErrAccessDenied
	}

	stats, err := s.gameRepo.GetStatsByPlayerID(ctx, pid)
	if err != nil {
		return nil, fmt.Errorf("get player stats: %w", err)
	}

	var avgScores float64
	if stats.GamesCount > 0 {
		avgScores = math.Round(float64(stats.SumScore)/float64(stats.GamesCount)*100) / 100
	}

	return &StatisticView{
		SumScore:    stats.SumScore,
		AvgScores:   avgScores,
		GamesCount:  stats.GamesCount,
		WinsCount:   stats.WinsCount,
		LossesCount: stats.LossesCount,
		DrawsCount:  stats.DrawsCount,
	}, nil
}

// --- internal helpers ---

func (s *GameService) loadRandomQuestionIDs(ctx context.Context) ([]pgtype.UUID, error) {
	questions, err := s.questionRepo.ListPublished(ctx, questionsPerGame)
	if err != nil {
		return nil, fmt.Errorf("list published questions: %w", err)
	}
	if len(questions) < questionsPerGame {
		return nil, ErrNotEnoughQuestions
	}

	ids := make([]pgtype.UUID, questionsPerGame)
	for i, q := range questions[:questionsPerGame] {
		ids[i] = q.ID
	}
	return ids, nil
}

func (s *GameService) recalculateScores(ctx context.Context, game *gamemodels.QuizGame) (*gamemodels.QuizGame, error) {
	fp1Answers, err := s.gameRepo.GetPlayerAnswers(ctx, game.ID, game.FirstPlayerID)
	if err != nil {
		return nil, fmt.Errorf("get first player answers: %w", err)
	}
	fp2Answers, err := s.gameRepo.GetPlayerAnswers(ctx, game.ID, game.SecondPlayerID)
	if err != nil {
		return nil, fmt.Errorf("get second player answers: %w", err)
	}

	p1Score := countCorrect(fp1Answers)
	p2Score := countCorrect(fp2Answers)

	// Apply bonus: player who finished first gets +1 if they have ≥1 correct answer
	// Bonus is only calculated when both players finished
	if game.FirstPlayerFinishedAt.Valid && game.SecondPlayerFinishedAt.Valid {
		if game.FirstPlayerFinishedAt.Time.Before(game.SecondPlayerFinishedAt.Time) {
			if p1Score > 0 {
				p1Score++
			}
		} else if game.SecondPlayerFinishedAt.Time.Before(game.FirstPlayerFinishedAt.Time) {
			if p2Score > 0 {
				p2Score++
			}
		}
	}

	game.FirstPlayerScore = p1Score
	game.SecondPlayerScore = p2Score
	return game, nil
}

func (s *GameService) buildGameView(ctx context.Context, game *gamemodels.QuizGame, currentPlayerID pgtype.UUID) (*GameView, error) {
	view := &GameView{
		Status: string(game.Status),
	}

	id, _ := game.ID.Value()
	view.ID, _ = id.(string)

	if game.CreatedAt.Valid {
		view.PairCreatedDate = game.CreatedAt.Time
	}
	if game.StartedAt.Valid {
		t := game.StartedAt.Time
		view.StartGameDate = &t
	}
	if game.FinishedAt.Valid {
		t := game.FinishedAt.Time
		view.FinishGameDate = &t
	}

	// First player progress
	fp1Login, err := s.users.GetLoginByID(ctx, game.FirstPlayerID)
	if err != nil {
		return nil, fmt.Errorf("get first player login: %w", err)
	}
	fp1ID, _ := game.FirstPlayerID.Value()
	fp1Answers, err := s.gameRepo.GetPlayerAnswers(ctx, game.ID, game.FirstPlayerID)
	if err != nil {
		return nil, fmt.Errorf("get first player answers: %w", err)
	}

	view.FirstPlayerProgress = &PlayerProgress{
		Player:  PlayerInfo{ID: fmt.Sprintf("%v", fp1ID), Login: fp1Login},
		Score:   game.FirstPlayerScore,
		Answers: toAnswerViews(fp1Answers),
	}

	// Second player progress and questions: only when Active or Finished
	if game.Status != gamemodels.GameStatusPending && game.SecondPlayerID.Valid {
		fp2Login, err := s.users.GetLoginByID(ctx, game.SecondPlayerID)
		if err != nil {
			return nil, fmt.Errorf("get second player login: %w", err)
		}
		fp2ID, _ := game.SecondPlayerID.Value()
		fp2Answers, err := s.gameRepo.GetPlayerAnswers(ctx, game.ID, game.SecondPlayerID)
		if err != nil {
			return nil, fmt.Errorf("get second player answers: %w", err)
		}

		view.SecondPlayerProgress = &PlayerProgress{
			Player:  PlayerInfo{ID: fmt.Sprintf("%v", fp2ID), Login: fp2Login},
			Score:   game.SecondPlayerScore,
			Answers: toAnswerViews(fp2Answers),
		}

		// Questions (body only, no correctAnswers for security)
		gameQuestions, err := s.gameRepo.GetGameQuestions(ctx, game.ID)
		if err != nil {
			return nil, fmt.Errorf("get game questions: %w", err)
		}
		view.Questions = toQuestionViews(gameQuestions)
	}

	return view, nil
}

func toAnswerViews(answers []*gamemodels.QuizGameAnswer) []*AnswerView {
	if answers == nil {
		return []*AnswerView{}
	}
	out := make([]*AnswerView, len(answers))
	for i, a := range answers {
		qid, _ := a.QuestionID.Value()
		out[i] = &AnswerView{
			QuestionID: fmt.Sprintf("%v", qid),
			IsCorrect:  a.IsCorrect,
			AddedAt:    a.AnsweredAt.Time,
		}
	}
	return out
}

func toQuestionViews(questions []*gamemodels.QuizGameQuestion) []*QuestionView {
	out := make([]*QuestionView, len(questions))
	for i, q := range questions {
		qid, _ := q.QuestionID.Value()
		out[i] = &QuestionView{
			ID:   fmt.Sprintf("%v", qid),
			Body: q.Body.String,
		}
	}
	return out
}

func countCorrect(answers []*gamemodels.QuizGameAnswer) int {
	n := 0
	for _, a := range answers {
		if a.IsCorrect {
			n++
		}
	}
	return n
}

// checkAnswer checks if the submitted answer matches any of the correct answers
// using case-insensitive comparison.
func checkAnswer(submitted string, correct []string) bool {
	submitted = strings.TrimSpace(strings.ToLower(submitted))
	for _, ca := range correct {
		if strings.TrimSpace(strings.ToLower(ca)) == submitted {
			return true
		}
	}
	return false
}

func parseUUID(s string) (pgtype.UUID, error) {
	var uid pgtype.UUID
	if err := uid.Scan(s); err != nil {
		return uid, err
	}
	return uid, nil
}

func uuidEqual(a, b pgtype.UUID) bool {
	return a.Valid && b.Valid && a.Bytes == b.Bytes
}
