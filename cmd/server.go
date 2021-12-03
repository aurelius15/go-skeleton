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
	switch v := i.(type) {
	case *config.ServerConfig:
		c.config = v
	default:
		log.Default().Panic("can not convert interface to struct")
	}
}

func (c *ServerCmd) Execute() {
	s := server.CreateServer(c.config.Port)
	s.Up()
}

func init() {
	CommandCollection[ServerCmdName] = &ServerCmd{}
}
