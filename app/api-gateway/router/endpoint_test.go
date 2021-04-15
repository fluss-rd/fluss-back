package router

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetResourceFromURL(t *testing.T) {
	c := require.New(t)

	url, err := url.Parse("/modules")
	c.Nil(err)

	resource, err := getResourceFromURL(*url)
	c.Nil(err)
	c.Equal("modules", resource)

	url, err = url.Parse("/something/modules")
	c.Nil(err)

	resource, err = getResourceFromURL(*url)
	c.Nil(err)
	c.Equal("modules", resource)
}
