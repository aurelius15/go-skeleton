package cmd

import (
	"github.com/aurelius15/go-skeleton/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerCmd_Init(t *testing.T) {
	cmd := CommandCollection[ServerCmdName]

	assert.NotNil(t, cmd)
	assert.Equal(t, &ServerCmd{}, cmd)
}

func TestServerCmd_BindConfig(t *testing.T) {
	cmd := CommandCollection[ServerCmdName]

	assert.NotPanicsf(t, func() {
		cmd.BindConfig(&config.ServerConfig{})
	}, "It can't bind only &config.ServerConfig")
	assert.Panicsf(t, func() {
		cmd.BindConfig(nil)
	}, "It can bind only &config.ServerConfig")
}
