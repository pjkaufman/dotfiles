package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func MustGetCommandOutput(programName, errorMsg string, args ...string) string {
	cmd := exec.Command(programName, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		WriteError(fmt.Sprintf(`%s: %s`, errorMsg, err))
	}

	return string(output)
}

func MustRunCommand(programName, errorMsg string, args ...string) {
	cmd := exec.Command(programName, args...)
	err := cmd.Run()
	if err != nil {
		WriteError(fmt.Sprintf(`%s: %s`, errorMsg, err))
	}
}

func MustChangeDirectoryTo(path string) {
	err := os.Chdir(path)

	if err != nil {
		WriteError(fmt.Sprintf(`failed to change directory to "%s": %s`, path, err))
	}
}

func GetCurrentDirector() (string, error) {
	return os.Getwd()
}
