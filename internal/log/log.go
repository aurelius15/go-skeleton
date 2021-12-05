package log

import "go.uber.org/zap"

var log = zap.L()

func Default() *zap.Logger {
	return log
}

func SetDefault(l *zap.Logger) {
	log = l
}
