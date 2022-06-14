package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

const (
	DebugLevelStr   string = "DEBUG"
	InfoLevelStr    string = "INFO"
	WarningLevelStr string = "WARNING"
	ErrorLevelStr   string = "ERROR"
)

var (
	Logger zap.Logger
)

func Init(logLevel string, isDev bool) *zap.SugaredLogger {
	encoderConfig := GetEncoding()
	loggerConfig := zap.Config{
		Level:         zap.NewAtomicLevelAt(levelToZapLevel(logLevel)),
		Development:   isDev,
		Encoding:      "json",
		EncoderConfig: encoderConfig,
		OutputPaths:   []string{"stdout"},
	}

	globalLogger, err := loggerConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("build zap logger from config error: %v", err))
	}
	globalLogger = globalLogger.WithOptions(zap.AddCallerSkip(1))
	_ = zap.ReplaceGlobals(globalLogger)
	_ = globalLogger.Sync()

	Logger := globalLogger.Sugar()
	fmt.Print("logger initialized")
	return Logger
}

func GetEncoding() zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return encoderConfig
}

func levelToZapLevel(s string) zapcore.Level {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case DebugLevelStr:
		fmt.Print("return form here")
		return zapcore.DebugLevel
	case InfoLevelStr:
		return zapcore.InfoLevel
	case WarningLevelStr:
		return zapcore.WarnLevel
	case ErrorLevelStr:
		return zapcore.ErrorLevel
	}
	fmt.Print("retrun")
	return zapcore.InfoLevel
}
