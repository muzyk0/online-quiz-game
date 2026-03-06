package dto

// CreateQuestionRequest is the SA request body to create a question.
type CreateQuestionRequest struct {
	Body           string   `json:"body" validate:"required,min=10,max=500"`
	CorrectAnswers []string `json:"correctAnswers" validate:"required,min=1,dive,required"`
}

// UpdateQuestionRequest is the SA request body to update a question.
type UpdateQuestionRequest struct {
	Body           string   `json:"body" validate:"required,min=10,max=500"`
	CorrectAnswers []string `json:"correctAnswers" validate:"required,min=1,dive,required"`
}

// PublishQuestionRequest is the SA request body to publish/unpublish a question.
type PublishQuestionRequest struct {
	Published *bool `json:"published" validate:"required"`
}
