package repository

import "errors"

var (
	// ErrNotFound not found
	ErrNotFound = errors.New("not found")
)

// ErrDuplicateFields duplicate fields
type ErrDuplicateFields struct {
	Field string
}

func (e ErrDuplicateFields) Error() string {
	return "duplicate fields"
}
