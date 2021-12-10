package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateServer(t *testing.T) {
	s := NewServer()

	assert.NotNil(t, s)
}
