package errors

import (
	"fmt"
)

type ServiceError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *ServiceError) Unwrap() error {
	return e.Err
}

func NewServiceError(code, message string, err error) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
