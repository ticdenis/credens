package logrus

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
}

func (logger *Logger) Log(msg string) {
	logrus.Println(msg)
}

func NewLogger() *Logger {
	return &Logger{}
}
