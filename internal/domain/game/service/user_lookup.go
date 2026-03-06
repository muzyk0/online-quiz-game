package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/muzyk0/online-quiz-game/internal/domain/user/repository"
)

// UserLookupAdapter adapts the user repository to the UserLookup interface
// needed by GameService.
type UserLookupAdapter struct {
	repo repository.UserRepositoryInterface
}

func NewUserLookupAdapter(repo repository.UserRepositoryInterface) UserLookup {
	return &UserLookupAdapter{repo: repo}
}

// GetLoginByID returns the user's login field.
func (a *UserLookupAdapter) GetLoginByID(ctx context.Context, userID pgtype.UUID) (string, error) {
	user, err := a.repo.GetByID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("get user for login lookup: %w", err)
	}
	return user.Login.String, nil
}
