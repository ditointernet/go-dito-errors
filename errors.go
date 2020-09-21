package errors

import (
	"net/http"
)

// ErrorKind is a code that reflects HTTP's response code semantics.
// It is supposed to be bypassed directly to the client as an HTTP code response in an HTTP context.
// Althought, its semantic should be respected in any kind of service context (e.g.: gRPC).
type ErrorKind int

const (
	// KindBadRequest reflects the semantic of 400 HTTP status code response.
	KindBadRequest ErrorKind = http.StatusBadRequest
	// KindNotFound reflects the semantic of 404 HTTP status code response.
	KindNotFound ErrorKind = http.StatusNotFound
	// KindUnexpected reflects the semantic of 500 HTTP status code response.
	KindUnexpected ErrorKind = http.StatusInternalServerError
)

// ErrorSeverity indicates how critic this error is considered by the system.
type ErrorSeverity int

const (
	// SeverityCritical indicates a critical error and, most of the time, needs an human intervantion.
	// Normally is related to an alert.
	SeverityCritical ErrorSeverity = 5
	// SeverityError indicates a commun error.
	SeverityError ErrorSeverity = 4
	// SeverityWarning indicates an error that is not that important and should not impact the flow of the request.
	SeverityWarning ErrorSeverity = 3
)

// Error is a representation of a domain error.
type Error struct {
	Kind     ErrorKind     `json:"kind"`
	Severity ErrorSeverity `json:"severity"`
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
		case ErrorKind:
			err.Kind = arg
		case ErrorSeverity:
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

// Severity is a method that returns the severity a given error.
// If not a domain error, it returns a SeverityError as default.
func Severity(err error) ErrorSeverity {
	e, ok := err.(*Error)
	if !ok {
		return SeverityError
	}

	return e.Severity
}

// Kind is a method that returns the kind a given error.
// If not a domain error, it returns a KindUnexpected as default.
func Kind(err error) ErrorKind {
	e, ok := err.(*Error)
	if !ok {
		return KindUnexpected
	}

	return e.Kind
}
