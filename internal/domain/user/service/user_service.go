package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/muzyk0/online-quiz-game/internal/domain/user/models"
	"github.com/muzyk0/online-quiz-game/internal/domain/user/repository"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

// Sentinel errors
var (
	ErrUserAlreadyExists   = errors.New("user with this email already exists")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrCredentialsRequired = errors.New("email and password are required")
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrInvalidUserID       = errors.New("invalid user id")
)

// SACreateUserInput contains data for admin-created users.
type SACreateUserInput struct {
	Login    string
	Password string //nolint:gosec
	Email    string
}

// SAListUsersInput holds query params for the SA users list.
type SAListUsersInput struct {
	SearchLoginTerm string
	SearchEmailTerm string
	SortBy          string
	SortDirection   string
	PageNumber      int
	PageSize        int
}

// PaginatedUsersOutput is the paginated list of users for SA.
type PaginatedUsersOutput struct {
	PagesCount int
	Page       int
	PageSize   int
	TotalCount int
	Items      []*UserOutput
}

// UserServiceInterface defines the interface for user-related operations
type UserServiceInterface interface {
	Register(ctx context.Context, input RegisterUserInput) (*UserOutput, error)
	Login(ctx context.Context, input LoginUserInput) (*UserOutput, error)
	LoginByLoginOrEmail(ctx context.Context, loginOrEmail, password string) (*UserOutput, error)
	GetUser(ctx context.Context, userID string) (*UserOutput, error)
	UpdateProfile(ctx context.Context, userID string, input UpdateProfileInput) (*UserOutput, error)
	ChangeEmail(ctx context.Context, userID, currentPassword, newEmail string) error
	ChangePassword(ctx context.Context, userID, currentPassword, newPassword string) error
	DeleteUser(ctx context.Context, userID string) error

	// SA operations
	SACreateUser(ctx context.Context, input SACreateUserInput) (*UserOutput, error)
	SADeleteUser(ctx context.Context, userID string) error
	SAListUsers(ctx context.Context, input SAListUsersInput) (*PaginatedUsersOutput, error)
}

// UserService implements business logic for user operations.
type UserService struct {
	repo repository.UserRepositoryInterface
}

// NewUserService creates a new UserService instance.
func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

// RegisterUserInput contains the data required to register a new user.
type RegisterUserInput struct {
	Email     string
	Password  string //nolint:gosec // Service input field name for user registration
	FirstName string
	LastName  string
	AvatarUrl string
}

// LoginUserInput contains the data required for user login.
type LoginUserInput struct {
	Email    string
	Password string //nolint:gosec // Service input field name for user login
}

// UpdateUserInput contains fields for updating a user (all optional).
type UpdateUserInput struct {
	Email     *string
	Password  *string //nolint:gosec // Service input field name for user update
	FirstName *string
	LastName  *string
	AvatarUrl *string
}

// UpdateProfileInput contains fields for updating user profile information.
type UpdateProfileInput struct {
	FirstName *string
	LastName  *string
	AvatarUrl *string
}

// UserOutput represents the user data returned by service operations.
type UserOutput struct {
	ID        string
	Login     string
	Email     string
	FirstName string
	LastName  string
	AvatarUrl string
	CreatedAt pgtype.Timestamptz
}

// Register creates a new user account with the provided registration data.
func (s *UserService) Register(ctx context.Context, input RegisterUserInput) (*UserOutput, error) {
	// Validate input
	if input.Email == "" || input.Password == "" {
		return nil, ErrCredentialsRequired
	}

	// Check if user already exists
	existingUser, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		// If error is "user not found", continue with registration
		if !errors.Is(err, repository.ErrUserNotFound) {
			// Surface other database errors
			return nil, fmt.Errorf("failed to check existing user: %w", err)
		}
	}
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := models.User{
		Email: input.Email,
		PasswordHash: pgtype.Text{
			String: string(hashedPassword),
			Valid:  true,
		},
		FirstName: pgtype.Text{
			String: input.FirstName,
			Valid:  input.FirstName != "",
		},
		LastName: pgtype.Text{
			String: input.LastName,
			Valid:  input.LastName != "",
		},
		AvatarUrl: pgtype.Text{
			String: input.AvatarUrl,
			Valid:  input.AvatarUrl != "",
		},
		IsVerified: pgtype.Bool{
			Bool:  false,
			Valid: true,
		},
	}

	createdUser, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return toUserOutput(createdUser), nil
}

// Login authenticates a user with email and password.
func (s *UserService) Login(ctx context.Context, input LoginUserInput) (*UserOutput, error) {
	// Validate input
	if input.Email == "" || input.Password == "" {
		return nil, ErrCredentialsRequired
	}

	// Get user by email
	user, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Check if password hash is valid
	if !user.PasswordHash.Valid {
		return nil, ErrInvalidCredentials
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(input.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return toUserOutput(user), nil
}

// LoginByLoginOrEmail authenticates a user by either login or email with password.
func (s *UserService) LoginByLoginOrEmail(ctx context.Context, loginOrEmail, password string) (*UserOutput, error) {
	if loginOrEmail == "" || password == "" {
		return nil, ErrCredentialsRequired
	}

	// Try by email first
	user, err := s.repo.GetByEmail(ctx, loginOrEmail)
	if err != nil {
		if !errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		// Try by login
		user, err = s.repo.GetByLogin(ctx, loginOrEmail)
		if err != nil {
			return nil, ErrInvalidCredentials
		}
	}

	if !user.PasswordHash.Valid {
		return nil, ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}
	return toUserOutput(user), nil
}

// GetUser retrieves a user by their ID.
func (s *UserService) GetUser(ctx context.Context, userID string) (*UserOutput, error) {
	id := pgtype.UUID{}
	if err := id.Scan(userID); err != nil {
		return nil, ErrInvalidUserID
	}

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return toUserOutput(user), nil
}

// DeleteUser permanently removes a user account.
func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	id := pgtype.UUID{}
	if err := id.Scan(userID); err != nil {
		return ErrInvalidUserID
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// UpdateProfile updates only non-sensitive profile information (firstName, lastName, avatarUrl)
func (s *UserService) UpdateProfile(ctx context.Context, userID string, input UpdateProfileInput) (*UserOutput, error) {
	id := pgtype.UUID{}
	if err := id.Scan(userID); err != nil {
		return nil, ErrInvalidUserID
	}

	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Update only profile fields (no email or password)
	if input.FirstName != nil {
		user.FirstName = pgtype.Text{
			String: *input.FirstName,
			Valid:  true,
		}
	}
	if input.LastName != nil {
		user.LastName = pgtype.Text{
			String: *input.LastName,
			Valid:  true,
		}
	}
	if input.AvatarUrl != nil {
		user.AvatarUrl = pgtype.Text{
			String: *input.AvatarUrl,
			Valid:  true,
		}
	}

	updatedUser, err := s.repo.Update(ctx, *user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return toUserOutput(updatedUser), nil
}

// ChangeEmail changes the user's email address with password verification
func (s *UserService) ChangeEmail(ctx context.Context, userID, currentPassword, newEmail string) error {
	id := pgtype.UUID{}
	if err := id.Scan(userID); err != nil {
		return ErrInvalidUserID
	}

	// Get current user
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Verify current password
	if !user.PasswordHash.Valid {
		return ErrInvalidPassword
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(currentPassword)); err != nil {
		return ErrInvalidPassword
	}

	// Check if new email is already in use by another account
	existingUser, err := s.repo.GetByEmail(ctx, newEmail)
	if err == nil && existingUser.ID != user.ID {
		return ErrUserAlreadyExists
	}

	// Update email
	user.Email = newEmail

	_, err = s.repo.Update(ctx, *user)
	if err != nil {
		return fmt.Errorf("failed to update user email: %w", err)
	}

	return nil
}

// ChangePassword changes the user's password with current password verification
func (s *UserService) ChangePassword(ctx context.Context, userID, currentPassword, newPassword string) error {
	id := pgtype.UUID{}
	if err := id.Scan(userID); err != nil {
		return ErrInvalidUserID
	}

	// Get current user
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Verify current password
	if !user.PasswordHash.Valid {
		return ErrInvalidPassword
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(currentPassword)); err != nil {
		return ErrInvalidPassword
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Update password
	user.PasswordHash = pgtype.Text{
		String: string(hashedPassword),
		Valid:  true,
	}

	_, err = s.repo.Update(ctx, *user)
	if err != nil {
		return fmt.Errorf("failed to update user password: %w", err)
	}

	return nil
}

// SACreateUser creates a new user from the SA panel (bypasses email verification).
var ErrLoginAlreadyExists = errors.New("user with this login already exists")

func (s *UserService) SACreateUser(ctx context.Context, input SACreateUserInput) (*UserOutput, error) {
	// Check email uniqueness
	if existing, err := s.repo.GetByEmail(ctx, input.Email); err == nil && existing != nil {
		return nil, ErrUserAlreadyExists
	}

	// Check login uniqueness
	if existing, err := s.repo.GetByLogin(ctx, input.Login); err == nil && existing != nil {
		return nil, ErrLoginAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user := models.User{
		Login: pgtype.Text{String: input.Login, Valid: true},
		Email: input.Email,
		PasswordHash: pgtype.Text{
			String: string(hashedPassword),
			Valid:  true,
		},
		IsVerified: pgtype.Bool{Bool: true, Valid: true},
	}

	created, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return toUserOutput(created), nil
}

// SADeleteUser deletes a user by ID (SA operation).
func (s *UserService) SADeleteUser(ctx context.Context, userID string) error {
	id := pgtype.UUID{}
	if err := id.Scan(userID); err != nil {
		return ErrUserNotFound
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserNotFound
		}
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}

// SAListUsers lists users for the SA panel with filtering and pagination.
func (s *UserService) SAListUsers(ctx context.Context, input SAListUsersInput) (*PaginatedUsersOutput, error) {
	f := repository.SAListFilter{
		SearchLoginTerm: input.SearchLoginTerm,
		SearchEmailTerm: input.SearchEmailTerm,
		SortBy:          input.SortBy,
		SortDirection:   input.SortDirection,
		PageNumber:      input.PageNumber,
		PageSize:        input.PageSize,
	}

	items, total, err := s.repo.ListSA(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("list SA users: %w", err)
	}

	pageSize := input.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	pagesCount := 0
	if total > 0 {
		pagesCount = (total + pageSize - 1) / pageSize
	}

	out := make([]*UserOutput, len(items))
	for i, u := range items {
		out[i] = toUserOutput(u)
	}

	return &PaginatedUsersOutput{
		PagesCount: pagesCount,
		Page:       input.PageNumber,
		PageSize:   pageSize,
		TotalCount: total,
		Items:      out,
	}, nil
}

// toUserOutput converts a user model to UserOutput.
func toUserOutput(u *models.User) *UserOutput {
	return &UserOutput{
		ID:        u.ID.String(),
		Login:     u.Login.String,
		Email:     u.Email,
		FirstName: u.FirstName.String,
		LastName:  u.LastName.String,
		AvatarUrl: u.AvatarUrl.String,
		CreatedAt: u.CreatedAt,
	}
}
