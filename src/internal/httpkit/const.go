package httpkit

import "time"

const (
	HeaderContentType = "Content-Type"
	ContentTypeJSON   = "application/json"
	DateTimeLayout    = time.RFC3339

	VerdictFailure           = "failure"
	VerdictMissingParameters = "missing_parameters"
	VerdictSuccess           = "success"
	VerdictInvalidParameters = "invalid_parameters"
	VerdictExistedRecord     = "existed_record"
	VerdictRecordNotFound    = "record_not_found"
)
