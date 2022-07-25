package log

import (
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	Sugar = logger.Sugar()
}

func Debug(args ...interface{}) {
	Sugar.Debug(args)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	Sugar.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	Sugar.Info(args)
}

func Infow(msg string, keysAndValues ...interface{}) {
	Sugar.Infow(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	Sugar.Error(args)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	Sugar.Errorw(msg, keysAndValues...)
}
