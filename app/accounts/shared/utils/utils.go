package utils

import (
	"strings"

	"github.com/gofrs/uuid"
)

// GenerateID generates a uuid with a prefix
func GenerateID(prefix string) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", nil
	}

	return prefix + strings.ReplaceAll(id.String(), "-", ""), nil
}
