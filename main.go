package main

import (
	"github.com/aurelius15/go-skeleton/cmd"
	"github.com/aurelius15/go-skeleton/internal/config"
	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/aurelius15/go-skeleton/internal/reflection"
)

func main() {
	configs := config.ParseConfig()

	logger := log.NewLogger(configs.AppMode == config.ProdMode)
	log.SetDefault(logger)

	defer log.GracefulSync(logger)

	inter, err := reflection.FirstNotNilInterface(*configs)
	if err != nil {
		log.Default().Panic(err.Error())
	}

	cmdConfig, ok := inter.(config.Configure)
	if !ok {
		log.Default().Panic("wrong command's interface")
	}

	if command, ok := cmd.CommandCollection[cmdConfig.Command()]; ok {
		command.BindConfig(cmdConfig)
		command.Execute()
	}
}
