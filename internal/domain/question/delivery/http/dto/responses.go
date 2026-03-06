package dto

import (
	"time"

	"github.com/muzyk0/online-quiz-game/internal/domain/question/service"
)

// QuestionResponse is the API representation of a quiz question.
type QuestionResponse struct {
	ID             string    `json:"id"`
	Body           string    `json:"body"`
	CorrectAnswers []string  `json:"correctAnswers"`
	Published      bool      `json:"published"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

// PaginatedQuestionsResponse is the paginated list response.
type PaginatedQuestionsResponse struct {
	PagesCount int                 `json:"pagesCount"`
	Page       int                 `json:"page"`
	PageSize   int                 `json:"pageSize"`
	TotalCount int                 `json:"totalCount"`
	Items      []*QuestionResponse `json:"items"`
}

// FromServiceOutput converts a service QuestionOutput to QuestionResponse.
func FromServiceOutput(q *service.QuestionOutput) *QuestionResponse {
	r := &QuestionResponse{
		ID:             q.ID,
		Body:           q.Body,
		CorrectAnswers: q.CorrectAnswers,
		Published:      q.Published,
	}
	if q.CorrectAnswers == nil {
		r.CorrectAnswers = []string{}
	}
	if q.CreatedAt.Valid {
		r.CreatedAt = q.CreatedAt.Time
	}
	if q.UpdatedAt.Valid {
		t := q.UpdatedAt.Time
		r.UpdatedAt = &t
	}
	return r
}

// FromPaginatedOutput converts a service paginated output to response.
func FromPaginatedOutput(p *service.PaginatedQuestionsOutput) *PaginatedQuestionsResponse {
	items := make([]*QuestionResponse, len(p.Items))
	for i, q := range p.Items {
		items[i] = FromServiceOutput(q)
	}
	return &PaginatedQuestionsResponse{
		PagesCount: p.PagesCount,
		Page:       p.Page,
		PageSize:   p.PageSize,
		TotalCount: p.TotalCount,
		Items:      items,
	}
}

// ErrorResponse is a generic error response.
type ErrorResponse struct {
	Error string `json:"error"`
}
