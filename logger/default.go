package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var mainLogger *logrus.Logger

func InitLogger(level logrus.Level) {
	mainLogger = createLogger(level)
}

func Log() *logrus.Logger {
	if mainLogger == nil {
		panic("error on Log - need setup logger before using")
	}

	return mainLogger
}

func createLogger(level logrus.Level) *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stdout,
		Level:     level,
		Formatter: &logrus.JSONFormatter{},
	}
}
