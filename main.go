package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aurelius15/go-skeleton/cmd"
	"github.com/aurelius15/go-skeleton/internal/config"
	"github.com/aurelius15/go-skeleton/internal/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	configs := config.ParseConfig()
	cmdName, subConfig := getSubCommand(configs)

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

	if command, ok := cmd.CommandCollection[cmdName]; ok {
		command.BindConfig(subConfig)
		command.Execute()
	}
}

func getSubCommand(c *config.Config) (name string, inter interface{}) {
	v := reflect.ValueOf(*c)

	for i := 0; i < v.NumField(); i++ {
		subCmd := v.Field(i)

		if subCmd.Interface() != nil && subCmd.IsNil() {
			continue
		}

		inter = subCmd.Interface()
		name = strings.Split(strings.SplitN(v.Type().Field(i).Tag.Get("arg"), ",", 1)[0], ":")[1]

		break
	}

	return
}
