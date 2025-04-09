package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func Init() {
	raw, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	Log = raw.Sugar()
}
