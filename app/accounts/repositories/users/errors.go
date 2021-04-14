package repository

import "fmt"

// ErrDuplicateFields duplicate fields
type ErrDuplicateFields struct {
	Field string
}

func (e ErrDuplicateFields) Error() string {
	return fmt.Sprintf("duplicate fields: %s", e.Field)
}
