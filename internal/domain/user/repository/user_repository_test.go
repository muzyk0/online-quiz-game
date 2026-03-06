package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
	"github.com/muzyk0/online-quiz-game/internal/app/database"
	"github.com/muzyk0/online-quiz-game/internal/domain/user/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Integration tests for UserRepository require a real database.
// Unit-level encryption tests were removed as PII encryption was removed from the quiz service.
// Run integration tests with: go test -tags=integration ./...

func TestUserRepository_Placeholder(t *testing.T) {
	t.Skip("integration tests require a database connection")
}

func TestUserRepositoryUpdatePersistsPasswordHash(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	dbWrapper := &database.DB{DB: sqlxDB}
	repo := &UserRepository{db: dbWrapper}

	userID := uuid.NewString()
	user := models.User{
		ID:           mustPGUUID(t, userID),
		Email:        "user@example.com",
		PasswordHash: pgtype.Text{String: "hashed-password", Valid: true},
		FirstName:    pgtype.Text{String: "Jane", Valid: true},
		LastName:     pgtype.Text{String: "Doe", Valid: true},
		AvatarUrl:    pgtype.Text{String: "https://example.com/avatar.png", Valid: true},
		IsVerified:   pgtype.Bool{Bool: true, Valid: true},
	}

	rows := sqlmock.NewRows([]string{
		"id", "login", "email", "password_hash", "first_name", "last_name", "avatar_url", "is_verified", "created_at", "updated_at",
	}).AddRow(
		userID,
		"janedoe",
		user.Email,
		user.PasswordHash.String,
		user.FirstName.String,
		user.LastName.String,
		user.AvatarUrl.String,
		true,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(regexp.QuoteMeta(`
		UPDATE users SET
			email      = $1,
			password_hash = $2,
			first_name = $3,
			last_name  = $4,
			avatar_url = $5,
			is_verified = $6,
			updated_at = NOW()
		WHERE id = $7
		RETURNING `+userColumns)).
		WithArgs(
			user.Email,
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
		).
		WillReturnRows(rows)

	updated, err := repo.Update(t.Context(), user)

	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, user.PasswordHash.String, updated.PasswordHash.String)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func mustPGUUID(t *testing.T, id string) pgtype.UUID {
	t.Helper()
	var uid pgtype.UUID
	require.NoError(t, uid.Scan(id))
	return uid
}
