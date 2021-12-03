package cmd

import (
	"github.com/aurelius15/go-skeleton/internal/config"
	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/aurelius15/go-skeleton/internal/server"
)

const ServerCmdName = "server"

type ServerCmd struct {
	config *config.ServerConfig
}

func (c *ServerCmd) BindConfig(i interface{}) {
	c.config = i.(*config.ServerConfig)
}

func (c *ServerCmd) Execute() {
	log.Default().Info("Server up")
	server.Up(c.config.Port)
}

func init() {
	CommandCollection[ServerCmdName] = &ServerCmd{}
}
