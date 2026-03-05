package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/muzyk0/online-quiz-game/internal/domain/question/models"
	"github.com/muzyk0/online-quiz-game/internal/domain/question/repository"
)

// Sentinel errors
var (
	ErrQuestionNotFound        = errors.New("question not found")
	ErrBodyRequired            = errors.New("body is required")
	ErrCorrectAnswersRequired  = errors.New("at least one correct answer is required to publish")
)

// QuestionServiceInterface defines business operations for quiz questions.
type QuestionServiceInterface interface {
	CreateQuestion(ctx context.Context, input CreateQuestionInput) (*QuestionOutput, error)
	UpdateQuestion(ctx context.Context, id string, input UpdateQuestionInput) (*QuestionOutput, error)
	DeleteQuestion(ctx context.Context, id string) error
	PublishQuestion(ctx context.Context, id string, published bool) (*QuestionOutput, error)
	ListQuestions(ctx context.Context, input ListQuestionsInput) (*PaginatedQuestionsOutput, error)
}

// --- Input / Output types ---

type CreateQuestionInput struct {
	Body           string
	CorrectAnswers []string
}

type UpdateQuestionInput struct {
	Body           string
	CorrectAnswers []string
}

type ListQuestionsInput struct {
	BodySearchTerm  string
	PublishedStatus string
	SortBy          string
	SortDirection   string
	PageNumber      int
	PageSize        int
}

type QuestionOutput struct {
	ID             string
	Body           string
	CorrectAnswers []string
	Published      bool
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
}

type PaginatedQuestionsOutput struct {
	PagesCount int
	Page       int
	PageSize   int
	TotalCount int
	Items      []*QuestionOutput
}

// QuestionService implements QuestionServiceInterface.
type QuestionService struct {
	repo repository.QuestionRepositoryInterface
}

func NewQuestionService(repo repository.QuestionRepositoryInterface) QuestionServiceInterface {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) CreateQuestion(ctx context.Context, input CreateQuestionInput) (*QuestionOutput, error) {
	if input.Body == "" {
		return nil, ErrBodyRequired
	}

	q := models.QuizQuestion{
		Body:           input.Body,
		CorrectAnswers: input.CorrectAnswers,
		Published:      false,
	}

	created, err := s.repo.Create(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("create question: %w", err)
	}
	return toOutput(created), nil
}

func (s *QuestionService) UpdateQuestion(ctx context.Context, id string, input UpdateQuestionInput) (*QuestionOutput, error) {
	uid, err := parseUUID(id)
	if err != nil {
		return nil, ErrQuestionNotFound
	}

	existing, err := s.repo.GetByID(ctx, uid)
	if err != nil {
		if errors.Is(err, repository.ErrQuestionNotFound) {
			return nil, ErrQuestionNotFound
		}
		return nil, fmt.Errorf("get question: %w", err)
	}

	existing.Body = input.Body
	existing.CorrectAnswers = input.CorrectAnswers

	updated, err := s.repo.Update(ctx, *existing)
	if err != nil {
		if errors.Is(err, repository.ErrQuestionNotFound) {
			return nil, ErrQuestionNotFound
		}
		return nil, fmt.Errorf("update question: %w", err)
	}
	return toOutput(updated), nil
}

func (s *QuestionService) DeleteQuestion(ctx context.Context, id string) error {
	uid, err := parseUUID(id)
	if err != nil {
		return ErrQuestionNotFound
	}

	if err := s.repo.Delete(ctx, uid); err != nil {
		if errors.Is(err, repository.ErrQuestionNotFound) {
			return ErrQuestionNotFound
		}
		return fmt.Errorf("delete question: %w", err)
	}
	return nil
}

func (s *QuestionService) PublishQuestion(ctx context.Context, id string, published bool) (*QuestionOutput, error) {
	uid, err := parseUUID(id)
	if err != nil {
		return nil, ErrQuestionNotFound
	}

	// Validate: can only publish if there are correct answers
	if published {
		existing, err := s.repo.GetByID(ctx, uid)
		if err != nil {
			if errors.Is(err, repository.ErrQuestionNotFound) {
				return nil, ErrQuestionNotFound
			}
			return nil, fmt.Errorf("get question: %w", err)
		}
		if len(existing.CorrectAnswers) == 0 {
			return nil, ErrCorrectAnswersRequired
		}
	}

	updated, err := s.repo.SetPublished(ctx, uid, published)
	if err != nil {
		if errors.Is(err, repository.ErrQuestionNotFound) {
			return nil, ErrQuestionNotFound
		}
		return nil, fmt.Errorf("set published: %w", err)
	}
	return toOutput(updated), nil
}

func (s *QuestionService) ListQuestions(ctx context.Context, input ListQuestionsInput) (*PaginatedQuestionsOutput, error) {
	f := repository.ListFilter{
		BodySearchTerm:  input.BodySearchTerm,
		PublishedStatus: input.PublishedStatus,
		SortBy:          input.SortBy,
		SortDirection:   input.SortDirection,
		PageNumber:      input.PageNumber,
		PageSize:        input.PageSize,
	}

	items, total, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("list questions: %w", err)
	}

	pageSize := input.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	pagesCount := 0
	if total > 0 {
		pagesCount = (total + pageSize - 1) / pageSize
	}

	out := make([]*QuestionOutput, len(items))
	for i, q := range items {
		out[i] = toOutput(q)
	}

	return &PaginatedQuestionsOutput{
		PagesCount: pagesCount,
		Page:       input.PageNumber,
		PageSize:   pageSize,
		TotalCount: total,
		Items:      out,
	}, nil
}

// --- helpers ---

func toOutput(q *models.QuizQuestion) *QuestionOutput {
	id, _ := q.ID.Value()
	idStr, _ := id.(string)
	return &QuestionOutput{
		ID:             idStr,
		Body:           q.Body,
		CorrectAnswers: q.CorrectAnswers,
		Published:      q.Published,
		CreatedAt:      q.CreatedAt,
		UpdatedAt:      q.UpdatedAt,
	}
}

func parseUUID(s string) (pgtype.UUID, error) {
	var uid pgtype.UUID
	if err := uid.Scan(s); err != nil {
		return uid, err
	}
	return uid, nil
}
