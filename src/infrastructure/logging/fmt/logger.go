package fmt

import (
	"fmt"
)

type Logger struct {
}

func (logger *Logger) Log(msg string) {
	fmt.Printf(msg)
}

func NewLogger() *Logger {
	return &Logger{}
}
