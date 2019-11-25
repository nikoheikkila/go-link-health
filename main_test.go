package main

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func fakeLink(address string, status int) *Link {
	parsedURL, _ := url.Parse(address)

	return &Link{
		url:    parsedURL,
		status: status,
	}
}

func TestLink(t *testing.T) {
	goodLink := fakeLink("http://good.org", 200)
	badLink := fakeLink("http://bad.org", 500)

	assert.Equal(t, "http://good.org", goodLink.url.String())
	assert.Equal(t, "http://bad.org", badLink.url.String())
	assert.Equal(t, 200, goodLink.status)
	assert.Equal(t, 500, badLink.status)
}

func TestIsHealthy(t *testing.T) {
	goodLink := fakeLink("http://good.org", 200)
	badLink := fakeLink("http://bad.org", 300)

	assert.True(t, goodLink.isHealthy())
	assert.False(t, badLink.isHealthy())
}
