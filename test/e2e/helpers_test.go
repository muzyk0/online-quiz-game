//go:build e2e

package e2e

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/muzyk0/online-quiz-game/internal/app/config"
	"github.com/muzyk0/online-quiz-game/internal/app/database"
	appmiddleware "github.com/muzyk0/online-quiz-game/internal/app/middleware"
	authhttp "github.com/muzyk0/online-quiz-game/internal/domain/auth/delivery/http"
	gamehttp "github.com/muzyk0/online-quiz-game/internal/domain/game/delivery/http"
	gamerepo "github.com/muzyk0/online-quiz-game/internal/domain/game/repository"
	gameservice "github.com/muzyk0/online-quiz-game/internal/domain/game/service"
	healthhttp "github.com/muzyk0/online-quiz-game/internal/domain/health/delivery/http"
	questionhttp "github.com/muzyk0/online-quiz-game/internal/domain/question/delivery/http"
	questionrepo "github.com/muzyk0/online-quiz-game/internal/domain/question/repository"
	questionservice "github.com/muzyk0/online-quiz-game/internal/domain/question/service"
	testinghttp "github.com/muzyk0/online-quiz-game/internal/domain/testing/delivery/http"
	userhttp "github.com/muzyk0/online-quiz-game/internal/domain/user/delivery/http"
	userrepo "github.com/muzyk0/online-quiz-game/internal/domain/user/repository"
	userservice "github.com/muzyk0/online-quiz-game/internal/domain/user/service"
	"github.com/muzyk0/online-quiz-game/internal/pkg/auth"
	"github.com/muzyk0/online-quiz-game/internal/pkg/validation"
	"github.com/stretchr/testify/require"
)

// testServer holds the test server and HTTP client for e2e tests.
type testServer struct {
	server *httptest.Server
	client *http.Client
	saUser string // "login:password" for basic auth
}

// newTestServer starts a real echo app backed by a real database.
// Skips if DATABASE_URL is not set or DB is unreachable.
func newTestServer(t *testing.T) *testServer {
	t.Helper()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/quiz_db?sslmode=disable"
	}

	cfg := &config.Config{
		ServerEnv:            "test",
		DatabaseURL:          dbURL,
		DBMaxOpenConns:       5,
		DBMaxIdleConns:       2,
		DBConnMaxLifetimeMin: 5,
		JWTSecret:            "test-secret-for-e2e-tests-only",
		JWTAccessExpiryMin:   60,
		JWTRefreshExpiryDays: 7,
		CorsAllowedOrigins:   []string{"*"},
		SAAdminLogin:         "admin",
		SAAdminPassword:      "qwerty",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, cfg.DatabaseURL, cfg.DBMaxOpenConns, cfg.DBMaxIdleConns, cfg.DBConnMaxLifetimeMin)
	if err != nil {
		t.Skipf("skipping e2e test: cannot connect to database: %v", err)
	}

	tokenManager := auth.NewTokenManager(cfg.JWTSecret,
		time.Duration(cfg.JWTAccessExpiryMin)*time.Minute,
		time.Duration(cfg.JWTRefreshExpiryDays)*24*time.Hour,
	)

	userRepo := userrepo.NewUserRepository(db)
	userSvc := userservice.NewUserService(userRepo)

	healthHandler := healthhttp.NewHandler(db)
	userHandler := userhttp.NewHandler(userSvc, tokenManager)
	saUserHandler := userhttp.NewSAHandler(userSvc)
	authHandler := authhttp.NewHandler(userSvc, tokenManager)

	questionRepository := questionrepo.NewQuestionRepository(db)
	questionSvc := questionservice.NewQuestionService(questionRepository)
	questionHandler := questionhttp.NewHandler(questionSvc)

	gameRepository := gamerepo.NewGameRepository(db)
	userLookup := gameservice.NewUserLookupAdapter(userRepo)
	gameSvc := gameservice.NewGameService(gameRepository, questionRepository, userLookup)
	gameHandler := gamehttp.NewHandler(gameSvc)

	testingHandler := testinghttp.NewHandler(db)

	// Build Echo directly so we can skip the rate limiter in tests.
	e := echo.New()
	e.Validator = validation.NewValidator()
	e.HTTPErrorHandler = appmiddleware.CustomHTTPErrorHandler
	e.Use(appmiddleware.SecurityHeadersMiddleware())
	e.Use(appmiddleware.RequestIDMiddleware())
	e.Use(appmiddleware.RecoverMiddleware())
	e.Use(appmiddleware.CORSMiddleware(cfg.CorsAllowedOrigins))
	// RateLimiterMiddleware intentionally omitted for e2e tests.

	authMW := auth.JWTMiddleware(tokenManager)
	saMW := auth.BasicAuthMiddleware(cfg.SAAdminLogin, cfg.SAAdminPassword)

	healthhttp.RegisterRoutes(e, healthHandler)
	userhttp.RegisterRoutes(e, userHandler, authMW)
	authhttp.RegisterRoutes(e, authHandler, authMW)
	userhttp.RegisterSARoutes(e, saUserHandler, saMW)
	questionhttp.RegisterRoutes(e, questionHandler, saMW)
	gamehttp.RegisterRoutes(e, gameHandler, authMW)
	testinghttp.RegisterRoutes(e, testingHandler)

	ts := httptest.NewServer(e)
	t.Cleanup(func() {
		ts.Close()
		_ = db.Close()
	})

	return &testServer{
		server: ts,
		client: ts.Client(),
		saUser: cfg.SAAdminLogin + ":" + cfg.SAAdminPassword,
	}
}

// url builds the full URL for a path.
func (ts *testServer) url(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return ts.server.URL + path
}

// do executes an HTTP request and returns status + body bytes.
func (ts *testServer) do(t *testing.T, method, path string, body any, headers map[string]string) (int, []byte) {
	t.Helper()
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		require.NoError(t, err)
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, ts.url(path), bodyReader)
	require.NoError(t, err)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := ts.client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return resp.StatusCode, respBody
}

// basicAuthHeader returns an Authorization header value for SA basic auth.
func (ts *testServer) basicAuthHeader() map[string]string {
	encoded := base64.StdEncoding.EncodeToString([]byte(ts.saUser))
	return map[string]string{"Authorization": "Basic " + encoded}
}

// bearerHeader returns an Authorization header for JWT bearer auth.
func bearerHeader(token string) map[string]string {
	return map[string]string{"Authorization": "Bearer " + token}
}

// deleteAllData clears the database via the testing endpoint.
func (ts *testServer) deleteAllData(t *testing.T) {
	t.Helper()
	status, _ := ts.do(t, http.MethodDelete, "/api/testing/all-data", nil, nil)
	require.Equal(t, http.StatusNoContent, status)
}

// createQuestion creates a quiz question via SA API and returns its ID.
func (ts *testServer) createQuestion(t *testing.T, body string, correctAnswers []string) string {
	t.Helper()
	payload := map[string]any{
		"body":           body,
		"correctAnswers": correctAnswers,
	}
	status, respBody := ts.do(t, http.MethodPost, "/api/sa/quiz/questions", payload, ts.basicAuthHeader())
	require.Equal(t, http.StatusCreated, status, "createQuestion: %s", respBody)

	var resp struct {
		ID string `json:"id"`
	}
	require.NoError(t, json.Unmarshal(respBody, &resp))
	require.NotEmpty(t, resp.ID)
	return resp.ID
}

// publishQuestion publishes a question via SA API.
func (ts *testServer) publishQuestion(t *testing.T, questionID string) {
	t.Helper()
	status, respBody := ts.do(t, http.MethodPut,
		"/api/sa/quiz/questions/"+questionID+"/publish",
		map[string]bool{"published": true},
		ts.basicAuthHeader(),
	)
	require.Equal(t, http.StatusNoContent, status, "publishQuestion: %s", respBody)
}

// createAndPublishQuestions creates and publishes n questions, returns their IDs.
func (ts *testServer) createAndPublishQuestions(t *testing.T, n int) []string {
	t.Helper()
	ids := make([]string, n)
	for i := 0; i < n; i++ {
		body := fmt.Sprintf("Question number %d: what is %d+%d?", i+1, i, i+1)
		answer := fmt.Sprintf("%d", i+i+1)
		ids[i] = ts.createQuestion(t, body, []string{answer, "correct answer"})
		ts.publishQuestion(t, ids[i])
	}
	return ids
}

// saCreateUser creates a user via SA API, returns the user ID.
func (ts *testServer) saCreateUser(t *testing.T, login, email, password string) string {
	t.Helper()
	payload := map[string]string{
		"login":    login,
		"email":    email,
		"password": password,
	}
	status, respBody := ts.do(t, http.MethodPost, "/api/sa/users", payload, ts.basicAuthHeader())
	require.Equal(t, http.StatusCreated, status, "saCreateUser login=%s: %s", login, respBody)

	var resp struct {
		ID string `json:"id"`
	}
	require.NoError(t, json.Unmarshal(respBody, &resp))
	require.NotEmpty(t, resp.ID)
	return resp.ID
}

// login logs in a user and returns the access token.
func (ts *testServer) login(t *testing.T, loginOrEmail, password string) string {
	t.Helper()
	payload := map[string]string{
		"loginOrEmail": loginOrEmail,
		"password":     password,
	}
	status, respBody := ts.do(t, http.MethodPost, "/api/auth/login", payload, nil)
	require.Equal(t, http.StatusOK, status, "login %s: %s", loginOrEmail, respBody)

	var resp struct {
		AccessToken string `json:"accessToken"`
	}
	require.NoError(t, json.Unmarshal(respBody, &resp))
	require.NotEmpty(t, resp.AccessToken)
	return resp.AccessToken
}

// connectToGame calls POST /api/pair-game-quiz/pairs/connection and returns the game response body.
func (ts *testServer) connectToGame(t *testing.T, token string) (int, []byte) {
	t.Helper()
	return ts.do(t, http.MethodPost, "/api/pair-game-quiz/pairs/connection", nil, bearerHeader(token))
}

// submitAnswer calls POST /api/pair-game-quiz/pairs/my-current/answers.
func (ts *testServer) submitAnswer(t *testing.T, token, answer string) (int, []byte) {
	t.Helper()
	return ts.do(t, http.MethodPost, "/api/pair-game-quiz/pairs/my-current/answers",
		map[string]string{"answer": answer},
		bearerHeader(token),
	)
}

// getMyCurrentGame calls GET /api/pair-game-quiz/pairs/my-current.
func (ts *testServer) getMyCurrentGame(t *testing.T, token string) (int, []byte) {
	t.Helper()
	return ts.do(t, http.MethodGet, "/api/pair-game-quiz/pairs/my-current", nil, bearerHeader(token))
}

// getMyGames calls GET /api/pair-game-quiz/pairs/my.
func (ts *testServer) getMyGames(t *testing.T, token, rawQuery string) (int, []byte) {
	t.Helper()

	path := "/api/pair-game-quiz/pairs/my"
	if rawQuery != "" {
		path += "?" + rawQuery
	}

	return ts.do(t, http.MethodGet, path, nil, bearerHeader(token))
}

// getMyStatistic calls GET /api/pair-game-quiz/users/my-statistic.
func (ts *testServer) getMyStatistic(t *testing.T, token string) (int, []byte) {
	t.Helper()
	return ts.do(t, http.MethodGet, "/api/pair-game-quiz/users/my-statistic", nil, bearerHeader(token))
}

// getGameByID calls GET /api/pair-game-quiz/pairs/:id.
func (ts *testServer) getGameByID(t *testing.T, token, id string) (int, []byte) {
	t.Helper()
	return ts.do(t, http.MethodGet, "/api/pair-game-quiz/pairs/"+id, nil, bearerHeader(token))
}

// mustUnmarshalGame unmarshals response body into a game map.
func mustUnmarshalGame(t *testing.T, body []byte) map[string]any {
	t.Helper()
	var g map[string]any
	require.NoError(t, json.Unmarshal(body, &g), "body: %s", body)
	return g
}
