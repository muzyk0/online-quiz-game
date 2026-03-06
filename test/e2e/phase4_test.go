//go:build e2e

package e2e

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// gameFinishWait is slightly longer than the 10-second auto-finish timeout
// to give the server time to process the expiry on the next request.
const gameFinishWait = 11 * time.Second

// TestGameTimingFinishAfterWaiting mirrors "Homework 28 > Finish game after waiting 10 sec."
func TestGameTimingFinishAfterWaiting(t *testing.T) {
	ts := newTestServer(t)

	// Test 1: user1 answers all 5 correctly, user2 does nothing.
	// After 10s: game Finished, firstPlayer.score=6 (5+bonus), secondPlayer.score=0.
	t.Run("user1 finishes all 5 correct; user2 idles; after 10s: Finished with score 6:0", func(t *testing.T) {
		tokens := setup(t, ts, 2)

		s, b := ts.connectToGame(t, tokens[0])
		require.Equal(t, http.StatusOK, s, "user1 connect: %s", b)
		gameID := mustUnmarshalGame(t, b)["id"].(string)

		s, b = ts.connectToGame(t, tokens[1])
		require.Equal(t, http.StatusOK, s, "user2 connect: %s", b)

		for i := 0; i < 5; i++ {
			s, b = ts.submitAnswer(t, tokens[0], "correct answer")
			require.Equal(t, http.StatusOK, s, "user1 answer %d: %s", i+1, b)
		}

		time.Sleep(gameFinishWait)

		s, b = ts.getGameByID(t, tokens[0], gameID)
		require.Equal(t, http.StatusOK, s, "getGame: %s", b)
		g := mustUnmarshalGame(t, b)
		assert.Equal(t, "Finished", g["status"])
		assert.NotNil(t, g["finishGameDate"])

		fp := g["firstPlayerProgress"].(map[string]any)
		sp := g["secondPlayerProgress"].(map[string]any)
		assert.Equal(t, float64(6), fp["score"], "user1: 5 correct + 1 bonus")
		assert.Equal(t, float64(0), sp["score"], "user2: no answers")
	})

	// Test 2: user2 answers 3 correct, then user1 finishes all 5.
	// After 10s: my-current by user2 → 404; game → Finished 6:3.
	t.Run("user2 answers 3 correct, user1 finishes; after 10s: my-current 404, game Finished 6:3", func(t *testing.T) {
		tokens := setup(t, ts, 2)

		s, b := ts.connectToGame(t, tokens[0])
		require.Equal(t, http.StatusOK, s)
		gameID := mustUnmarshalGame(t, b)["id"].(string)

		s, b = ts.connectToGame(t, tokens[1])
		require.Equal(t, http.StatusOK, s)

		for i := 0; i < 3; i++ {
			s, b = ts.submitAnswer(t, tokens[1], "correct answer")
			require.Equal(t, http.StatusOK, s, "user2 answer %d: %s", i+1, b)
		}

		for i := 0; i < 5; i++ {
			s, b = ts.submitAnswer(t, tokens[0], "correct answer")
			require.Equal(t, http.StatusOK, s, "user1 answer %d: %s", i+1, b)
		}

		time.Sleep(gameFinishWait)

		s, _ = ts.getMyCurrentGame(t, tokens[1])
		assert.Equal(t, http.StatusNotFound, s, "user2 my-current should be 404 after timeout")

		s, b = ts.getGameByID(t, tokens[0], gameID)
		require.Equal(t, http.StatusOK, s)
		g := mustUnmarshalGame(t, b)
		assert.Equal(t, "Finished", g["status"])
		assert.NotNil(t, g["finishGameDate"])

		fp := g["firstPlayerProgress"].(map[string]any)
		sp := g["secondPlayerProgress"].(map[string]any)
		assert.Equal(t, float64(6), fp["score"], "user1: 5 correct + 1 bonus")
		assert.Equal(t, float64(3), sp["score"], "user2: 3 correct, no bonus")
	})

	// Test 3: user2 finishes first (3 incorrect + 2 correct), user1 has 3 correct.
	// Before timeout: my-current is still Active.
	// After 10s: my-current → 404, game → Finished 3:3.
	t.Run("user2 finishes first with 3 incorrect+2 correct; user1 has 3 correct; after 10s: Finished 3:3", func(t *testing.T) {
		tokens := setup(t, ts, 2)

		s, b := ts.connectToGame(t, tokens[0])
		require.Equal(t, http.StatusOK, s)
		gameID := mustUnmarshalGame(t, b)["id"].(string)

		s, b = ts.connectToGame(t, tokens[1])
		require.Equal(t, http.StatusOK, s)

		for i := 0; i < 3; i++ {
			s, b = ts.submitAnswer(t, tokens[1], "wrong answer")
			require.Equal(t, http.StatusOK, s, "user2 incorrect %d: %s", i+1, b)
		}

		for i := 0; i < 3; i++ {
			s, b = ts.submitAnswer(t, tokens[0], "correct answer")
			require.Equal(t, http.StatusOK, s, "user1 answer %d: %s", i+1, b)
		}

		for i := 0; i < 2; i++ {
			s, b = ts.submitAnswer(t, tokens[1], "correct answer")
			require.Equal(t, http.StatusOK, s, "user2 correct %d: %s", i+1, b)
		}

		// Before timeout: user1 hasn't finished, game is still Active.
		s, _ = ts.getMyCurrentGame(t, tokens[1])
		assert.Equal(t, http.StatusOK, s, "user2 my-current should be 200 before timeout")

		time.Sleep(gameFinishWait)

		s, _ = ts.getMyCurrentGame(t, tokens[1])
		assert.Equal(t, http.StatusNotFound, s, "user2 my-current should be 404 after timeout")

		s, b = ts.getGameByID(t, tokens[1], gameID)
		require.Equal(t, http.StatusOK, s)
		g := mustUnmarshalGame(t, b)
		assert.Equal(t, "Finished", g["status"])
		assert.NotNil(t, g["finishGameDate"])

		fp := g["firstPlayerProgress"].(map[string]any)
		sp := g["secondPlayerProgress"].(map[string]any)
		// user1 (first player): 3 correct, did not finish first → no bonus → 3
		// user2 (second player): 2 correct, finished first, ≥1 correct → +1 bonus → 3
		assert.Equal(t, float64(3), fp["score"], "user1: 3 correct, no bonus")
		assert.Equal(t, float64(3), sp["score"], "user2: 2 correct + 1 bonus")
	})

	// Test 4: two parallel games, both auto-finish after 10s.
	t.Run("two parallel games both auto-finish after 10s", func(t *testing.T) {
		tokens := setup(t, ts, 4)

		// Game1: user1 creates, user2 connects.
		s, b := ts.connectToGame(t, tokens[0])
		require.Equal(t, http.StatusOK, s, "game1 user1 connect: %s", b)
		game1ID := mustUnmarshalGame(t, b)["id"].(string)

		s, b = ts.connectToGame(t, tokens[1])
		require.Equal(t, http.StatusOK, s, "game1 user2 connect: %s", b)

		// Game1: user2 answers 3 incorrect.
		for i := 0; i < 3; i++ {
			s, b = ts.submitAnswer(t, tokens[1], "wrong answer")
			require.Equal(t, http.StatusOK, s, "game1 user2 wrong %d: %s", i+1, b)
		}

		// Game1: user1 answers 4 correct (not finished yet).
		for i := 0; i < 4; i++ {
			s, b = ts.submitAnswer(t, tokens[0], "correct answer")
			require.Equal(t, http.StatusOK, s, "game1 user1 correct %d: %s", i+1, b)
		}

		// Game2: user3 creates, user4 connects.
		s, b = ts.connectToGame(t, tokens[2])
		require.Equal(t, http.StatusOK, s, "game2 user3 connect: %s", b)

		s, b = ts.connectToGame(t, tokens[3])
		require.Equal(t, http.StatusOK, s, "game2 user4 connect: %s", b)
		game2ID := mustUnmarshalGame(t, b)["id"].(string)

		// Game2: user3 answers all 5 correctly (finishes first).
		for i := 0; i < 5; i++ {
			s, b = ts.submitAnswer(t, tokens[2], "correct answer")
			require.Equal(t, http.StatusOK, s, "game2 user3 answer %d: %s", i+1, b)
		}

		// Game2: user4 answers 2 correct (not finished).
		for i := 0; i < 2; i++ {
			s, b = ts.submitAnswer(t, tokens[3], "correct answer")
			require.Equal(t, http.StatusOK, s, "game2 user4 answer %d: %s", i+1, b)
		}

		// Game1: user2 answers 2 correct (finishes, 5 total).
		for i := 0; i < 2; i++ {
			s, b = ts.submitAnswer(t, tokens[1], "correct answer")
			require.Equal(t, http.StatusOK, s, "game1 user2 correct %d: %s", i+1, b)
		}

		time.Sleep(gameFinishWait)

		// Verify game1.
		s, b = ts.getGameByID(t, tokens[1], game1ID)
		require.Equal(t, http.StatusOK, s, "getGame1: %s", b)
		g1 := mustUnmarshalGame(t, b)
		assert.Equal(t, "Finished", g1["status"])
		assert.NotNil(t, g1["finishGameDate"])

		fp1 := g1["firstPlayerProgress"].(map[string]any)
		sp1 := g1["secondPlayerProgress"].(map[string]any)
		// user1 (first player): 4 correct, did not finish first → no bonus → 4
		// user2 (second player): 2 correct, finished first, ≥1 correct → +1 bonus → 3
		assert.Equal(t, float64(4), fp1["score"], "game1 user1 score")
		assert.Equal(t, float64(3), sp1["score"], "game1 user2 score")

		// Verify game2.
		s, b = ts.getGameByID(t, tokens[2], game2ID)
		require.Equal(t, http.StatusOK, s, "getGame2: %s", b)
		g2 := mustUnmarshalGame(t, b)
		assert.Equal(t, "Finished", g2["status"])
		assert.NotNil(t, g2["finishGameDate"])

		fp2 := g2["firstPlayerProgress"].(map[string]any)
		sp2 := g2["secondPlayerProgress"].(map[string]any)
		// user3 (first player): 5 correct, finished first → +1 bonus → 6
		// user4 (second player): 2 correct, did not finish → no bonus → 2
		assert.Equal(t, float64(6), fp2["score"], "game2 user3 score")
		assert.Equal(t, float64(2), sp2["score"], "game2 user4 score")
	})

	// Test 5: after auto-finish both players can join a new game (no 403).
	t.Run("after auto-finish, players can join a new game without 403", func(t *testing.T) {
		tokens := setup(t, ts, 2)

		s, b := ts.connectToGame(t, tokens[0])
		require.Equal(t, http.StatusOK, s)

		s, b = ts.connectToGame(t, tokens[1])
		require.Equal(t, http.StatusOK, s)

		for i := 0; i < 5; i++ {
			s, b = ts.submitAnswer(t, tokens[0], "correct answer")
			require.Equal(t, http.StatusOK, s, "user1 answer %d: %s", i+1, b)
		}

		time.Sleep(gameFinishWait)

		// Both players can now start/join a new game.
		s, b = ts.connectToGame(t, tokens[0])
		assert.Equal(t, http.StatusOK, s, "user1 new game connect: %s", b)

		s, b = ts.connectToGame(t, tokens[1])
		assert.Equal(t, http.StatusOK, s, "user2 new game connect: %s", b)
	})
}
