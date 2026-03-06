# Testing Patterns

## E2E Tests

E2E tests live in `test/e2e/`, require a running PostgreSQL instance (`DATABASE_URL`), and auto-skip if the database is unreachable.

```bash
# Run all e2e tests (race detector + coverage + summary)
make test-e2e

# Run a single e2e test
go test -tags e2e -v -run TestConcurrentJoin -timeout 60s ./test/e2e/
```

The server is started via `httptest.NewServer` with a real Echo application and a real database — no mocks.

## Goroutine-safe Channel Pattern

When `require.*` is called inside a goroutine, `t.FailNow()` triggers `runtime.Goexit()`, which terminates the goroutine before the channel send — causing the test to hang. Fix: put the send in a `defer`:

```go
ch := make(chan result, 2)
for _, tok := range tokens {
    tok := tok
    go func() {
        r := result{}
        defer func() { ch <- r }() // runs even on runtime.Goexit
        s, b := ts.connectToGame(t, tok)
        r.statusCode = s
        if s == http.StatusOK {
            var g map[string]any
            if json.Unmarshal(b, &g) == nil {
                r.gameID, _ = g["id"].(string)
                r.gameStatus, _ = g["status"].(string)
            }
        }
    }()
}
r1, r2 := <-ch, <-ch
```

## Testing TOCTOU Race Conditions

To verify that a race condition fix actually works:

1. Restore old files: `git checkout <fix-commit>^ -- <files>`
2. Add `time.Sleep(200ms)` in the repository method between SELECT and UPDATE/INSERT to widen the TOCTOU window
3. Run the e2e test — it should fail (500 / duplicate-key error)
4. Restore the fixed code:
   ```bash
   git restore --source=HEAD <files>
   git reset HEAD <files>
   ```
5. Run the test again — it should pass

**Why the artificial delay is needed:** without it, goroutines making HTTP calls to a local test server rarely overlap at the database level. `time.Sleep(200ms)` guarantees both SELECTs complete before the first UPDATE.

**Example:** `TestConcurrentJoin` in `test/e2e/game_test.go` — regression test for commit `d3046b0` (`FindPendingAndActivate` with `SELECT FOR UPDATE SKIP LOCKED`).

## Known Issues / Gotchas

### `git restore` restores from the index, not HEAD

`git checkout <commit>^ -- <file>` updates **both the index and the working tree**. After that, `git restore <file>` will restore the old version again (from the index). Specify the source explicitly:

```bash
git restore --source=HEAD <file>
git reset HEAD <file>
```

### pgx UNIQUE violation detection

Use `database.IsUniqueViolation(err)` from `internal/app/database/pgxerrors.go`:

```go
if database.IsUniqueViolation(err) {
    return nil, ErrDuplicate
}
```

The helper wraps `pgconn.PgError` + SQLSTATE `23505` so repositories don't need to import `pgconn` directly.

### Rate limit middleware returns AppError

`AuthRateLimitMiddleware` returns `*apperrors.AppError`, it does not write via `c.JSON`. Tests must assert through `require.ErrorAs`, not `rec.Code`:

```go
err := handler(c)
require.Error(t, err)
var appErr *apperrors.AppError
require.ErrorAs(t, err, &appErr)
assert.Equal(t, http.StatusTooManyRequests, appErr.Code)
```
