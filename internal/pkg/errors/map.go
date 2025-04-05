package errors

import "net/http"

const (
	CodeInvalidRequest   = "INVALID_REQUEST"
	CodeMissingMandatory = "MISSING_FIELD"
	CodeInvalidFormat    = "INVALID_FORMAT"
	CodeDuplicateRequest = "DUPLICATE_REQUEST"
	CodeNotFound         = "NOT_FOUND"
	CodeInternal         = "INTERNAL_ERROR"
)

func HTTPStatusFromCode(code string) int {
	switch code {
	case CodeInvalidRequest,
		CodeMissingMandatory,
		CodeInvalidFormat,
		CodeDuplicateRequest:
		return http.StatusBadRequest
	case CodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
