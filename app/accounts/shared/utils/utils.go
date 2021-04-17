package utils

import (
	"encoding/json"
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

// Converts a struct to a map while maintaining the json alias as keys
func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}
