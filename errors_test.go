package errors_test

import (
	e "errors"
	"testing"

	errors "github.com/ditointernet/go-dito-errors"
)

func TestSeverity(t *testing.T) {
	tt := []struct {
		name     string
		err      error
		expected errors.ErrorSeverity
	}{
		{
			name:     "it should return a default severity if the error isn't a domain error",
			err:      e.New("mock error"),
			expected: errors.SeverityError,
		},

		{
			name:     "it should return the error severity",
			err:      errors.New("mock error", errors.SeverityCritical),
			expected: errors.SeverityCritical,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := errors.Severity(tc.err)

			if res != tc.expected {
				t.Errorf("it should return '%d', got '%d'", tc.expected, res)
			}
		})
	}
}

func TestKind(t *testing.T) {
	tt := []struct {
		name     string
		err      error
		expected errors.ErrorKind
	}{
		{
			name:     "it should return a default kind if the error isn't a domain error",
			err:      e.New("mock error"),
			expected: errors.KindUnexpected,
		},

		{
			name:     "it should return the error kind",
			err:      errors.New("mock error", errors.KindBadRequest),
			expected: errors.KindBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := errors.Kind(tc.err)

			if res != tc.expected {
				t.Errorf("it should return '%d', got '%d'", tc.expected, res)
			}
		})
	}
}
