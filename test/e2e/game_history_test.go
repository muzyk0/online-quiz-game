//go:build e2e

package e2e

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type answerAttempt struct {
	token  string
	answer string
}

type myGamesResponse struct {
	PagesCount int              `json:"pagesCount"`
	Page       int              `json:"page"`
	PageSize   int              `json:"pageSize"`
	TotalCount int              `json:"totalCount"`
	Items      []myGameListItem `json:"items"`
}

type myGameListItem struct {
	ID              string     `json:"id"`
	Status          string     `json:"status"`
	PairCreatedDate time.Time  `json:"pairCreatedDate"`
	StartGameDate   *time.Time `json:"startGameDate"`
	FinishGameDate  *time.Time `json:"finishGameDate"`
}

type myStatisticResponse struct {
	SumScore    int     `json:"sumScore"`
	AvgScores   float64 `json:"avgScores"`
	GamesCount  int     `json:"gamesCount"`
	WinsCount   int     `json:"winsCount"`
	LossesCount int     `json:"lossesCount"`
	DrawsCount  int     `json:"drawsCount"`
}

func TestMyGamesHistory(t *testing.T) {
	ts := newTestServer(t)
	tokens := setup(t, ts, 2)

	firstFinished := playGame(
		t,
		ts,
		tokens[0],
		tokens[1],
		firstPlayerWinsThreeToTwo(tokens[0], tokens[1]),
	)
	assertFinishedScore(t, firstFinished, 3, 2)

	secondFinished := playGame(
		t,
		ts,
		tokens[0],
		tokens[1],
		drawTwoToTwo(tokens[0], tokens[1]),
	)
	assertFinishedScore(t, secondFinished, 2, 2)

	thirdFinished := playGame(
		t,
		ts,
		tokens[0],
		tokens[1],
		secondPlayerWinsFiveToZero(tokens[0], tokens[1]),
	)
	assertFinishedScore(t, thirdFinished, 0, 5)

	activeGameID := startGame(t, ts, tokens[0], tokens[1])

	status, body := ts.getMyGames(t, tokens[0], "sortBy=status&sortDirection=desc&pageNumber=1&pageSize=10")
	require.Equal(t, http.StatusOK, status, "getMyGames: %s", body)

	var response myGamesResponse
	require.NoError(t, json.Unmarshal(body, &response), "body: %s", body)

	assert.Equal(t, 1, response.Page)
	assert.Equal(t, 10, response.PageSize)
	assert.Equal(t, 1, response.PagesCount)
	assert.Equal(t, 4, response.TotalCount)
	require.Len(t, response.Items, 4)

	assert.Equal(t, []string{"Finished", "Finished", "Finished", "Active"}, collectStatuses(response.Items))
	assert.Equal(t, thirdFinished["id"], response.Items[0].ID)
	assert.Equal(t, secondFinished["id"], response.Items[1].ID)
	assert.Equal(t, firstFinished["id"], response.Items[2].ID)
	assert.Equal(t, activeGameID, response.Items[3].ID)

	assert.True(t, notBefore(response.Items[0].PairCreatedDate, response.Items[1].PairCreatedDate))
	assert.True(t, notBefore(response.Items[1].PairCreatedDate, response.Items[2].PairCreatedDate))
	assert.NotNil(t, response.Items[0].FinishGameDate)
	assert.NotNil(t, response.Items[1].FinishGameDate)
	assert.NotNil(t, response.Items[2].FinishGameDate)
	assert.NotNil(t, response.Items[3].StartGameDate)
	assert.Nil(t, response.Items[3].FinishGameDate)
}

func TestMyStatistic(t *testing.T) {
	ts := newTestServer(t)
	tokens := setup(t, ts, 4)

	firstFinished := playGame(
		t,
		ts,
		tokens[0],
		tokens[1],
		firstPlayerWinsThreeToTwo(tokens[0], tokens[1]),
	)
	assertFinishedScore(t, firstFinished, 3, 2)

	secondFinished := playGame(
		t,
		ts,
		tokens[0],
		tokens[2],
		drawTwoToTwo(tokens[0], tokens[2]),
	)
	assertFinishedScore(t, secondFinished, 2, 2)

	thirdFinished := playGame(
		t,
		ts,
		tokens[0],
		tokens[3],
		secondPlayerWinsFiveToZero(tokens[0], tokens[3]),
	)
	assertFinishedScore(t, thirdFinished, 0, 5)

	startGame(t, ts, tokens[0], tokens[1])

	status, body := ts.getMyStatistic(t, tokens[0])
	require.Equal(t, http.StatusOK, status, "getMyStatistic: %s", body)

	var response myStatisticResponse
	require.NoError(t, json.Unmarshal(body, &response), "body: %s", body)

	assert.Equal(t, 5, response.SumScore)
	assert.InDelta(t, 1.67, response.AvgScores, 0.001)
	assert.Equal(t, 3, response.GamesCount)
	assert.Equal(t, 1, response.WinsCount)
	assert.Equal(t, 1, response.LossesCount)
	assert.Equal(t, 1, response.DrawsCount)
}

func startGame(t *testing.T, ts *testServer, creatorToken, joinerToken string) string {
	t.Helper()

	status, body := ts.connectToGame(t, creatorToken)
	require.Equal(t, http.StatusOK, status, "create game: %s", body)

	pendingGame := mustUnmarshalGame(t, body)
	gameID, ok := pendingGame["id"].(string)
	require.True(t, ok)
	require.Equal(t, "PendingSecondPlayer", pendingGame["status"])

	status, body = ts.connectToGame(t, joinerToken)
	require.Equal(t, http.StatusOK, status, "join game: %s", body)

	activeGame := mustUnmarshalGame(t, body)
	require.Equal(t, gameID, activeGame["id"])
	require.Equal(t, "Active", activeGame["status"])

	return gameID
}

func playGame(t *testing.T, ts *testServer, creatorToken, joinerToken string, attempts []answerAttempt) map[string]any {
	t.Helper()

	gameID := startGame(t, ts, creatorToken, joinerToken)

	for i, attempt := range attempts {
		status, body := ts.submitAnswer(t, attempt.token, attempt.answer)
		require.Equal(t, http.StatusOK, status, "answer %d: %s", i+1, body)
	}

	status, body := ts.getGameByID(t, creatorToken, gameID)
	require.Equal(t, http.StatusOK, status, "get finished game: %s", body)

	return mustUnmarshalGame(t, body)
}

func assertFinishedScore(t *testing.T, game map[string]any, firstPlayerScore, secondPlayerScore int) {
	t.Helper()

	require.Equal(t, "Finished", game["status"])
	require.NotNil(t, game["finishGameDate"])

	firstProgress := game["firstPlayerProgress"].(map[string]any)
	secondProgress := game["secondPlayerProgress"].(map[string]any)

	assert.Equal(t, float64(firstPlayerScore), firstProgress["score"])
	assert.Equal(t, float64(secondPlayerScore), secondProgress["score"])
}

func collectStatuses(items []myGameListItem) []string {
	statuses := make([]string, len(items))
	for i := range items {
		statuses[i] = items[i].Status
	}
	return statuses
}

func notBefore(left, right time.Time) bool {
	return left.After(right) || left.Equal(right)
}

func firstPlayerWinsThreeToTwo(firstPlayerToken, secondPlayerToken string) []answerAttempt {
	return []answerAttempt{
		{token: firstPlayerToken, answer: "correct answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "correct answer"},
		{token: secondPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "wrong answer"},
		{token: firstPlayerToken, answer: "correct answer"},
		{token: firstPlayerToken, answer: "correct answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
	}
}

func drawTwoToTwo(firstPlayerToken, secondPlayerToken string) []answerAttempt {
	return []answerAttempt{
		{token: firstPlayerToken, answer: "correct answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "correct answer"},
		{token: secondPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "wrong answer"},
		{token: firstPlayerToken, answer: "correct answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
	}
}

func secondPlayerWinsFiveToZero(firstPlayerToken, secondPlayerToken string) []answerAttempt {
	return []answerAttempt{
		{token: secondPlayerToken, answer: "correct answer"},
		{token: secondPlayerToken, answer: "correct answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "correct answer"},
		{token: secondPlayerToken, answer: "correct answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
		{token: firstPlayerToken, answer: "wrong answer"},
		{token: secondPlayerToken, answer: "correct answer"},
	}
}
