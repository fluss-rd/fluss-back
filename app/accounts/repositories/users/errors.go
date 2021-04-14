package repository

// ErrDuplicateFields duplicate fields
type ErrDuplicateFields struct {
	Field string
}

func (e ErrDuplicateFields) Error() string {
	return "duplicate fields"
}
