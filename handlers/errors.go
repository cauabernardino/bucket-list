package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse structures the error response in the API
type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: http.StatusNotFound, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Bad request"}
)

// Render handles inserting status code to error response
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

// ErrorRenderer handles rendering of Bad Request errors
func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: http.StatusBadRequest,
		StatusText: "Bad request",
		Message:    err.Error(),
	}
}

// ServerErrorRenderer handles rendering of Internal Server errors
func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: http.StatusInternalServerError,
		StatusText: "Internal server error",
		Message:    err.Error(),
	}
}
