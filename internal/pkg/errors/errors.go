package errors

import (
	"fmt"
	"runtime/debug"
)

type ServiceError struct {
	code       string
	message    string
	err        error
	stacktrace string
}

func (e ServiceError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %s - %v", e.code, e.message, e.err)
	}
	return fmt.Sprintf("%s: %s", e.code, e.message)
}

func (e ServiceError) Code() string {
	return e.code
}

func (e ServiceError) Message() string {
	return e.message
}

func (e ServiceError) Unwrap() error {
	return e.err
}

func (e ServiceError) Stacktrace() string {
	return e.stacktrace
}

// Constructor with optional stack trace
func New(code, message string, err error) ServiceError {
	return ServiceError{
		code:       code,
		message:    message,
		err:        err,
		stacktrace: string(debug.Stack()),
	}
}

func ErrInvalidRequest(err error) ServiceError {
	return New(CodeInvalidRequest, "Invalid request", err)
}

func ErrMissingMandatoryField(field string, err error) ServiceError {
	return New(CodeMissingMandatory, fmt.Sprintf("Missing mandatory field: %s", field), err)
}

func ErrInvalidFieldFormat(field string, err error) ServiceError {
	return New(CodeInvalidFormat, fmt.Sprintf("Invalid format for field: %s", field), err)
}

func ErrDuplicateRequest(reason string) ServiceError {
	return New(CodeDuplicateRequest, fmt.Sprintf("Duplicate request: %s", reason), nil)
}

func ErrNotFound(entity string, err error) ServiceError {
	return New(CodeNotFound, fmt.Sprintf("%s not found", entity), err)
}
