package httpkit

import "time"

const (
	HeaderContentType = "Content-Type"
	ContentTypeJSON   = "application/json"
	DateTimeLayout    = time.RFC3339

	VerdictMissingAuthorization = "missing_authorization"
	VerdictUnknownAuthorization = "unknown_authorization"

	VerdictFailure           = "failure"
	VerdictMissingParameters = "missing_parameters"
	VerdictSuccess           = "success"
)
