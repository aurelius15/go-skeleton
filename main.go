package main

import (
	"fmt"

	"github.com/aurelius15/go-skeleton/cmd"
	"github.com/aurelius15/go-skeleton/internal/config"
	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/aurelius15/go-skeleton/internal/reflection"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	configs := config.ParseConfig()
	inter, err := reflection.FirstNotNilInterface(*configs)

	if err != nil {
		log.Default().Panic(err.Error())
	}

	cmdConfig, ok := inter.(config.Configure)
	if !ok {
		log.Default().Panic("wrong command's interface")
	}

	var l *zap.Logger

	switch configs.AppMode {
	case config.LocalMode:
		l, _ = zap.NewDevelopment()
	default:
		l, _ = zap.NewProduction()

		gin.SetMode(gin.ReleaseMode)
	}

	log.SetDefault(l)
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			fmt.Println("zap-logging: " + err.Error())
		}
	}(l)

	if command, ok := cmd.CommandCollection[cmdConfig.Command()]; ok {
		command.BindConfig(cmdConfig)
		command.Execute()
	}
}
