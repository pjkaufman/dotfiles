package utils

import (
	"os"
	"strings"
)

func WriteError(errorMsg string) {
	if !strings.HasSuffix(errorMsg, "\n") {
		errorMsg += "\n"
	}

	os.Stderr.WriteString(errorMsg)
}

func WriteOut(msg string) {
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}

	os.Stdout.WriteString(msg)
}
