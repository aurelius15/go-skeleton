package log

import (
	"fmt"

	"go.uber.org/zap"
)

var log = zap.L()

func Default() *zap.Logger {
	return log
}

func SetDefault(l *zap.Logger) {
	log = l
}

func NewLogger(isProd bool) (l *zap.Logger) {
	if isProd {
		l, _ = zap.NewProduction()
	} else {
		l, _ = zap.NewDevelopment()
	}

	return
}

func GracefulSync(l *zap.Logger) {
	err := l.Sync()
	if err != nil {
		fmt.Println("zap-logging: " + err.Error())
	}
}
