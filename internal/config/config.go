package config

import (
	"github.com/alexflint/go-arg"
)

const (
	LocalMode = "local"
	ProdMode  = "prod"
)

type Configure interface {
	Config(fieldName string) string
	Command() string
}

type Config struct {
	ServerConfig        *ServerConfig        `arg:"subcommand:server"`
	NumberOfUsersConfig *NumberOfUsersConfig `arg:"subcommand:number-users"`
	AppMode             string               `arg:"-m,--app-mode,env:APP_MODE" default:"prod" help:"application mode [local,prod]"`
	RedisPort           string               `arg:"--redis-port,required,env:REDIS_PORT" placeholder:":6379"`
}

func ParseConfig() *Config {
	c := Config{}
	p := arg.MustParse(&c)

	if p.Subcommand() == nil {
		p.Fail("missing subcommand")
	}

	return &c
}
