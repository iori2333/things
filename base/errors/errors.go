package errors

import "fmt"

type Status int

const (
	StatusInvalid      Status = 400
	StatusUnauthorized Status = 401
	StatusForbidden    Status = 403
	StatusNotFound     Status = 404
	StatusConflict     Status = 409
	StatusDefault      Status = 500
)

type Error struct {
	Status    Status `json:"status"`
	ErrorType string `json:"error"`
	Message   string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorType, e.Message)
}

func New(status Status, errType, format string, args ...any) error {
	return &Error{
		Status:    status,
		ErrorType: errType,
		Message:   fmt.Sprintf(format, args...),
	}
}

func Invalid(errType, format string, args ...any) error {
	return New(StatusInvalid, errType, format, args...)
}

func Unauthorized(errType, format string, args ...any) error {
	return New(StatusUnauthorized, errType, format, args...)
}

func Forbidden(errType, format string, args ...any) error {
	return New(StatusForbidden, errType, format, args...)
}

func NotFound(errType, format string, args ...any) error {
	return New(StatusNotFound, errType, format, args...)
}

func Conflict(errType, format string, args ...any) error {
	return New(StatusConflict, errType, format, args...)
}

func Default(errType, format string, args ...any) error {
	return New(StatusDefault, errType, format, args...)
}

func Wrap(status Status, errType string, err error) error {
	if err == nil {
		return nil
	}
	return &Error{
		Status:    status,
		ErrorType: errType,
		Message:   err.Error(),
	}
}
