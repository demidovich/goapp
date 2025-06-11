package errors

import std_errors "errors"

const (
	ErrUnauthorized        = "Unauthorized"
	ErrForbidden           = "Forbidden"
	ErrRequestTimeout      = "Request timeout"
	ErrTooManyRequests     = "Too many requests"
	ErrBadRequest          = "Bad request"
	ErrNotFound            = "Not found"
	ErrLogicError          = "Logic error"
	ErrInternalServerError = "Internal server error"
)

var (
	Unauthorized        = std_errors.New("Unauthorized")
	Forbidden           = std_errors.New("Forbidden")
	RequestTimeout      = std_errors.New("Request timeout")
	TooManyRequests     = std_errors.New("Too many requests")
	BadRequest          = std_errors.New("Bad request")
	NotFound            = std_errors.New("Not found")
	LogicError          = std_errors.New("Logic error")
	InternalServerError = std_errors.New("Internal server error")
)

func As(err error, target any) bool {
	return std_errors.As(err, target)
}

func Is(err, target error) bool {
	return std_errors.Is(err, target)
}

func New(text string) error {
	return std_errors.New(text)
}
