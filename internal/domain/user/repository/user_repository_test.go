package repository

import (
	"testing"
)

// Integration tests for UserRepository require a real database.
// Unit-level encryption tests were removed as PII encryption was removed from the quiz service.
// Run integration tests with: go test -tags=integration ./...

func TestUserRepository_Placeholder(t *testing.T) {
	t.Skip("integration tests require a database connection")
}
