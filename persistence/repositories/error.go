package repositories

import "fmt"

type PersistenceError struct {
	Message string
	Code    ErrorCode
	err     error
}

// Error implements the error interface for CustomError.
func (e *PersistenceError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Unwrap returns the wrapped error.
func (e *PersistenceError) Unwrap() error {
	return e.err
}

// NewPersistenceError creates a new PersistenceError with the given code and message.
func NewPersistenceError(code ErrorCode, message string, err error) *PersistenceError {
	return &PersistenceError{
		Code:    code,
		Message: message,
		err:     err,
	}
}

// ErrorCode is a custom type for error codes.
type ErrorCode int

const (
	ErrNotFound     ErrorCode = 1
	ErrDuplicateRow ErrorCode = 2
	ErrInternal     ErrorCode = 3
)
