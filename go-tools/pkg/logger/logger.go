package logger

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Logger interface {
	WriteInfo(msg string)
	WriteWarn(msg string)
	WriteError(msg string)
	GetInputString(prompt string) string
	GetInputInt(prompt string) int
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

func (lg LoggerHandler) GetInputString(prompt string) string {
	fmt.Println(prompt)
	// based on https://stackoverflow.com/a/20895629 since for some reason spaces were not read properly by fmt.Scanln
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		lg.WriteError(fmt.Sprintf("failed to read in the string provided: %s", err))
	}

	response = strings.TrimRight(response, "\n")

	return response
}

func (lg LoggerHandler) GetInputInt(prompt string) int {
	fmt.Println(prompt)
	var response int
	fmt.Scanf("%d", &response)

	return response
}
