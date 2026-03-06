//go:build e2e

package e2e

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type topPlayerItem struct {
	SumScore    int     `json:"sumScore"`
	AvgScores   float64 `json:"avgScores"`
	GamesCount  int     `json:"gamesCount"`
	WinsCount   int     `json:"winsCount"`
	LossesCount int     `json:"lossesCount"`
	DrawsCount  int     `json:"drawsCount"`
	Player      struct {
		ID    string `json:"id"`
		Login string `json:"login"`
	} `json:"player"`
}

type topPlayersResponse struct {
	PagesCount int             `json:"pagesCount"`
	Page       int             `json:"page"`
	PageSize   int             `json:"pageSize"`
	TotalCount int             `json:"totalCount"`
	Items      []topPlayerItem `json:"items"`
}

// TestTopPlayersEndpoint verifies GET /api/pair-game-quiz/users/top
// is public (no auth), paginated, and sortable.
func TestTopPlayersEndpoint(t *testing.T) {
	ts := newTestServer(t)
	ts.deleteAllData(t)

	// Setup: create 3 users and play several games.
	tokens := setup(t, ts, 3)
	t1, t2, t3 := tokens[0], tokens[1], tokens[2]

	// Game 1: user1 (first) vs user2 (second) → user1 wins 3-2
	playGame(t, ts, t1, t2, firstPlayerWinsThreeToTwo(t1, t2))

	// Game 2: user1 (first) vs user3 (second) → user1 wins 3-2
	playGame(t, ts, t1, t3, firstPlayerWinsThreeToTwo(t1, t3))

	// Game 3: user2 (first) vs user3 (second) → user2 wins 3-2
	playGame(t, ts, t2, t3, firstPlayerWinsThreeToTwo(t2, t3))

	// GET /api/pair-game-quiz/users/top — no auth required
	status, body := ts.do(t, http.MethodGet, "/api/pair-game-quiz/users/top", nil, nil)
	require.Equal(t, http.StatusOK, status, "top players endpoint should be public; body: %s", body)

	var resp topPlayersResponse
	require.NoError(t, json.Unmarshal(body, &resp))

	assert.Equal(t, 3, resp.TotalCount)
	assert.Equal(t, 1, resp.Page)
	assert.Equal(t, 10, resp.PageSize)
	assert.Equal(t, 1, resp.PagesCount)
	assert.Len(t, resp.Items, 3)

	// Default sort: avgScores desc, sumScore desc — user with most wins/scores first.
	// All players have same avg (2.5 with bonus) but different sumScores.
	for _, item := range resp.Items {
		assert.NotEmpty(t, item.Player.ID)
		assert.NotEmpty(t, item.Player.Login)
		assert.GreaterOrEqual(t, item.GamesCount, 1)
	}
}

// TestTopPlayersPagination verifies pageSize and pageNumber params.
func TestTopPlayersPagination(t *testing.T) {
	ts := newTestServer(t)
	ts.deleteAllData(t)

	tokens := setup(t, ts, 3)
	t1, t2, t3 := tokens[0], tokens[1], tokens[2]

	playGame(t, ts, t1, t2, firstPlayerWinsThreeToTwo(t1, t2))
	playGame(t, ts, t1, t3, firstPlayerWinsThreeToTwo(t1, t3))
	playGame(t, ts, t2, t3, firstPlayerWinsThreeToTwo(t2, t3))

	// Page 1 with pageSize=2
	status, body := ts.do(t, http.MethodGet, "/api/pair-game-quiz/users/top?pageSize=2&pageNumber=1", nil, nil)
	require.Equal(t, http.StatusOK, status)

	var p1 topPlayersResponse
	require.NoError(t, json.Unmarshal(body, &p1))
	assert.Equal(t, 3, p1.TotalCount)
	assert.Equal(t, 2, p1.PageSize)
	assert.Equal(t, 2, p1.PagesCount)
	assert.Len(t, p1.Items, 2)

	// Page 2 with pageSize=2
	status, body = ts.do(t, http.MethodGet, "/api/pair-game-quiz/users/top?pageSize=2&pageNumber=2", nil, nil)
	require.Equal(t, http.StatusOK, status)

	var p2 topPlayersResponse
	require.NoError(t, json.Unmarshal(body, &p2))
	assert.Len(t, p2.Items, 1)
	assert.Equal(t, 2, p2.Page)
}

// TestTopPlayersCustomSort verifies the sort query parameter.
func TestTopPlayersCustomSort(t *testing.T) {
	ts := newTestServer(t)
	ts.deleteAllData(t)

	tokens := setup(t, ts, 3)
	t1, t2, t3 := tokens[0], tokens[1], tokens[2]

	playGame(t, ts, t1, t2, firstPlayerWinsThreeToTwo(t1, t2))
	playGame(t, ts, t1, t3, firstPlayerWinsThreeToTwo(t1, t3))
	playGame(t, ts, t2, t3, firstPlayerWinsThreeToTwo(t2, t3))

	// Sort by winsCount desc (URL-encode the space)
	path := "/api/pair-game-quiz/users/top?sort=" + url.QueryEscape("winsCount desc")
	status, body := ts.do(t, http.MethodGet, path, nil, nil)
	require.Equal(t, http.StatusOK, status)

	var resp topPlayersResponse
	require.NoError(t, json.Unmarshal(body, &resp))
	require.Len(t, resp.Items, 3)

	// Verify descending winsCount order
	for i := 1; i < len(resp.Items); i++ {
		assert.GreaterOrEqual(t, resp.Items[i-1].WinsCount, resp.Items[i].WinsCount)
	}
}
