package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Logger interface {
	WriteInfo(msg string)
	WriteWarn(msg string)
	WriteError(msg string)
}

type LoggerHandler struct{}

func NewLoggerHandler() *LoggerHandler {
	return &LoggerHandler{}
}

func (lg LoggerHandler) WriteError(msg string) {
	color.New(color.FgRed).Fprintln(os.Stderr, msg)
	os.Exit(-1)
}

func (lg LoggerHandler) WriteInfo(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}

func (lg LoggerHandler) WriteWarn(msg string) {
	color.New(color.FgYellow).Fprintln(os.Stdout, msg)
}
