package cmd

import (
	"github.com/aurelius15/go-skeleton/internal/config"
	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/aurelius15/go-skeleton/internal/server"
	"go.uber.org/zap"
)

type ServerCmd struct {
	webEngine server.WebEngine
	config    config.Configure
}

func (c *ServerCmd) BindConfig(i config.Configure) {
	c.config = i
	c.webEngine = server.NewServer()
}

func (c *ServerCmd) Execute() {
	if c.config == nil || c.webEngine == nil {
		log.Default().Fatal("before Execute trigger BindConfig")
	}

	port := c.config.Config("Port")

	if err := c.webEngine.Run(port); err != nil {
		log.Default().Fatal("Error during running server", zap.Error(err))
	}
}

func init() {
	CommandCollection[config.ServerCmd] = &ServerCmd{}
}
