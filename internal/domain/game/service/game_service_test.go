package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	gamemodels "github.com/muzyk0/online-quiz-game/internal/domain/game/models"
	gamerepo "github.com/muzyk0/online-quiz-game/internal/domain/game/repository"
	questionmodels "github.com/muzyk0/online-quiz-game/internal/domain/question/models"
	questionrepo "github.com/muzyk0/online-quiz-game/internal/domain/question/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubUserLookup struct{}

func (stubUserLookup) GetLoginByID(context.Context, pgtype.UUID) (string, error) {
	return "player", nil
}

type questionRepoStub struct {
	listPublishedFunc func(context.Context, int) ([]*questionmodels.QuizQuestion, error)
}

var _ questionrepo.QuestionRepositoryInterface = (*questionRepoStub)(nil)

func testUUID() string {
	return uuid.NewString()
}

func pgUUID(t *testing.T, id string) pgtype.UUID {
	t.Helper()
	var uid pgtype.UUID
	require.NoError(t, uid.Scan(id))
	return uid
}

func (s *questionRepoStub) Create(context.Context, questionmodels.QuizQuestion) (*questionmodels.QuizQuestion, error) {
	panic("not used in test")
}

func (s *questionRepoStub) GetByID(context.Context, pgtype.UUID) (*questionmodels.QuizQuestion, error) {
	panic("not used in test")
}

func (s *questionRepoStub) Update(context.Context, questionmodels.QuizQuestion) (*questionmodels.QuizQuestion, error) {
	panic("not used in test")
}

func (s *questionRepoStub) Delete(context.Context, pgtype.UUID) error {
	panic("not used in test")
}

func (s *questionRepoStub) SetPublished(context.Context, pgtype.UUID, bool) (*questionmodels.QuizQuestion, error) {
	panic("not used in test")
}

func (s *questionRepoStub) List(context.Context, questionrepo.ListFilter) ([]*questionmodels.QuizQuestion, int, error) {
	panic("not used in test")
}

func (s *questionRepoStub) ListPublished(ctx context.Context, limit int) ([]*questionmodels.QuizQuestion, error) {
	return s.listPublishedFunc(ctx, limit)
}

func TestGameServiceJoinOrCreateGameDoesNotActivateWithoutEnoughQuestions(t *testing.T) {
	playerID := testUUID()

	gameRepo := &GameRepositoryInterfaceMock{
		GetActiveByPlayerIDFunc: func(context.Context, pgtype.UUID) (*gamemodels.QuizGame, error) {
			return nil, gamerepo.ErrGameNotFound
		},
		FindPendingAndActivateFunc: func(context.Context, pgtype.UUID, []pgtype.UUID) (*gamemodels.QuizGame, error) {
			t.Fatal("FindPendingAndActivate should not be called when fewer than five questions exist")
			return nil, nil
		},
	}
	questionRepo := &questionRepoStub{
		listPublishedFunc: func(context.Context, int) ([]*questionmodels.QuizQuestion, error) {
			return []*questionmodels.QuizQuestion{
				{ID: pgUUID(t, testUUID())},
				{ID: pgUUID(t, testUUID())},
				{ID: pgUUID(t, testUUID())},
				{ID: pgUUID(t, testUUID())},
			}, nil
		},
	}

	svc := NewGameService(gameRepo, questionRepo, stubUserLookup{})

	_, err := svc.JoinOrCreateGame(context.Background(), playerID)

	require.Error(t, err)
	assert.ErrorIs(t, err, ErrNotEnoughQuestions)
}

func TestGameServiceJoinOrCreateGamePropagatesAtomicActivationError(t *testing.T) {
	playerID := testUUID()
	expectedErr := errors.New("write failed")

	gameRepo := &GameRepositoryInterfaceMock{
		GetActiveByPlayerIDFunc: func(context.Context, pgtype.UUID) (*gamemodels.QuizGame, error) {
			return nil, gamerepo.ErrGameNotFound
		},
		FindPendingAndActivateFunc: func(context.Context, pgtype.UUID, []pgtype.UUID) (*gamemodels.QuizGame, error) {
			return nil, expectedErr
		},
	}
	questionRepo := &questionRepoStub{
		listPublishedFunc: func(context.Context, int) ([]*questionmodels.QuizQuestion, error) {
			return []*questionmodels.QuizQuestion{
				{ID: pgUUID(t, testUUID())},
				{ID: pgUUID(t, testUUID())},
				{ID: pgUUID(t, testUUID())},
				{ID: pgUUID(t, testUUID())},
				{ID: pgUUID(t, testUUID())},
			}, nil
		},
	}

	svc := NewGameService(gameRepo, questionRepo, stubUserLookup{})

	_, err := svc.JoinOrCreateGame(context.Background(), playerID)

	require.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
}

func TestGetTopPlayersReturnsPagedResults(t *testing.T) {
	gameRepo := &GameRepositoryInterfaceMock{
		GetTopPlayersFunc: func(_ context.Context, filter gamerepo.TopPlayersFilter) ([]*gamerepo.TopPlayerStats, int, error) {
			return []*gamerepo.TopPlayerStats{
				{PlayerID: "id-1", PlayerLogin: "alice", GamesCount: 3, SumScore: 9, AvgScores: 3.0, WinsCount: 3, LossesCount: 0, DrawsCount: 0},
				{PlayerID: "id-2", PlayerLogin: "bob", GamesCount: 2, SumScore: 4, AvgScores: 2.0, WinsCount: 1, LossesCount: 1, DrawsCount: 0},
			}, 2, nil
		},
	}

	svc := NewGameService(gameRepo, &questionRepoStub{}, stubUserLookup{})

	result, err := svc.GetTopPlayers(context.Background(), TopPlayersInput{
		Sort:       []string{"avgScores desc"},
		PageNumber: 1,
		PageSize:   10,
	})

	require.NoError(t, err)
	require.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Items, 2)
	assert.Equal(t, "alice", result.Items[0].Player.Login)
	assert.Equal(t, 3, result.Items[0].WinsCount)
	assert.InDelta(t, 3.0, result.Items[0].AvgScores, 0.001)
}

func TestGetTopPlayersDefaultSortApplied(t *testing.T) {
	var capturedFilter gamerepo.TopPlayersFilter
	gameRepo := &GameRepositoryInterfaceMock{
		GetTopPlayersFunc: func(_ context.Context, filter gamerepo.TopPlayersFilter) ([]*gamerepo.TopPlayerStats, int, error) {
			capturedFilter = filter
			return nil, 0, nil
		},
	}

	svc := NewGameService(gameRepo, &questionRepoStub{}, stubUserLookup{})

	_, err := svc.GetTopPlayers(context.Background(), TopPlayersInput{})
	require.NoError(t, err)

	// Empty sort slice is passed through to repo; repo applies default SQL ordering
	assert.Empty(t, capturedFilter.Sort)
	assert.Equal(t, 1, capturedFilter.PageNumber)
	assert.Equal(t, 10, capturedFilter.PageSize)
}
