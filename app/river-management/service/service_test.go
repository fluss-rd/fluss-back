package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidPhoneNumber(t *testing.T) {
	c := require.New(t)

	c.False(isValidPhoneNumber("8097538038"))
	c.True(isValidPhoneNumber("+18097538038"))
}
