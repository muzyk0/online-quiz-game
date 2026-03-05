package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenTypeAccess  = "access"
	tokenTypeRefresh = "refresh"
)

// Claims represents the JWT claims
type Claims struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	UserType  string `json:"user_type"`
	TokenType string `json:"token_type"`
	TokenID   string `json:"token_id,omitempty"` // For refresh tokens only (enables rotation/blacklisting)
	jwt.RegisteredClaims
}

// TokenManager handles JWT token operations
type TokenManager struct {
	secret        []byte
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

// NewTokenManager creates a new TokenManager
func NewTokenManager(secret string, accessExpiry, refreshExpiry time.Duration) *TokenManager {
	return &TokenManager{
		secret:        []byte(secret),
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (tm *TokenManager) parseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return tm.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ValidateToken validates an access token and returns the claims.
func (tm *TokenManager) ValidateToken(tokenString string) (*Claims, error) {
	return tm.validateTokenType(tokenString, tokenTypeAccess)
}

// ValidateRefreshToken validates a refresh token and returns the claims.
func (tm *TokenManager) ValidateRefreshToken(tokenString string) (*Claims, error) {
	return tm.validateTokenType(tokenString, tokenTypeRefresh)
}

func (tm *TokenManager) validateTokenType(tokenString, expectedType string) (*Claims, error) {
	claims, err := tm.parseToken(tokenString)
	if err != nil {
		return nil, err
	}
	if claims.TokenType != expectedType {
		return nil, fmt.Errorf("unexpected token type: %s", claims.TokenType)
	}
	return claims, nil
}

// GenerateAccessToken generates a short-lived access token for API authentication.
func (tm *TokenManager) GenerateAccessToken(userID, email, userType string) (string, error) {
	claims := Claims{
		UserID:    userID,
		Email:     email,
		UserType:  userType,
		TokenType: tokenTypeAccess,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.accessExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "online-quiz-game",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(tm.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign access token: %w", err)
	}
	return signedToken, nil
}

// GenerateRefreshToken generates a long-lived refresh token with a unique token ID for rotation support.
func (tm *TokenManager) GenerateRefreshToken(userID, email, userType, tokenID string) (string, error) {
	claims := Claims{
		UserID:    userID,
		Email:     email,
		UserType:  userType,
		TokenType: tokenTypeRefresh,
		TokenID:   tokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "online-quiz-game",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(tm.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign refresh token: %w", err)
	}
	return signedToken, nil
}
