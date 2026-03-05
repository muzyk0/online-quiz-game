//go:build e2e

package e2e

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSAQuestionsCRUD mirrors "Homework 25 > Quiz questions CRUD by SA".
func TestSAQuestionsCRUD(t *testing.T) {
	ts := newTestServer(t)

	t.Run("DELETE /testing/all-data: should remove all data; status 204", func(t *testing.T) {
		ts.deleteAllData(t)
	})

	t.Run("POST /sa/quiz/questions: should create new question; status 201", func(t *testing.T) {
		ts.deleteAllData(t)

		id := ts.createQuestion(t, "What is the capital of France?", []string{"Paris"})
		assert.NotEmpty(t, id)

		// Verify via GET
		status, body := ts.do(t, http.MethodGet, "/api/sa/quiz/questions", nil, ts.basicAuthHeader())
		require.Equal(t, http.StatusOK, status)

		var resp map[string]any
		require.NoError(t, json.Unmarshal(body, &resp))
		items := resp["items"].([]any)
		assert.Len(t, items, 1)
		q := items[0].(map[string]any)
		assert.Equal(t, "What is the capital of France?", q["body"])
		assert.Equal(t, false, q["published"])
	})

	t.Run("GET /sa/quiz/questions: should return questions array with pagination; status 200", func(t *testing.T) {
		ts.deleteAllData(t)

		// Create 5 questions
		for i := 0; i < 5; i++ {
			ts.createQuestion(t,
				strings.Repeat("a", 10)+" question number "+string(rune('0'+i)),
				[]string{"answer"},
			)
		}

		status, body := ts.do(t, http.MethodGet, "/api/sa/quiz/questions", nil, ts.basicAuthHeader())
		require.Equal(t, http.StatusOK, status)

		var resp map[string]any
		require.NoError(t, json.Unmarshal(body, &resp))
		assert.Equal(t, float64(1), resp["page"])
		assert.Equal(t, float64(5), resp["totalCount"])
		items := resp["items"].([]any)
		assert.Len(t, items, 5)
	})

	t.Run("PUT /sa/quiz/questions/:id: should update quiz question; status 204", func(t *testing.T) {
		ts.deleteAllData(t)

		id := ts.createQuestion(t, "Original question body text", []string{"original answer"})

		status, body := ts.do(t, http.MethodPut, "/api/sa/quiz/questions/"+id,
			map[string]any{
				"body":           "Updated question body text here",
				"correctAnswers": []string{"updated answer"},
			},
			ts.basicAuthHeader(),
		)
		require.Equal(t, http.StatusNoContent, status, "update: %s", body)

		// Verify update via GET
		getStatus, getBody := ts.do(t, http.MethodGet, "/api/sa/quiz/questions", nil, ts.basicAuthHeader())
		require.Equal(t, http.StatusOK, getStatus)

		var resp map[string]any
		require.NoError(t, json.Unmarshal(getBody, &resp))
		items := resp["items"].([]any)
		require.Len(t, items, 1)
		q := items[0].(map[string]any)
		assert.Equal(t, "Updated question body text here", q["body"])
	})

	t.Run("DELETE /sa/quiz/questions/:id: should delete question by id; status 204", func(t *testing.T) {
		ts.deleteAllData(t)

		id := ts.createQuestion(t, "Question to be deleted soon", []string{"answer"})

		status, body := ts.do(t, http.MethodDelete, "/api/sa/quiz/questions/"+id, nil, ts.basicAuthHeader())
		require.Equal(t, http.StatusNoContent, status, "delete: %s", body)

		// Verify deletion
		getStatus, getBody := ts.do(t, http.MethodGet, "/api/sa/quiz/questions", nil, ts.basicAuthHeader())
		require.Equal(t, http.StatusOK, getStatus)

		var resp map[string]any
		require.NoError(t, json.Unmarshal(getBody, &resp))
		assert.Equal(t, float64(0), resp["totalCount"])
	})

	t.Run("PUT /sa/quiz/questions/:id/publish: should update publish status; status 204", func(t *testing.T) {
		ts.deleteAllData(t)

		id := ts.createQuestion(t, "Question to be published now", []string{"answer"})

		// Publish
		status, body := ts.do(t, http.MethodPut, "/api/sa/quiz/questions/"+id+"/publish",
			map[string]bool{"published": true},
			ts.basicAuthHeader(),
		)
		require.Equal(t, http.StatusNoContent, status, "publish: %s", body)

		// Verify published=true
		getStatus, getBody := ts.do(t, http.MethodGet, "/api/sa/quiz/questions", nil, ts.basicAuthHeader())
		require.Equal(t, http.StatusOK, getStatus)

		var resp map[string]any
		require.NoError(t, json.Unmarshal(getBody, &resp))
		items := resp["items"].([]any)
		require.Len(t, items, 1)
		q := items[0].(map[string]any)
		assert.Equal(t, true, q["published"])

		// Unpublish
		status2, body2 := ts.do(t, http.MethodPut, "/api/sa/quiz/questions/"+id+"/publish",
			map[string]bool{"published": false},
			ts.basicAuthHeader(),
		)
		require.Equal(t, http.StatusNoContent, status2, "unpublish: %s", body2)
	})

	t.Run("GET,POST,PUT,DELETE /sa/quiz/questions: should return 401 if auth credentials incorrect", func(t *testing.T) {
		ts.deleteAllData(t)

		wrongAuth := map[string]string{"Authorization": "Basic " + "wronguser:wrongpass"}

		// POST with wrong auth
		s, _ := ts.do(t, http.MethodPost, "/api/sa/quiz/questions",
			map[string]any{"body": "Question body here", "correctAnswers": []string{"answer"}},
			wrongAuth,
		)
		assert.Equal(t, http.StatusUnauthorized, s)

		// Create a question with correct auth for remaining tests
		id := ts.createQuestion(t, "Question for auth test here", []string{"answer"})

		// GET with wrong auth
		s, _ = ts.do(t, http.MethodGet, "/api/sa/quiz/questions", nil, wrongAuth)
		assert.Equal(t, http.StatusUnauthorized, s)

		// PUT with wrong auth
		s, _ = ts.do(t, http.MethodPut, "/api/sa/quiz/questions/"+id,
			map[string]any{"body": "Updated body text here", "correctAnswers": []string{"a"}},
			wrongAuth,
		)
		assert.Equal(t, http.StatusUnauthorized, s)

		// DELETE with wrong auth
		s, _ = ts.do(t, http.MethodDelete, "/api/sa/quiz/questions/"+id, nil, wrongAuth)
		assert.Equal(t, http.StatusUnauthorized, s)
	})

	t.Run("PUT,DELETE /sa/quiz/questions/:id: should return 404 if id not found", func(t *testing.T) {
		nonExistentID := "00000000-0000-0000-0000-000000000001"

		s, _ := ts.do(t, http.MethodPut, "/api/sa/quiz/questions/"+nonExistentID,
			map[string]any{"body": "Updated body text here", "correctAnswers": []string{"answer"}},
			ts.basicAuthHeader(),
		)
		assert.Equal(t, http.StatusNotFound, s)

		s, _ = ts.do(t, http.MethodDelete, "/api/sa/quiz/questions/"+nonExistentID, nil, ts.basicAuthHeader())
		assert.Equal(t, http.StatusNotFound, s)
	})
}

// TestSAQuestionsBodyValidation mirrors "Homework 25 > Quiz questions CRUD by SA > Questions body validation".
func TestSAQuestionsBodyValidation(t *testing.T) {
	ts := newTestServer(t)
	ts.deleteAllData(t)

	// Create one question for PUT/publish tests
	questionID := ts.createQuestion(t, "Valid question for validation tests", []string{"answer"})

	t.Run("POST /sa/quiz/questions: should return 400 if body is incorrect", func(t *testing.T) {
		tooLong := strings.Repeat("x", 501)

		// body too long
		s, body := ts.do(t, http.MethodPost, "/api/sa/quiz/questions",
			map[string]any{"body": tooLong, "correctAnswers": []string{"answer"}},
			ts.basicAuthHeader(),
		)
		assert.Equal(t, http.StatusBadRequest, s)
		assertErrorField(t, body, "body")

		// body too short
		s, body = ts.do(t, http.MethodPost, "/api/sa/quiz/questions",
			map[string]any{"body": "short", "correctAnswers": []string{"answer"}},
			ts.basicAuthHeader(),
		)
		assert.Equal(t, http.StatusBadRequest, s)
		assertErrorField(t, body, "body")

		// missing correctAnswers
		s, body = ts.do(t, http.MethodPost, "/api/sa/quiz/questions",
			map[string]any{"body": "Valid question body text here"},
			ts.basicAuthHeader(),
		)
		assert.Equal(t, http.StatusBadRequest, s)
		assertErrorField(t, body, "correctAnswers")

		// empty correctAnswers array
		s, body = ts.do(t, http.MethodPost, "/api/sa/quiz/questions",
			map[string]any{"body": "Valid question body text here", "correctAnswers": []string{}},
			ts.basicAuthHeader(),
		)
		assert.Equal(t, http.StatusBadRequest, s)
		assertErrorField(t, body, "correctAnswers")
	})

	t.Run("PUT /sa/quiz/questions/:id: should return 400 if body is incorrect", func(t *testing.T) {
		tooLong := strings.Repeat("x", 501)

		// body too long
		s, body := ts.do(t, http.MethodPut, "/api/sa/quiz/questions/"+questionID,
			map[string]any{"body": tooLong, "correctAnswers": []string{"answer"}},
			ts.basicAuthHeader(),
		)
		assert.Equal(t, http.StatusBadRequest, s)
		assertErrorField(t, body, "body")

		// missing correctAnswers
		s, body = ts.do(t, http.MethodPut, "/api/sa/quiz/questions/"+questionID,
			map[string]any{"body": "Valid updated question body text"},
			ts.basicAuthHeader(),
		)
		assert.Equal(t, http.StatusBadRequest, s)
		assertErrorField(t, body, "correctAnswers")
	})

	t.Run("PUT /sa/quiz/questions/:id/publish: should return 400 if body is incorrect", func(t *testing.T) {
		// Send string "true" instead of boolean true
		s, body := ts.do(t, http.MethodPut, "/api/sa/quiz/questions/"+questionID+"/publish",
			map[string]any{"published": "true"},
			ts.basicAuthHeader(),
		)
		assert.Equal(t, http.StatusBadRequest, s)
		assertErrorField(t, body, "published")
	})
}

// assertErrorField checks that the error response contains a field error for the given field name.
func assertErrorField(t *testing.T, body []byte, field string) {
	t.Helper()
	var resp map[string]any
	require.NoError(t, json.Unmarshal(body, &resp), "body: %s", body)

	msgs, ok := resp["errorsMessages"]
	require.True(t, ok, "response should have errorsMessages, got: %s", body)

	arr := msgs.([]any)
	require.NotEmpty(t, arr, "errorsMessages should not be empty, got: %s", body)

	for _, item := range arr {
		m := item.(map[string]any)
		if m["field"] == field {
			return
		}
	}
	t.Errorf("expected error for field %q, got: %s", field, body)
}
