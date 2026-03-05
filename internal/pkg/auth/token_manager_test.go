package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestTokenManager() *TokenManager {
	return NewTokenManager("test-secret", 15*time.Minute, 7*24*time.Hour)
}

func TestNewTokenManager(t *testing.T) {
	secret := "test-secret"
	tm := NewTokenManager(secret, 15*time.Minute, 7*24*time.Hour)

	assert.Equal(t, []byte(secret), tm.secret)
	assert.Equal(t, 15*time.Minute, tm.accessExpiry)
	assert.Equal(t, 7*24*time.Hour, tm.refreshExpiry)
}

func TestValidateToken(t *testing.T) {
	tm := newTestTokenManager()

	userID := "user-123"
	email := "test@example.com"
	userType := "user"

	tokenString, err := tm.GenerateAccessToken(userID, email, userType)
	require.NoError(t, err)

	claims, err := tm.ValidateToken(tokenString)
	require.NoError(t, err)

	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)
	assert.Equal(t, userType, claims.UserType)
	assert.Equal(t, "online-quiz-game", claims.Issuer)
}

func TestValidateTokenInvalid(t *testing.T) {
	tm := newTestTokenManager()

	// Test with invalid token
	_, err := tm.ValidateToken("invalid-token")
	require.Error(t, err)

	// Test with wrong secret
	wrongTm := NewTokenManager("wrong-secret", 15*time.Minute, 7*24*time.Hour)
	tokenString, err := tm.GenerateAccessToken("user-123", "test@example.com", "user")
	require.NoError(t, err)

	_, err = wrongTm.ValidateToken(tokenString)
	assert.Error(t, err)
}

func TestGenerateAccessToken(t *testing.T) {
	tm := newTestTokenManager()

	userID := "user-123"
	email := "test@example.com"
	userType := "user"

	tokenString, err := tm.GenerateAccessToken(userID, email, userType)
	require.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte("test-secret"), nil
	})
	require.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(*Claims)
	require.True(t, ok)

	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)
	assert.Equal(t, userType, claims.UserType)
	assert.Equal(t, "online-quiz-game", claims.Issuer)
	assert.Empty(t, claims.TokenID, "Access tokens should not have TokenID")
}

func TestGenerateRefreshToken(t *testing.T) {
	tm := newTestTokenManager()

	userID := "user-123"
	email := "test@example.com"
	userType := "user"
	tokenID := "token-id-123"

	tokenString, err := tm.GenerateRefreshToken(userID, email, userType, tokenID)
	require.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte("test-secret"), nil
	})
	require.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(*Claims)
	require.True(t, ok)

	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, email, claims.Email)
	assert.Equal(t, userType, claims.UserType)
	assert.Equal(t, tokenID, claims.TokenID)
	assert.Equal(t, "online-quiz-game", claims.Issuer)
}
