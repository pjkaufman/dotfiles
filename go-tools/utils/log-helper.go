package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// type Logger interface {
// }

func WriteError(errorMsg string) {
	color.New(color.FgRed).Fprintln(os.Stderr, errorMsg)
	os.Exit(-1)
}

func WriteInfo(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}

func WriteWarn(msg string) {
	color.New(color.FgYellow).Fprintln(os.Stdout, msg)
}
