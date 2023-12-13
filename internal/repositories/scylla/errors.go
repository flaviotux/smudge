package scylla

import (
	"errors"
	"fmt"
)

type (
	ValidationError struct {
		Name string // Field or edge name.
		err  error
	}
)

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return e.err.Error()
}

// Unwrap implements the errors.Wrapper interface.
func (e *ValidationError) Unwrap() error {
	return e.err
}

// IsValidationError returns a boolean indicating whether the error is a validation error.
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	var e *ValidationError
	return errors.As(err, &e)
}

type (
	// NotFoundError returns when trying to update an
	// entity, and it was not found in the database.
	NotFoundError struct {
		table string
		id    string
	}
)

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("record with id %v not found in table %s", e.id, e.table)
}
