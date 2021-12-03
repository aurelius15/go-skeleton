package config

import (
	"github.com/alexflint/go-arg"
)

const (
	LocalMode = "local"
	ProdMode  = "prod"
)

type ServerConfig struct {
	Port string `arg:"-p,--port,required"`
}

type Config struct {
	ServerConfig *ServerConfig `arg:"subcommand:server"`
	AppMode      string        `arg:"-m,--app-mode,env:APP_MODE" default:"prod" help:"application mode [local,prod]"`
}

func ParseConfig() *Config {
	c := Config{}
	arg.MustParse(&c)

	return &c
}
