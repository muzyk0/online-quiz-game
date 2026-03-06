package config

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"strconv"
	"strings"
)

// Config holds the application configuration
type Config struct {
	ServerHost string
	ServerPort int
	ServerEnv  string

	DatabaseURL          string
	DBMaxOpenConns       int
	DBMaxIdleConns       int
	DBConnMaxLifetimeMin int

	JWTSecret            string //nolint:gosec
	JWTAccessExpiryMin   int
	JWTRefreshExpiryDays int

	CorsAllowedOrigins []string

	// Super Admin credentials for SA API (Basic auth)
	SAAdminLogin    string
	SAAdminPassword string //nolint:gosec
}

// Load loads the configuration from environment variables
func Load() *Config {
	serverEnv := getEnvOrDefault("SERVER_ENV", "development")
	jwtSecret := getEnvOrDefault("JWT_SECRET", "")

	if serverEnv != "development" && jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set in production environments")
	}

	if jwtSecret == "" {
		b := make([]byte, 32)
		if _, err := rand.Read(b); err != nil {
			log.Fatal("Failed to generate random JWT secret:", err)
		}
		jwtSecret = base64.StdEncoding.EncodeToString(b)
		log.Println("WARNING: Using generated temporary JWT secret for development. Set JWT_SECRET for persistence.")
	}

	return &Config{
		ServerHost: getEnvOrDefault("SERVER_HOST", "localhost"),
		ServerPort: getIntEnvOrDefault("SERVER_PORT", 8080),
		ServerEnv:  serverEnv,

		DatabaseURL:          getEnvOrDefault("DATABASE_URL", "postgres://user:password@localhost:5432/quiz_db?sslmode=disable"),
		DBMaxOpenConns:       getIntEnvOrDefault("DB_MAX_OPEN_CONNS", 25),
		DBMaxIdleConns:       getIntEnvOrDefault("DB_MAX_IDLE_CONNS", 5),
		DBConnMaxLifetimeMin: getIntEnvOrDefault("DB_CONN_MAX_LIFETIME_MIN", 5),

		JWTSecret:            jwtSecret,
		JWTAccessExpiryMin:   getIntEnvOrDefault("JWT_ACCESS_EXPIRY_MIN", 15),
		JWTRefreshExpiryDays: getIntEnvOrDefault("JWT_REFRESH_EXPIRY_DAYS", 7),

		CorsAllowedOrigins: getSliceEnvOrDefault("CORS_ALLOWED_ORIGINS", []string{
			"http://localhost:3000",
		}),
		SAAdminLogin:    getEnvOrDefault("SA_ADMIN_LOGIN", "admin"),
		SAAdminPassword: getEnvOrDefault("SA_ADMIN_PASSWORD", "admin"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getIntEnvOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getSliceEnvOrDefault(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		slice := strings.Split(value, ",")
		for i, v := range slice {
			slice[i] = strings.TrimSpace(v)
		}
		return slice
	}
	return defaultValue
}
