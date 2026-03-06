package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/muzyk0/online-quiz-game/internal/app/database"
	"github.com/muzyk0/online-quiz-game/internal/domain/user/models"
)

// Sentinel errors for user repository
var (
	ErrUserNotFound = errors.New("user not found")
)

//go:generate go run github.com/matryer/moq@latest -out ../service/mock_user_repository_test.go -pkg service . UserRepositoryInterface

// SAListFilter holds options for the SA users list endpoint.
type SAListFilter struct {
	SearchLoginTerm string
	SearchEmailTerm string
	SortBy          string // "login" | "email" | "createdAt"
	SortDirection   string // "asc" | "desc"
	PageNumber      int
	PageSize        int
}

// UserRepositoryInterface defines the interface for user database operations
type UserRepositoryInterface interface {
	Create(ctx context.Context, user models.User) (*models.User, error)
	GetByID(ctx context.Context, id pgtype.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByLogin(ctx context.Context, login string) (*models.User, error)
	Update(ctx context.Context, user models.User) (*models.User, error)
	Delete(ctx context.Context, id pgtype.UUID) error
	ListSA(ctx context.Context, f SAListFilter) ([]*models.User, int, error)
}

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

const userColumns = `id, login, email, password_hash, first_name, last_name, avatar_url, is_verified, created_at, updated_at`

// saUserColumns selects user fields for SA list endpoints, substituting NULL for
// password_hash so the sensitive hash is never fetched over the wire.
const saUserColumns = `id, login, email, NULL AS password_hash, first_name, last_name, avatar_url, is_verified, created_at, updated_at`

// Create inserts a new user into the database
func (r *UserRepository) Create(ctx context.Context, user models.User) (*models.User, error) {
	isVerified := user.IsVerified
	if !isVerified.Valid {
		isVerified = pgtype.Bool{Bool: false, Valid: true}
	}

	query := `
		INSERT INTO users (login, email, password_hash, first_name, last_name, avatar_url, is_verified)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING ` + userColumns

	var created models.User
	err := r.db.QueryRowxContext(ctx, query,
		user.Login,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.AvatarUrl,
		isVerified,
	).StructScan(&created)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &created, nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id pgtype.UUID) (*models.User, error) {
	query := `SELECT ` + userColumns + ` FROM users WHERE id = $1`

	var user models.User
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT ` + userColumns + ` FROM users WHERE email = $1`

	var user models.User
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}

// GetByLogin retrieves a user by login
func (r *UserRepository) GetByLogin(ctx context.Context, login string) (*models.User, error) {
	query := `SELECT ` + userColumns + ` FROM users WHERE login = $1`

	var user models.User
	err := r.db.GetContext(ctx, &user, query, login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by login: %w", err)
	}
	return &user, nil
}

// Update modifies an existing user
func (r *UserRepository) Update(ctx context.Context, user models.User) (*models.User, error) {
	query := `
		UPDATE users SET
			email      = $1,
			password_hash = $2,
			first_name = $3,
			last_name  = $4,
			avatar_url = $5,
			is_verified = $6,
			updated_at = NOW()
		WHERE id = $7
		RETURNING ` + userColumns

	var updated models.User
	err := r.db.QueryRowxContext(ctx, query,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.AvatarUrl,
		user.IsVerified,
		user.ID,
	).StructScan(&updated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return &updated, nil
}

// Delete removes a user by ID
func (r *UserRepository) Delete(ctx context.Context, id pgtype.UUID) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

// ListSA retrieves users for the SA API with filtering, sorting, and pagination.
func (r *UserRepository) ListSA(ctx context.Context, f SAListFilter) ([]*models.User, int, error) {
	var conditions []string
	var args []any
	argIdx := 1

	if f.SearchLoginTerm != "" {
		conditions = append(conditions, fmt.Sprintf("login ILIKE $%d", argIdx))
		args = append(args, "%"+f.SearchLoginTerm+"%")
		argIdx++
	}
	if f.SearchEmailTerm != "" {
		conditions = append(conditions, fmt.Sprintf("email ILIKE $%d", argIdx))
		args = append(args, "%"+f.SearchEmailTerm+"%")
		argIdx++
	}

	where := ""
	if len(conditions) > 0 {
		where = "WHERE " + strings.Join(conditions, " AND ")
	}

	sortCol := saValidSortColumn(f.SortBy)
	sortDir := saValidSortDir(f.SortDirection)

	var total int
	if err := r.db.GetContext(ctx, &total, fmt.Sprintf(`SELECT COUNT(*) FROM users %s`, where), args...); err != nil {
		return nil, 0, fmt.Errorf("count SA users: %w", err)
	}

	pageSize := f.PageSize
	if pageSize < 1 || pageSize > 20 {
		pageSize = 10
	}
	pageNumber := f.PageNumber
	if pageNumber < 1 {
		pageNumber = 1
	}
	offset := (pageNumber - 1) * pageSize

	dataArgs := append(args, pageSize, offset)
	dataQuery := fmt.Sprintf(
		`SELECT `+saUserColumns+` FROM users %s ORDER BY %s %s LIMIT $%d OFFSET $%d`,
		where, sortCol, sortDir, argIdx, argIdx+1,
	)

	var users []*models.User
	if err := r.db.SelectContext(ctx, &users, dataQuery, dataArgs...); err != nil {
		return nil, 0, fmt.Errorf("list SA users: %w", err)
	}
	return users, total, nil
}

func saValidSortColumn(s string) string {
	switch s {
	case "login":
		return "login"
	case "email":
		return "email"
	default:
		return "created_at"
	}
}

func saValidSortDir(s string) string {
	if strings.ToLower(s) == "asc" {
		return "ASC"
	}
	return "DESC"
}
