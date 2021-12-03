package log

import "go.uber.org/zap"

var log, _ = zap.NewProduction()

func Default() *zap.Logger {
	return log
}

func SetDefault(l *zap.Logger) {
	log = l
}
