package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/muzyk0/online-quiz-game/internal/app/database"
	"github.com/muzyk0/online-quiz-game/internal/domain/question/models"
)

var ErrQuestionNotFound = errors.New("question not found")

// ListFilter holds pagination and filtering options for listing questions.
type ListFilter struct {
	BodySearchTerm  string
	PublishedStatus string // "all" | "published" | "notPublished"
	SortBy          string // "createdAt" | "updatedAt" | "body"
	SortDirection   string // "asc" | "desc"
	PageNumber      int
	PageSize        int
}

//go:generate go run github.com/matryer/moq@latest -out ../service/mock_question_repository_test.go -pkg service . QuestionRepositoryInterface

// QuestionRepositoryInterface defines DB operations for quiz questions.
type QuestionRepositoryInterface interface {
	Create(ctx context.Context, q models.QuizQuestion) (*models.QuizQuestion, error)
	GetByID(ctx context.Context, id pgtype.UUID) (*models.QuizQuestion, error)
	Update(ctx context.Context, q models.QuizQuestion) (*models.QuizQuestion, error)
	Delete(ctx context.Context, id pgtype.UUID) error
	SetPublished(ctx context.Context, id pgtype.UUID, published bool) (*models.QuizQuestion, error)
	List(ctx context.Context, f ListFilter) ([]*models.QuizQuestion, int, error)
	ListPublished(ctx context.Context, limit int) ([]*models.QuizQuestion, error)
}

// QuestionRepository implements QuestionRepositoryInterface.
type QuestionRepository struct {
	db *database.DB
}

func NewQuestionRepository(db *database.DB) QuestionRepositoryInterface {
	return &QuestionRepository{db: db}
}

const columns = `id, body, correct_answers, published, created_at, updated_at`

func (r *QuestionRepository) Create(ctx context.Context, q models.QuizQuestion) (*models.QuizQuestion, error) {
	query := `
		INSERT INTO quiz_questions (body, correct_answers, published)
		VALUES (:body, :correct_answers, :published)
		RETURNING ` + columns

	rows, err := r.db.NamedQueryContext(ctx, query, q)
	if err != nil {
		return nil, fmt.Errorf("create question: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var out models.QuizQuestion
		if err := rows.StructScan(&out); err != nil {
			return nil, fmt.Errorf("scan created question: %w", err)
		}
		return &out, nil
	}
	return nil, errors.New("no row returned after insert")
}

func (r *QuestionRepository) GetByID(ctx context.Context, id pgtype.UUID) (*models.QuizQuestion, error) {
	query := `SELECT ` + columns + ` FROM quiz_questions WHERE id = $1`

	var q models.QuizQuestion
	if err := r.db.GetContext(ctx, &q, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrQuestionNotFound
		}
		return nil, fmt.Errorf("get question by id: %w", err)
	}
	return &q, nil
}

func (r *QuestionRepository) Update(ctx context.Context, q models.QuizQuestion) (*models.QuizQuestion, error) {
	query := `
		UPDATE quiz_questions
		SET body = :body, correct_answers = :correct_answers, updated_at = NOW()
		WHERE id = :id
		RETURNING ` + columns

	rows, err := r.db.NamedQueryContext(ctx, query, q)
	if err != nil {
		return nil, fmt.Errorf("update question: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var out models.QuizQuestion
		if err := rows.StructScan(&out); err != nil {
			return nil, fmt.Errorf("scan updated question: %w", err)
		}
		return &out, nil
	}
	return nil, ErrQuestionNotFound
}

func (r *QuestionRepository) Delete(ctx context.Context, id pgtype.UUID) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM quiz_questions WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete question: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return ErrQuestionNotFound
	}
	return nil
}

func (r *QuestionRepository) List(ctx context.Context, f ListFilter) ([]*models.QuizQuestion, int, error) {
	// Build WHERE clause
	var conditions []string
	var args []any
	argIdx := 1

	if f.BodySearchTerm != "" {
		conditions = append(conditions, fmt.Sprintf("body ILIKE $%d", argIdx))
		args = append(args, "%"+f.BodySearchTerm+"%")
		argIdx++
	}

	switch f.PublishedStatus {
	case "published":
		conditions = append(conditions, fmt.Sprintf("published = $%d", argIdx))
		args = append(args, true)
		argIdx++
	case "notPublished":
		conditions = append(conditions, fmt.Sprintf("published = $%d", argIdx))
		args = append(args, false)
		argIdx++
	}

	where := ""
	if len(conditions) > 0 {
		where = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Validate sort
	sortCol := validSortColumn(f.SortBy)
	sortDir := validSortDir(f.SortDirection)

	// Count total
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM quiz_questions %s`, where)
	var total int
	if err := r.db.GetContext(ctx, &total, countQuery, args...); err != nil {
		return nil, 0, fmt.Errorf("count questions: %w", err)
	}

	// Fetch page
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
		`SELECT %s FROM quiz_questions %s ORDER BY %s %s LIMIT $%d OFFSET $%d`,
		columns, where, sortCol, sortDir, argIdx, argIdx+1,
	)

	var questions []*models.QuizQuestion
	if err := r.db.SelectContext(ctx, &questions, dataQuery, dataArgs...); err != nil {
		return nil, 0, fmt.Errorf("list questions: %w", err)
	}
	return questions, total, nil
}

func (r *QuestionRepository) ListPublished(ctx context.Context, limit int) ([]*models.QuizQuestion, error) {
	query := fmt.Sprintf(
		`SELECT %s FROM quiz_questions WHERE published = true ORDER BY RANDOM() LIMIT $1`,
		columns,
	)
	var questions []*models.QuizQuestion
	if err := r.db.SelectContext(ctx, &questions, query, limit); err != nil {
		return nil, fmt.Errorf("list published questions: %w", err)
	}
	return questions, nil
}

// SetPublished updates the published flag for a question.
func (r *QuestionRepository) SetPublished(ctx context.Context, id pgtype.UUID, published bool) (*models.QuizQuestion, error) {
	query := `
		UPDATE quiz_questions
		SET published = $2, updated_at = NOW()
		WHERE id = $1
		RETURNING ` + columns

	var q models.QuizQuestion
	if err := r.db.GetContext(ctx, &q, query, id, published); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrQuestionNotFound
		}
		return nil, fmt.Errorf("set published: %w", err)
	}
	return &q, nil
}

func validSortColumn(s string) string {
	switch s {
	case "body":
		return "body"
	case "updatedAt":
		return "updated_at"
	default:
		return "created_at"
	}
}

func validSortDir(s string) string {
	if strings.ToLower(s) == "asc" {
		return "ASC"
	}
	return "DESC"
}
