package log

import (
	"fmt"
	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger

func Logger() *zap.SugaredLogger {

	if sugarLogger == nil {
		fmt.Println("init logger")
		logger, _ := zap.NewProduction()
		sugarLogger = logger.Sugar()
		defer logger.Sync()
	}

	return sugarLogger
}
