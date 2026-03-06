package http

import (
	"encoding/json"
	"errors"
	"fmt"
	nethttp "net/http"

	"github.com/labstack/echo/v4"

	"github.com/muzyk0/online-quiz-game/internal/domain/question/delivery/http/dto"
	"github.com/muzyk0/online-quiz-game/internal/domain/question/service"
	"github.com/muzyk0/online-quiz-game/internal/pkg/apperrors"
)

// Handler handles SA HTTP requests for quiz questions.
type Handler struct {
	svc service.QuestionServiceInterface
}

func NewHandler(svc service.QuestionServiceInterface) *Handler {
	return &Handler{svc: svc}
}

// GetQuestions godoc
//
//	@Summary		List quiz questions (SA)
//	@Tags			SA/Questions
//	@Produce		json
//	@Param			bodySearchTerm	query		string	false	"Search term in body"
//	@Param			publishedStatus	query		string	false	"all | published | notPublished (default: all)"
//	@Param			sortBy			query		string	false	"createdAt | updatedAt | body (default: createdAt)"
//	@Param			sortDirection	query		string	false	"asc | desc (default: desc)"
//	@Param			pageNumber		query		int		false	"Page number (default 1)"
//	@Param			pageSize		query		int		false	"Page size 1-20 (default 10)"
//	@Success		200	{object}	dto.PaginatedQuestionsResponse
//	@Failure		401	{object}	dto.ErrorResponse
//	@Security		BasicAuth
//	@Router			/api/sa/quiz/questions [get]
func (h *Handler) GetQuestions(c echo.Context) error {
	pageSize := queryIntOrDefault(c, "pageSize", 10)
	if pageSize > 20 {
		pageSize = 20
	}
	input := service.ListQuestionsInput{
		BodySearchTerm:  c.QueryParam("bodySearchTerm"),
		PublishedStatus: queryParamOrDefault(c, "publishedStatus", "all"),
		SortBy:          queryParamOrDefault(c, "sortBy", "createdAt"),
		SortDirection:   queryParamOrDefault(c, "sortDirection", "desc"),
		PageNumber:      queryIntOrDefault(c, "pageNumber", 1),
		PageSize:        pageSize,
	}

	result, err := h.svc.ListQuestions(c.Request().Context(), input)
	if err != nil {
		return mapQuestionError(err)
	}
	return c.JSON(nethttp.StatusOK, dto.FromPaginatedOutput(result))
}

// CreateQuestion godoc
//
//	@Summary		Create a quiz question (SA)
//	@Tags			SA/Questions
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.CreateQuestionRequest	true	"Question data"
//	@Success		201		{object}	dto.QuestionResponse
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Security		BasicAuth
//	@Router			/api/sa/quiz/questions [post]
func (h *Handler) CreateQuestion(c echo.Context) error {
	var req dto.CreateQuestionRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(nethttp.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	result, err := h.svc.CreateQuestion(c.Request().Context(), service.CreateQuestionInput{
		Body:           req.Body,
		CorrectAnswers: req.CorrectAnswers,
	})
	if err != nil {
		return mapQuestionError(err)
	}
	return c.JSON(nethttp.StatusCreated, dto.FromServiceOutput(result))
}

// UpdateQuestion godoc
//
//	@Summary		Update a quiz question (SA)
//	@Tags			SA/Questions
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"Question ID"
//	@Param			body	body		dto.UpdateQuestionRequest	true	"Question data"
//	@Success		204
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Security		BasicAuth
//	@Router			/api/sa/quiz/questions/{id} [put]
func (h *Handler) UpdateQuestion(c echo.Context) error {
	id := c.Param("id")

	var req dto.UpdateQuestionRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(nethttp.StatusBadRequest, "Invalid request body")
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	_, err := h.svc.UpdateQuestion(c.Request().Context(), id, service.UpdateQuestionInput{
		Body:           req.Body,
		CorrectAnswers: req.CorrectAnswers,
	})
	if err != nil {
		return mapQuestionError(err)
	}
	return c.NoContent(nethttp.StatusNoContent)
}

// DeleteQuestion godoc
//
//	@Summary		Delete a quiz question (SA)
//	@Tags			SA/Questions
//	@Param			id	path	string	true	"Question ID"
//	@Success		204
//	@Failure		401	{object}	dto.ErrorResponse
//	@Failure		404	{object}	dto.ErrorResponse
//	@Security		BasicAuth
//	@Router			/api/sa/quiz/questions/{id} [delete]
func (h *Handler) DeleteQuestion(c echo.Context) error {
	id := c.Param("id")

	if err := h.svc.DeleteQuestion(c.Request().Context(), id); err != nil {
		return mapQuestionError(err)
	}
	return c.NoContent(nethttp.StatusNoContent)
}

// PublishQuestion godoc
//
//	@Summary		Publish or unpublish a quiz question (SA)
//	@Tags			SA/Questions
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"Question ID"
//	@Param			body	body		dto.PublishQuestionRequest	true	"Publish flag"
//	@Success		204
//	@Failure		400		{object}	dto.ErrorResponse
//	@Failure		401		{object}	dto.ErrorResponse
//	@Failure		404		{object}	dto.ErrorResponse
//	@Security		BasicAuth
//	@Router			/api/sa/quiz/questions/{id}/publish [put]
func (h *Handler) PublishQuestion(c echo.Context) error {
	id := c.Param("id")

	var req dto.PublishQuestionRequest
	if err := c.Bind(&req); err != nil {
		var typeErr *json.UnmarshalTypeError
		if errors.As(err, &typeErr) {
			field := typeErr.Field
			if field == "" {
				field = "published" // For top-level JSON type errors
			}
			return apperrors.NewValidationError(map[string]string{
				field: "must be a boolean",
			})
		}
		return apperrors.BadRequest("Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	if req.Published == nil {
		return apperrors.NewValidationError(map[string]string{
			"published": "field is required",
		})
	}

	_, err := h.svc.PublishQuestion(c.Request().Context(), id, *req.Published)
	if err != nil {
		return mapQuestionError(err)
	}
	return c.NoContent(nethttp.StatusNoContent)
}

// --- helpers ---

func queryParamOrDefault(c echo.Context, key, def string) string {
	v := c.QueryParam(key)
	if v == "" {
		return def
	}
	return v
}

func queryIntOrDefault(c echo.Context, key string, def int) int {
	v := c.QueryParam(key)
	if v == "" {
		return def
	}
	var n int
	if _, err := fmt.Sscanf(v, "%d", &n); err != nil || n < 1 {
		return def
	}
	return n
}
