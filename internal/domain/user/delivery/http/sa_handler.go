package http

import (
	"errors"
	"fmt"
	nethttp "net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/muzyk0/online-quiz-game/internal/platform/http/apperrors"

	userservice "github.com/muzyk0/online-quiz-game/internal/domain/user/service"
)

// SAHandler handles super-admin HTTP requests for user management.
type SAHandler struct {
	svc userservice.UserServiceInterface
}

func NewSAHandler(svc userservice.UserServiceInterface) *SAHandler {
	return &SAHandler{svc: svc}
}

// SAUserViewModel is the response body for SA user endpoints.
type SAUserViewModel struct {
	ID        string    `json:"id"`
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

// SAUserInputModel is the request body for SA create user.
type SAUserInputModel struct {
	Login    string `json:"login"    validate:"required,min=3,max=10,login_pattern"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Email    string `json:"email"    validate:"required,email"`
}

// SAPaginatedUsersResponse is the paginated list of users.
type SAPaginatedUsersResponse struct {
	PagesCount int               `json:"pagesCount"`
	Page       int               `json:"page"`
	PageSize   int               `json:"pageSize"`
	TotalCount int               `json:"totalCount"`
	Items      []SAUserViewModel `json:"items"`
}

// SAGetUsers godoc
//
//	@Summary		Returns all users (SA)
//	@Tags			SA/Users
//	@Produce		json
//	@Param			sortBy			query	string	false	"createdAt | login | email (default: createdAt)"
//	@Param			sortDirection	query	string	false	"asc | desc (default: desc)"
//	@Param			pageNumber		query	int		false	"Page number (default 1)"
//	@Param			pageSize		query	int		false	"Page size 1-20 (default 10)"
//	@Param			searchLoginTerm	query	string	false	"Search by login (contains)"
//	@Param			searchEmailTerm	query	string	false	"Search by email (contains)"
//	@Success		200	{object}	SAPaginatedUsersResponse
//	@Failure		401	{object}	SAErrorResponse
//	@Security		BasicAuth
//	@Router			/api/sa/users [get]
func (h *SAHandler) GetUsers(c echo.Context) error {
	pageSize := queryIntOrDefault(c, "pageSize", 10)
	if pageSize > 20 {
		pageSize = 20
	}
	input := userservice.SAListUsersInput{
		SearchLoginTerm: c.QueryParam("searchLoginTerm"),
		SearchEmailTerm: c.QueryParam("searchEmailTerm"),
		SortBy:          queryParamOrDefault(c, "sortBy", "createdAt"),
		SortDirection:   queryParamOrDefault(c, "sortDirection", "desc"),
		PageNumber:      queryIntOrDefault(c, "pageNumber", 1),
		PageSize:        pageSize,
	}

	result, err := h.svc.SAListUsers(c.Request().Context(), input)
	if err != nil {
		return apperrors.Internal("failed to list users").Wrap(err)
	}

	items := make([]SAUserViewModel, len(result.Items))
	for i, u := range result.Items {
		items[i] = toSAUserViewModel(u)
	}

	return c.JSON(nethttp.StatusOK, SAPaginatedUsersResponse{
		PagesCount: result.PagesCount,
		Page:       result.Page,
		PageSize:   result.PageSize,
		TotalCount: result.TotalCount,
		Items:      items,
	})
}

// SACreateUser godoc
//
//	@Summary		Add new user to the system (SA)
//	@Tags			SA/Users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		SAUserInputModel	true	"User data"
//	@Success		201		{object}	SAUserViewModel
//	@Failure		400		{object}	SAErrorResponse
//	@Failure		401		{object}	SAErrorResponse
//	@Security		BasicAuth
//	@Router			/api/sa/users [post]
func (h *SAHandler) CreateUser(c echo.Context) error {
	var req SAUserInputModel
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(nethttp.StatusBadRequest, "invalid request body")
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	result, err := h.svc.SACreateUser(c.Request().Context(), userservice.SACreateUserInput{
		Login:    req.Login,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return mapSAUserError(err)
	}
	return c.JSON(nethttp.StatusCreated, toSAUserViewModel(result))
}

// SADeleteUser godoc
//
//	@Summary		Delete user by ID (SA)
//	@Tags			SA/Users
//	@Param			id	path	string	true	"User ID"
//	@Success		204
//	@Failure		401	{object}	SAErrorResponse
//	@Failure		404	{object}	SAErrorResponse
//	@Security		BasicAuth
//	@Router			/api/sa/users/{id} [delete]
func (h *SAHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if err := h.svc.SADeleteUser(c.Request().Context(), id); err != nil {
		return mapSAUserError(err)
	}
	return c.NoContent(nethttp.StatusNoContent)
}

// SAErrorResponse is a generic error body for SA endpoints.
type SAErrorResponse struct {
	ErrorsMessages []SAFieldError `json:"errorsMessages"`
}

// SAFieldError describes a single validation error.
type SAFieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

func mapSAUserError(err error) error {
	switch {
	case errors.Is(err, userservice.ErrUserNotFound):
		return apperrors.NotFound("user not found")
	case errors.Is(err, userservice.ErrUserAlreadyExists):
		return apperrors.BadRequest(fmt.Sprintf("user with this email already exists"))
	case errors.Is(err, userservice.ErrLoginAlreadyExists):
		return apperrors.BadRequest("user with this login already exists")
	default:
		return apperrors.Internal("failed to process request").Wrap(err)
	}
}

func queryParamOrDefault(c echo.Context, key, def string) string {
	if v := c.QueryParam(key); v != "" {
		return v
	}
	return def
}

func queryIntOrDefault(c echo.Context, key string, def int) int {
	v := c.QueryParam(key)
	if v == "" {
		return def
	}
	n := 0
	for _, ch := range v {
		if ch < '0' || ch > '9' {
			return def
		}
		n = n*10 + int(ch-'0')
	}
	return n
}

func toSAUserViewModel(u *userservice.UserOutput) SAUserViewModel {
	vm := SAUserViewModel{
		ID:    u.ID,
		Login: u.Login,
		Email: u.Email,
	}
	if u.CreatedAt.Valid {
		vm.CreatedAt = u.CreatedAt.Time
	}
	return vm
}
