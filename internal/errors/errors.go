package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Status  int           `json:"-"`
	Code    string        `json:"code"`
	Title   string        `json:"title"`
	Detail  string        `json:"detail,omitempty"`
	Err     error         `json:"-"`
	Meta    any           `json:"meta,omitempty"`
	Service *ServiceError `json:"service,omitempty"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Title, e.Err)
	}
	return e.Title
}

func (e *AppError) Unwrap() error {
	return e.Err
}

type ErrorResponse struct {
	Errors []AppError `json:"errors"`
}

func NewValidationError(detail string, err error) *AppError {
	return &AppError{
		Status: http.StatusBadRequest,
		Code:   CodeValidationError,
		Title:  TitleValidationError,
		Detail: detail,
		Err:    err,
	}
}

func NewNotFound(detail string, err error) *AppError {
	return &AppError{
		Status: http.StatusNotFound,
		Code:   CodeNotFound,
		Title:  TitleNotFound,
		Detail: detail,
		Err:    err,
	}
}

func NewUnauthorized(detail string, err error) *AppError {
	return &AppError{
		Status: http.StatusUnauthorized,
		Code:   CodeUnauthorized,
		Title:  TitleUnauthorized,
		Detail: detail,
		Err:    err,
	}
}

func NewForbidden(detail string, err error) *AppError {
	return &AppError{
		Status: http.StatusForbidden,
		Code:   CodeForbidden,
		Title:  TitleForbidden,
		Detail: detail,
		Err:    err,
	}
}

func NewBadRequest(detail string, err error) *AppError {
	return &AppError{
		Status: http.StatusBadRequest,
		Code:   CodeBadRequest,
		Title:  TitleBadRequest,
		Detail: detail,
		Err:    err,
	}
}

func NewInternalError(err error) *AppError {
	appErr := &AppError{
		Status: http.StatusInternalServerError,
		Code:   CodeInternalError,
		Title:  TitleInternalError,
		Detail: "An internal error occurred",
		Err:    err,
	}

	if se, ok := err.(*ServiceError); ok {
		appErr.Service = se
		appErr.Detail = se.Message
	}

	return appErr
}

func NewConflict(detail string, err error) *AppError {
	return &AppError{
		Status: http.StatusConflict,
		Code:   CodeConflict,
		Title:  TitleConflict,
		Detail: detail,
		Err:    err,
	}
}

func NewMethodNotAllowed(method string) *AppError {
	return &AppError{
		Status: http.StatusMethodNotAllowed,
		Code:   CodeMethodNotAllowed,
		Title:  TitleMethodNotAllowed,
		Detail: fmt.Sprintf("Method %s is not allowed", method),
	}
}

const (
	CodeValidationError  = "VALIDATION_ERROR"
	CodeNotFound         = "NOT_FOUND"
	CodeUnauthorized     = "UNAUTHORIZED"
	CodeForbidden        = "FORBIDDEN"
	CodeInternalError    = "INTERNAL_ERROR"
	CodeBadRequest       = "BAD_REQUEST"
	CodeConflict         = "CONFLICT"
	CodeTooManyRequests  = "TOO_MANY_REQUESTS"
	CodeMethodNotAllowed = "METHOD_NOT_ALLOWED"
	CodeUnsupportedMedia = "UNSUPPORTED_MEDIA_TYPE"
)

const (
	TitleValidationError  = "Validation Error"
	TitleNotFound         = "Resource Not Found"
	TitleUnauthorized     = "Unauthorized"
	TitleForbidden        = "Forbidden"
	TitleInternalError    = "Internal Server Error"
	TitleBadRequest       = "Bad Request"
	TitleConflict         = "Conflict"
	TitleTooManyRequests  = "Too Many Requests"
	TitleMethodNotAllowed = "Method Not Allowed"
)
