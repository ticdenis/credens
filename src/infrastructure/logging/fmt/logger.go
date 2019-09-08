package fmt

import (
	"credens/src/infrastructure/logging"
	"fmt"
)

type Logger struct {
}

func (logger *Logger) Log(msg string) {
	fmt.Printf(msg)
}

func NewLogger() logging.Logger {
	return new(Logger)
}
