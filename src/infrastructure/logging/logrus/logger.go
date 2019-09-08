package logrus

import (
	"credens/src/infrastructure/logging"
	"github.com/sirupsen/logrus"
)

type Logger struct {
}

func (logger *Logger) Log(msg string) {
	logrus.Println(msg)
}

func NewLogger() logging.Logger {
	return new(Logger)
}
