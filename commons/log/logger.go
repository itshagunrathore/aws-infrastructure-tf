package log

import (
	"go.uber.org/zap"
	nativeLogger "log"
)

var log *zap.SugaredLogger

func InitiateLogger(logLevel string, stage string) {
	isDev := stage == "dev"
	nativeLogger.Println("\t\tINFO \tconfiguring logging with level", logLevel)
	log = Init(logLevel, isDev)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	log.Debugw(msg, keysAndValues...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	log.Errorw(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	log.Warnw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	log.Infow(msg, keysAndValues...)
}

func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args)
}

func Panic(args ...interface{}) {
	log.Panic(args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args)
}
