package dto

// AnswerRequest is the body for submitting an answer.
type AnswerRequest struct {
	Answer string `json:"answer" validate:"required"`
}
