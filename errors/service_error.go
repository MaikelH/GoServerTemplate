package errors

import (
	"fmt"
	"net/http"
)

type ServiceError struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`
	Err     error     `json:"-"`
}

func (s ServiceError) Error() string {
	return fmt.Sprintf("Service Error: %s", s.Message)
}

func (s ServiceError) Unwrap() error {
	return s.Err
}

func NewServiceError(code ErrorCode, err error, message string) ServiceError {
	return ServiceError{
		Message: message,
		Err:     err,
		Code:    code,
	}
}

type ErrorCode int

const (
	ErrServerError                 ErrorCode = 1
	ErrMissingAuthenticationHeader ErrorCode = 2
	ErrInvalidToken                ErrorCode = 3
)

var ErrorCodeHTTPMap = map[ErrorCode]int{
	ErrInvalidToken: http.StatusUnauthorized,
}

func ConvertErrorCodeToHTTPCode(code ErrorCode) int {
	if val, ok := ErrorCodeHTTPMap[code]; ok {
		return val
	}
	return http.StatusInternalServerError
}
