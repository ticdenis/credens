package fmt

import (
	"fmt"
)

type FmtLogger struct {
}

func NewFmtLogger() *FmtLogger {
	return &FmtLogger{}
}

func (logger *FmtLogger) Log(msg string) {
	fmt.Printf(msg)
}
