package errors

import (
	"errors"
	"testing"
)

func TestServiceErrorWrap(t *testing.T) {
	rootErr := errors.New("DB connection failed")
	err := ErrInvalidRequest(rootErr)

	if !errors.Is(err, rootErr) {
		t.Errorf("Expected wrapped error to match root error")
	}

	if err.Code() != CodeInvalidRequest {
		t.Errorf("Expected code %s, got %s", CodeInvalidRequest, err.Code())
	}

	if err.Stacktrace() == "" {
		t.Error("Expected non-empty stacktrace")
	}
}
