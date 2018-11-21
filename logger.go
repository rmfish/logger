package gorequest

import (
	"go.uber.org/zap"
	"sync"
)

var once = sync.Once{}
var logger *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if logger != nil {
		return logger
	} else {
		once.Do(func() {
			log, _ := zap.NewDevelopment()
			logger = log.Sugar()
		})
		return logger
	}
}
