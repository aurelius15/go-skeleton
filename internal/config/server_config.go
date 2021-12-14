package config

import (
	"github.com/aurelius15/go-skeleton/internal/reflection"
)

const ServerCmd = "server"

type ServerConfig struct {
	Port string `arg:"-p,--port,required,env:APP_PORT" placeholder:":8080"`
}

func (c *ServerConfig) Config(fieldName string) (s string) {
	s, _ = reflection.StringFieldByName(c, fieldName)
	return
}

func (c *ServerConfig) Command() string {
	return ServerCmd
}
