package http

import (
	"strings"
)

// NewError creates a new Error instance.
func NewError(status Status, message ...string) IError {
	if len(message) > 1 {
		return &Error{status, strings.Join(message, " ")}
	}

	return &Error{Status: status, Message: message[0]}
}

// StatusCode returns the HTTP status code.
func (e *Error) StatusCode() Status {
	return e.Status
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}
