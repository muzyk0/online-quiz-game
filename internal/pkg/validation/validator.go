package validation

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/muzyk0/online-quiz-game/internal/pkg/apperrors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var loginPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

// CustomValidator wraps the validator instance
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator creates a new validator instance
func NewValidator() *CustomValidator {
	v := validator.New()
	// Use JSON tag names in validation errors instead of Go struct field names.
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	_ = v.RegisterValidation("login_pattern", func(fl validator.FieldLevel) bool {
		return loginPattern.MatchString(fl.Field().String())
	})
	return &CustomValidator{validator: v}
}

// Validate validates a struct and returns an apperrors.AppError with field details
func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			details := formatValidationErrors(validationErrors)
			return apperrors.NewValidationError(details)
		}
		// Non-validation error (e.g., invalid struct)
		return apperrors.BadRequest(fmt.Sprintf("validation failed: %s", err.Error()))
	}
	return nil
}

// formatValidationErrors converts validator errors to a map of field -> message
func formatValidationErrors(errs validator.ValidationErrors) map[string]string {
	details := make(map[string]string, len(errs))
	for _, e := range errs {
		details[e.Field()] = formatFieldError(e)
	}
	return details
}

// formatFieldError formats a single field error
func formatFieldError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "must be a valid email address"
	case "min":
		return fmt.Sprintf("must be at least %s characters long", e.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters long", e.Param())
	case "url":
		return "must be a valid URL"
	case "uuid":
		return "must be a valid UUID"
	case "login_pattern":
		return "must contain only letters, numbers, underscore or hyphen"
	default:
		return fmt.Sprintf("failed validation on %s", e.Tag())
	}
}

// ValidationMiddleware creates a middleware that validates request bodies
func ValidationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Validation is handled by the custom validator via c.Validate()
			return next(c)
		}
	}
}
