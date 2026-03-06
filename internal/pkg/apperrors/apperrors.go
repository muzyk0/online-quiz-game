// Package apperrors provides a unified error type for the application.
//
// AppError carries an HTTP status code and a safe client message.
// Handlers return AppError values; the centralized error handler
// (middleware.CustomHTTPErrorHandler) converts them to JSON responses.
//
// Usage in handlers:
//
//	return apperrors.NotFound("Game not found")
//	return apperrors.Internal("Failed to generate token").Wrap(err)
//
// Usage for validation:
//
//	return apperrors.NewValidationError(map[string]string{
//	    "email": "must be a valid email address",
//	})
package apperrors

import (
	"fmt"
	"net/http"
)

// FieldError describes a single field validation error matching the spec format.
type FieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

// AppError is the single application error type.
// It implements the error interface and carries HTTP semantics.
type AppError struct {
	// Code is the HTTP status code.
	Code int `json:"-"`
	// Message is the safe message sent to the client.
	Message string `json:"error"`
	// Details contains field-level validation errors (optional).
	Details map[string]string `json:"details,omitempty"`
	// Err is the underlying cause (logged server-side, never sent to client).
	Err error `json:"-"`
}

// ErrorsMessages returns the Details as a slice of FieldError for spec-compliant output.
func (e *AppError) ErrorsMessages() []FieldError {
	if len(e.Details) == 0 {
		return []FieldError{{Message: e.Message, Field: ""}}
	}
	out := make([]FieldError, 0, len(e.Details))
	for field, msg := range e.Details {
		out = append(out, FieldError{Message: msg, Field: field})
	}
	return out
}

// Error implements the error interface.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error for errors.Is/As support.
func (e *AppError) Unwrap() error {
	return e.Err
}

// Wrap attaches an underlying cause and returns a new AppError.
// The original is not mutated.
func (e *AppError) Wrap(err error) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: e.Message,
		Details: copyDetails(e.Details),
		Err:     err,
	}
}

// WithMessage returns a copy with a different client message.
func (e *AppError) WithMessage(msg string) *AppError {
	return &AppError{
		Code:    e.Code,
		Message: msg,
		Details: copyDetails(e.Details),
		Err:     e.Err,
	}
}

// --- Constructors ---

// New creates an AppError with the given status code and message.
func New(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// BadRequest creates a 400 error.
func BadRequest(message string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: message}
}

// Unauthorized creates a 401 error.
func Unauthorized(message string) *AppError {
	return &AppError{Code: http.StatusUnauthorized, Message: message}
}

// Forbidden creates a 403 error.
func Forbidden(message string) *AppError {
	return &AppError{Code: http.StatusForbidden, Message: message}
}

// NotFound creates a 404 error.
func NotFound(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

// Conflict creates a 409 error.
func Conflict(message string) *AppError {
	return &AppError{Code: http.StatusConflict, Message: message}
}

// TooManyRequests creates a 429 error.
func TooManyRequests(message string) *AppError {
	return &AppError{Code: http.StatusTooManyRequests, Message: message}
}

// Internal creates a 500 error.
func Internal(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}

// BadGateway creates a 502 error.
func BadGateway(message string) *AppError {
	return &AppError{Code: http.StatusBadGateway, Message: message}
}

// NewValidationError creates a 400 error with field-level details.
func NewValidationError(details map[string]string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "Validation failed",
		Details: copyDetails(details),
	}
}

// copyDetails returns a shallow copy of a Details map to prevent shared mutable state.
func copyDetails(d map[string]string) map[string]string {
	if d == nil {
		return nil
	}
	out := make(map[string]string, len(d))
	for k, v := range d {
		out[k] = v
	}
	return out
}
