package config

import (
	"github.com/aurelius15/go-skeleton/internal/reflection"
)

const ServerCmd = "server"

type ServerConfig struct {
	Port string `arg:"-p,--port,required" placeholder:":8080"`
}

func (c *ServerConfig) Config(fieldName string) (string, error) {
	return reflection.StringFieldByName(c, fieldName)
}

func (c *ServerConfig) Command() string {
	return ServerCmd
}
