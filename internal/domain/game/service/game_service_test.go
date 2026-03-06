package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	gamemodels "github.com/muzyk0/online-quiz-game/internal/domain/game/models"
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
	playerUUID := pgUUID(t, playerID)
	pendingGameID := pgUUID(t, testUUID())

	gameRepo := &GameRepositoryInterfaceMock{
		IsPlayerInActiveGameFunc: func(context.Context, pgtype.UUID) (bool, error) {
			return false, nil
		},
		FindPendingFunc: func(context.Context) (*gamemodels.QuizGame, error) {
			return &gamemodels.QuizGame{
				ID:            pendingGameID,
				FirstPlayerID: playerUUID,
				Status:        gamemodels.GameStatusPending,
			}, nil
		},
		ActivateGameWithQuestionsFunc: func(context.Context, pgtype.UUID, pgtype.UUID, []pgtype.UUID) (*gamemodels.QuizGame, error) {
			t.Fatal("pending game should not be activated when fewer than five questions exist")
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
	playerUUID := pgUUID(t, playerID)
	pendingGameID := pgUUID(t, testUUID())
	expectedErr := errors.New("write failed")

	gameRepo := &GameRepositoryInterfaceMock{
		IsPlayerInActiveGameFunc: func(context.Context, pgtype.UUID) (bool, error) {
			return false, nil
		},
		FindPendingFunc: func(context.Context) (*gamemodels.QuizGame, error) {
			return &gamemodels.QuizGame{
				ID:            pendingGameID,
				FirstPlayerID: playerUUID,
				Status:        gamemodels.GameStatusPending,
			}, nil
		},
		ActivateGameWithQuestionsFunc: func(context.Context, pgtype.UUID, pgtype.UUID, []pgtype.UUID) (*gamemodels.QuizGame, error) {
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
