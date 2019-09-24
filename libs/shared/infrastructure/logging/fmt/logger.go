package fmt

import (
	logging2 "credens/libs/shared/infrastructure/logging"
	"fmt"
)

type Logger struct {
}

func (logger *Logger) Log(msg string) {
	fmt.Printf(msg)
}

func NewLogger() logging2.Logger {
	return new(Logger)
}
