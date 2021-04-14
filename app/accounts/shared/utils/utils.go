package utils

import (
	"strings"

	"github.com/gofrs/uuid"
)

func GenerateID(prefix string) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", nil
	}

	return prefix + strings.ReplaceAll(id.String(), "-", ""), nil
}
