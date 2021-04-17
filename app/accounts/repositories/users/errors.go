package repository

import "errors"

var (
	// ErrNotFound not found
	ErrNotFound = errors.New("not found")
	// ErrNothingToUpdate nothing to update
	ErrNothingToUpdate = errors.New("nothing to update")
	// ErrMissingUserID missing user id
	ErrMissingUserID = errors.New("missing user id")
)

// ErrDuplicateFields duplicate fields
type ErrDuplicateFields struct {
	Field string
}

func (e ErrDuplicateFields) Error() string {
	return "duplicate fields"
}
