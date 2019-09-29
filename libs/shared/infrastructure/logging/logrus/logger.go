package logrus

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{}
}

func (logger *LogrusLogger) Log(msg string) {
	logrus.Println(msg)
}
