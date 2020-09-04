package errors

import (
	"net/http"
)

// Kind is a code that reflects HTTP's response code semantics.
// It is supposed to be bypassed directly to the client as an HTTP code response in an HTTP context.
// Althought, its semantic should be respected in any kind of service context (e.g.: gRPC).
type Kind int

const (
	// KindBadRequest reflects the semantic of 400 HTTP status code response.
	KindBadRequest Kind = http.StatusBadRequest
	// KindNotFound reflects the semantic of 404 HTTP status code response.
	KindNotFound Kind = http.StatusNotFound
	// KindUnexpected reflects the semantic of 500 HTTP status code response.
	KindUnexpected Kind = http.StatusInternalServerError
)

// Severity indicates how critic this error is considered by the system.
type Severity int

const (
	// SeverityCritical indicates a critical error and, most of the time, needs an human intervantion.
	// Normally is related to an alert.
	SeverityCritical Severity = 5
	// SeverityError indicates a commun error.
	SeverityError Severity = 4
	// SeverityWarning indicates an error that is not that important and should not impact the flow of the request.
	SeverityWarning Severity = 3
)

// Error is a representation of a domain error.
type Error struct {
	Kind     Kind     `json:"kind"`
	Severity Severity `json:"severity"`
	message  string
}

// New is a construction function of a domain error.
func New(args ...interface{}) *Error {
	err := Error{
		Kind:     KindUnexpected,
		Severity: SeverityError,
		message:  "Internal server error",
	}

	for _, arg := range args {
		switch arg := arg.(type) {
		case error:
			err.message = arg.Error()
		case Kind:
			err.Kind = arg
		case Severity:
			err.Severity = arg
		case string:
			err.message = arg
		}
	}

	return &err
}

// Error is a method that returns errors's message.
func (err Error) Error() string {
	return err.message
}
