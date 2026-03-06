# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

All Go commands run from the repo root (no `backend/` subdirectory needed — the Makefile's `cd backend` references are stale from a previous project template):

```bash
# Build
go build -o bin/server ./cmd/server/

# Run server
go run ./cmd/server/

# Run all tests
go test ./...

# Run a single test
go test ./internal/pkg/apperrors/... -run TestAppError
go test ./internal/domain/user/service/... -run TestUserService_GetByID

# Run tests with verbose output
go test -v ./internal/...

# Format
go fmt ./...

# Lint (requires golangci-lint)
golangci-lint run

# Generate mocks (mockery)
go generate ./...
```

### Database migrations

```bash
# Apply migrations (requires DATABASE_URL env var)
go run cmd/migrate/main.go -action up

# Rollback
go run cmd/migrate/main.go -action down

# Create new migration
migrate create -ext sql -dir internal/app/database/migrations <name>
```

### E2E tests

Require a running PostgreSQL database (`DATABASE_URL` env var). Tests auto-skip if DB is unreachable.

```bash
# Run all e2e tests (race detector + coverage + summary)
make test-e2e

# Run a single e2e test
go test -tags e2e -v -run TestConcurrentJoin -timeout 60s ./test/e2e/
```

## Environment Variables

Required for non-development usage:

| Variable | Default | Notes |
|----------|---------|-------|
| `DATABASE_URL` | `postgres://user:password@localhost:5432/quiz_db?sslmode=disable` | PostgreSQL connection string |
| `JWT_SECRET` | auto-generated | Must be set in production |
| `SERVER_PORT` | `8080` | |
| `SERVER_ENV` | `development` | |
| `SA_ADMIN_LOGIN` | `admin` | Basic auth for SA API |
| `SA_ADMIN_PASSWORD` | `admin` | Basic auth for SA API — CI test runner expects `qwerty` |
| `CORS_ALLOWED_ORIGINS` | `http://localhost:3000` | Comma-separated list |
| `JWT_EXPIRY_MINUTES` | `10` | Access token lifetime (spec requires ≥5 min) |

Copy `.env.example` → `.env` if present, or export vars directly.

## Architecture

### Layer structure (domain-driven)

Each domain follows: `delivery/http` → `service` → `repository` → DB

```
internal/
├── app/
│   ├── app.go              # Wires all domains together (DI root)
│   ├── config/             # Env-based config
│   ├── database/           # postgres.go (pgx/v5 + sqlx), migrations/
│   ├── middleware/         # Echo middleware: CORS, rate limit, error handler, etc.
│   └── server/             # Echo setup, middleware pipeline
├── domain/
│   ├── auth/               # Login (email+password), token refresh, code exchange
│   ├── game/               # Pair quiz game logic
│   ├── health/             # GET /healthz
│   ├── question/           # SA CRUD for quiz questions
│   └── user/               # Player registration/profile + SA user management
└── pkg/
    ├── apperrors/          # Unified AppError type with HTTP codes
    ├── auth/               # JWT token manager, BasicAuth middleware, CodeStore
    ├── helpers/            # UUID gen, request helpers, testutil
    ├── logger/             # Structured JSON logger
    └── validation/         # go-playground/validator wrapper with custom tags
```

### Authentication

Two separate auth schemes, applied at route registration in `app.go`:

- **Player API** (`/api/...`): JWT Bearer token — `auth.JWTMiddleware`. User ID stored in context as `"user_id"`.
- **SA (super-admin) API** (`/sa/...`): HTTP Basic Auth — `auth.BasicAuthMiddleware`. Credentials from `SA_ADMIN_LOGIN` / `SA_ADMIN_PASSWORD`.

### Error handling

All handlers return `*apperrors.AppError` values. The centralized handler in `middleware.CustomHTTPErrorHandler` converts them to the spec-compliant JSON format:

```json
{"errorsMessages": [{"message": "...", "field": "..."}]}
```

Never return raw `error` or `echo.HTTPError` from domain handlers — always use `apperrors.*` constructors.

When `c.Bind(&req)` fails due to a JSON type mismatch (e.g. string sent for a bool field), detect `*json.UnmarshalTypeError` to return a field-specific error:

```go
var typeErr *json.UnmarshalTypeError
if errors.As(err, &typeErr) {
    return apperrors.NewValidationError(map[string]string{typeErr.Field: "must be a boolean"})
}
return apperrors.BadRequest("Invalid request body")
```

### Validation

Struct validation uses `go-playground/validator/v10` with a custom wrapper. Custom tags:
- `login_pattern` — enforces `^[a-zA-Z0-9_-]+$` (letters, digits, `_`, `-`)

Call `c.Validate(&req)` in handlers; the validator returns `*apperrors.AppError` with field-level details.

`internal/pkg/validation/validator.go` uses `RegisterTagNameFunc` so validation error field names use JSON tag names (`"body"`, `"loginOrEmail"`) instead of Go struct field names (`"Body"`, `"LoginOrEmail"`). This is required for spec-compliant error responses.

### API Patterns

Key endpoint behaviors confirmed against the spec:

- **`POST /api/auth/login`**: accepts `{loginOrEmail, password}` (not `email`). Returns **only** `{accessToken}`. Tries email lookup first, falls back to login username (`LoginByLoginOrEmail` in user service).
- **`GET /api/auth/me`**: returns `{email, login, userId}` for the JWT-authenticated caller.
- **`POST /pair-game-quiz/pairs/my-current/answers`**: returns **403** when the caller is not currently in any active game (not 404). Uses `ErrGameNotActive` sentinel.
- **`GET /pair-game-quiz/pairs/:id`**: returns **400** (not 404) when `:id` is not a valid UUID format. Uses `ErrInvalidGameID` sentinel.
- **`POST /pair-game-quiz/pairs/my-current/answers` response**: `{questionId, answerStatus, addedAt}` where `answerStatus` is `"Correct"` or `"Incorrect"`.

### Game domain specifics

Quiz games are pair-based (2 players). Key invariants:
- **5 questions** selected randomly from published questions when second player joins.
- A player can only be in one active/pending game at a time.
- Bonus point: if a player finishes all answers first AND has ≥1 correct answer.
- Game statuses: `PendingSecondPlayer` → `Active` → `Finished`.
- While `PendingSecondPlayer`: `secondPlayerProgress`, `questions`, `startGameDate`, `finishGameDate` are `null`.

`GameService.SubmitAnswer` answers the **next unanswered question** for the player (sequential, one attempt per question).

### Database schema

Two migration files:
- `000001_init_schema` — `users` table with `login` (nullable varchar, unique), `email`, `password_hash`, etc.
- `000004_quiz_schema` — `quiz_questions`, `quiz_games`, `quiz_game_questions`, `quiz_game_answers`.

### Testing patterns

See [`docs/dev/testing-patterns.md`](docs/dev/testing-patterns.md) for detailed patterns including:
- Goroutine-safe channel pattern for concurrent e2e tests
- TOCTOU race condition verification workflow
- pgconn UNIQUE violation detection
- `git restore --source=HEAD` gotcha

### Spec compliance

The canonical API spec is `docs/specification/h28.sa.json` (OpenAPI) and `docs/specification/README.md`. Key constraints enforced in code:
- `pageSize` max is **20** (capped in SA handlers for questions and users).
- Login field pattern: `^[a-zA-Z0-9_-]*$`.
- `GET /pair-game-quiz/pairs/my-current` returns 404 if no active/pending game.

## Known Issues / Gotchas

See [`docs/dev/testing-patterns.md`](docs/dev/testing-patterns.md) for full details. Quick reference:
- **`git restore <file>`** restores from index, not HEAD — use `git restore --source=HEAD <file>` after `git checkout <commit>^ -- <file>`.
- **pgx UNIQUE violation**: use `database.IsUniqueViolation(err)` from `internal/app/database/pgxerrors.go`.
- **Rate limit middleware** returns `*apperrors.AppError`, not writes via `c.JSON`.
