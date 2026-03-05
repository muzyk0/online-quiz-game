package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	originalServerHost := os.Getenv("SERVER_HOST")
	originalServerPort := os.Getenv("SERVER_PORT")
	originalJWTExpiry := os.Getenv("JWT_ACCESS_EXPIRY_MIN")
	originalCors := os.Getenv("CORS_ALLOWED_ORIGINS")

	defer func() {
		_ = os.Setenv("SERVER_HOST", originalServerHost)
		_ = os.Setenv("SERVER_PORT", originalServerPort)
		_ = os.Setenv("JWT_ACCESS_EXPIRY_MIN", originalJWTExpiry)
		_ = os.Setenv("CORS_ALLOWED_ORIGINS", originalCors)
	}()

	_ = os.Unsetenv("SERVER_HOST")
	_ = os.Unsetenv("SERVER_PORT")
	_ = os.Unsetenv("JWT_ACCESS_EXPIRY_MIN")
	_ = os.Unsetenv("CORS_ALLOWED_ORIGINS")

	cfg := Load()
	assert.Equal(t, "localhost", cfg.ServerHost)
	assert.Equal(t, 8080, cfg.ServerPort)
	assert.Equal(t, 15, cfg.JWTAccessExpiryMin)
	assert.Equal(t, []string{"http://localhost:3000"}, cfg.CorsAllowedOrigins)

	_ = os.Setenv("SERVER_HOST", "production-host")
	_ = os.Setenv("SERVER_PORT", "9000")
	_ = os.Setenv("JWT_ACCESS_EXPIRY_MIN", "30")
	_ = os.Setenv("CORS_ALLOWED_ORIGINS", "https://example.com, https://api.example.com")

	cfg = Load()
	assert.Equal(t, "production-host", cfg.ServerHost)
	assert.Equal(t, 9000, cfg.ServerPort)
	assert.Equal(t, 30, cfg.JWTAccessExpiryMin)
	assert.Equal(t, []string{"https://example.com", "https://api.example.com"}, cfg.CorsAllowedOrigins)
}

func TestGetEnvOrDefault(t *testing.T) {
	originalValue := os.Getenv("TEST_VAR")
	defer func() { _ = os.Setenv("TEST_VAR", originalValue) }()

	_ = os.Unsetenv("TEST_VAR")
	result := getEnvOrDefault("TEST_VAR", "default_value")
	assert.Equal(t, "default_value", result)

	_ = os.Setenv("TEST_VAR", "custom_value")
	result = getEnvOrDefault("TEST_VAR", "default_value")
	assert.Equal(t, "custom_value", result)
}

func TestGetIntEnvOrDefault(t *testing.T) {
	originalValue := os.Getenv("TEST_INT_VAR")
	defer func() { _ = os.Setenv("TEST_INT_VAR", originalValue) }()

	_ = os.Unsetenv("TEST_INT_VAR")
	result := getIntEnvOrDefault("TEST_INT_VAR", 42)
	assert.Equal(t, 42, result)

	_ = os.Setenv("TEST_INT_VAR", "123")
	result = getIntEnvOrDefault("TEST_INT_VAR", 42)
	assert.Equal(t, 123, result)

	_ = os.Setenv("TEST_INT_VAR", "not_a_number")
	result = getIntEnvOrDefault("TEST_INT_VAR", 42)
	assert.Equal(t, 42, result)
}

func TestGetSliceEnvOrDefault(t *testing.T) {
	originalValue := os.Getenv("TEST_SLICE_VAR")
	defer func() { _ = os.Setenv("TEST_SLICE_VAR", originalValue) }()

	_ = os.Unsetenv("TEST_SLICE_VAR")
	result := getSliceEnvOrDefault("TEST_SLICE_VAR", []string{"default1", "default2"})
	assert.Equal(t, []string{"default1", "default2"}, result)

	_ = os.Setenv("TEST_SLICE_VAR", "value1, value2 , value3")
	result = getSliceEnvOrDefault("TEST_SLICE_VAR", []string{"default1", "default2"})
	assert.Equal(t, []string{"value1", "value2", "value3"}, result)
}
