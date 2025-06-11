package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type messageBody struct {
	Message string `json:"message"`
}

type dataBody struct {
	Data any `json:"data"`
}

type paginatedBody struct {
	Data       any         `json:"data"`
	Pagination pagintation `json:"pagination"`
}

type pagintation struct {
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page,omitempty"`
}

type errorBody struct {
	Error     string `json:"error,omitempty"`
	RequestID string `json:"request_id,omitempty"`
	// Caller    string   `json:"caller,omitempty"`
	// Stack     []string `json:"stacktrace,omitempty"`
}

type response struct {
	ctx    echo.Context
	status int
}

func NewResponse(ctx echo.Context) *response {
	return &response{
		status: http.StatusOK,
		ctx:    ctx,
	}
}

func (r *response) body(data any) error {
	var err error
	if responseJSONPretty {
		err = r.ctx.JSONPretty(
			r.status,
			data,
			responseJSONPrettyIndent,
		)
	} else {
		err = r.ctx.JSON(
			r.status,
			data,
		)
	}

	return err
}

func (r *response) WithStatus(status int) *response {
	r.status = status
	return r
}

func (r *response) Message(message string) error {
	return r.body(
		messageBody{Message: message},
	)
}

func (r *response) Raw(data any) error {
	return r.body(data)
}

func (r *response) Item(item any) error {
	return r.body(
		dataBody{Data: item},
	)
}

func (r *response) ItemCreated(item any) error {
	r.WithStatus(http.StatusCreated)
	return r.body(
		dataBody{Data: item},
	)
}

func (r *response) List(list []any) error {
	return r.body(
		dataBody{Data: list},
	)
}

func (r *response) ListPaginated(list []any, perPage, currentPage, nextPage int) error {
	return r.body(
		paginatedBody{
			Data: list,
			Pagination: pagintation{
				PerPage:     perPage,
				CurrentPage: currentPage,
				NextPage:    nextPage,
			},
		},
	)
}

func (r *response) Error(err error) error {
	er := errorBody{
		Error:     err.Error(),
		RequestID: r.requestID(),
	}

	return r.body(er)
}

func (r *response) requestID() string {
	value := r.ctx.Request().Header.Get(echo.HeaderXRequestID)
	if value == "" {
		value = r.ctx.Response().Header().Get(echo.HeaderXRequestID)
	}

	return value
}
