package main

import (
	"github.com/aurelius15/go-skeleton/cmd"
	"github.com/aurelius15/go-skeleton/internal/config"
	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/aurelius15/go-skeleton/internal/reflection"
	"github.com/aurelius15/go-skeleton/internal/storage"
	"github.com/go-redis/redis/v8"
)

func main() {
	configs := config.ParseConfig()

	logger := log.NewLogger(configs.AppMode == config.ProdMode)
	log.SetDefault(logger)

	defer log.GracefulSync(logger)

	storage.SetInstance(redis.NewClient(&redis.Options{
		Addr: configs.RedisPort,
	}))

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
