package services

import (
	"encoding/json"
	"fmt"
	"log/slog"
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

func (s ServiceError) ToJSON() []byte {
	bytes, err := json.Marshal(s)
	if err != nil {
		slog.Error("failed to marshal json", "error", err)
		return nil
	}

	return bytes
}

func (s ServiceError) StatusCode() int {
	return ConvertErrorCodeToHTTPCode(s.Code)
}

func (s ServiceError) DetailMsg() string {
	return s.Message
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
	ErrMissingExternalID           ErrorCode = 4
	ErrMissingUsername             ErrorCode = 5
	ErrDuplicateUser               ErrorCode = 6
	ErrExternalIDMissing           ErrorCode = 7
	ErrNotFound                    ErrorCode = 8
	ErrUserNotInContext            ErrorCode = 9
	ErrDuplicateShopName           ErrorCode = 10
	ErrMissingParameter            ErrorCode = 11
	ErrUserNotConnectedToShop      ErrorCode = 12
	ErrInactiveCurrency            ErrorCode = 13
)

// nolint:gochecknoglobals
var errorCodeHTTPMap = map[ErrorCode]int{
	ErrInvalidToken:           http.StatusUnauthorized,
	ErrMissingExternalID:      http.StatusBadRequest,
	ErrMissingUsername:        http.StatusBadRequest,
	ErrNotFound:               http.StatusNotFound,
	ErrExternalIDMissing:      http.StatusConflict,
	ErrUserNotInContext:       http.StatusUnauthorized,
	ErrDuplicateShopName:      http.StatusBadRequest,
	ErrMissingParameter:       http.StatusBadRequest,
	ErrUserNotConnectedToShop: http.StatusForbidden,
	ErrInactiveCurrency:       http.StatusBadRequest,
}

func ConvertErrorCodeToHTTPCode(code ErrorCode) int {
	if val, ok := errorCodeHTTPMap[code]; ok {
		return val
	}
	return http.StatusInternalServerError
}
