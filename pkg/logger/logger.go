package logger

import (
	"sync"

	"go.uber.org/zap"
)

var logOnce sync.Once
var log *zap.SugaredLogger

func Log() *zap.SugaredLogger {
	if log == nil {
		logOnce.Do(
			func() {
				raw, err := zap.NewProduction()
				if err != nil {
					panic(err)
				}
				log = raw.Sugar()
			},
		)
	}

	return log
}
