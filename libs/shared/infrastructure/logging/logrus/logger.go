package logrus

import (
	logging2 "credens/libs/shared/infrastructure/logging"
	"github.com/sirupsen/logrus"
)

type Logger struct {
}

func (logger *Logger) Log(msg string) {
	logrus.Println(msg)
}

func NewLogger() logging2.Logger {
	return new(Logger)
}
