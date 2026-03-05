package dto

// RefreshRequest represents the request body for token refresh (mobile clients)
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"` //nolint:gosec // API field name for token refresh request
}

// ChangeEmailRequest represents the request body for changing email
type ChangeEmailRequest struct {
	CurrentPassword string `json:"current_password" validate:"required,min=6"`
	NewEmail        string `json:"new_email" validate:"required,email"`
}

// ChangePasswordRequest represents the request body for changing password
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required,min=6"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}

