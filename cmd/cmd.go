package cmd

import "github.com/aurelius15/go-skeleton/internal/config"

var CommandCollection = make(map[string]Command)

type Command interface {
	BindConfig(configure config.Configure)
	Execute()
}
