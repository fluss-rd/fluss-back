package router

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetResourceFromURL(t *testing.T) {
	c := require.New(t)

	resource, err := getResourceFromEndpoint(Endpoint{Path: "/modules"})
	c.Nil(err)
	c.Equal("modules", resource)

	resource, err = getResourceFromEndpoint(Endpoint{Path: "/something/modules"})
	c.Nil(err)
	c.Equal("modules", resource)

	resource, err = getResourceFromEndpoint(Endpoint{Path: "/modules/{id}"})
	c.Nil(err)
	c.Equal("modules", resource)

	resource, err = getResourceFromEndpoint(Endpoint{Path: "/rivers/{id}/modules/{id}"})
	c.Nil(err)
	c.Equal("modules", resource)
}
