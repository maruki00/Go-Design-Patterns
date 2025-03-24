package more_test

import (
	"more"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer_Default(t *testing.T) {
	s, err := more.NewServer()
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "", s.Host)
	assert.Equal(t, 0, s.Port)
}

func TestNewServer_WithPort(t *testing.T) {
	s, err := more.NewServer(more.WithPort(8080))
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, 8080, s.Port)
}

func TestNewServer_WithHost(t *testing.T) {
	s, err := more.NewServer(more.WithHost("localhost"))
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "localhost", s.Host)
}

func TestNewServer_WithHostAndPort(t *testing.T) {
	s, err := more.NewServer(
		more.WithHost("127.0.0.1"),
		more.WithPort(9090),
	)
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "127.0.0.1", s.Host)
	assert.Equal(t, 9090, s.Port)
}
