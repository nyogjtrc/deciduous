package apierrors

import "fmt"

// error codes
const (
	_ = uint(iota)
	CodeBadRequest
	CodeBadHeader
	CodeUnauthorized
	CodeDataNotFound
)

// AppError struct with code and message
type AppError struct {
	code    uint
	message string
	err     error
}

// New create Error instance
func New(code uint, message string, err error) *AppError {
	return &AppError{code, message, err}
}

// Error getter
func (e *AppError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.message)
}

// Code getter
func (e *AppError) Code() uint {
	return e.code
}

// Message getter
func (e *AppError) Message() string {
	return e.message
}

// ErrorSummry getter
func (e *AppError) ErrorSummry() string {
	return fmt.Sprintf("%d: %s, [%s]", e.code, e.message, e.err.Error())
}
