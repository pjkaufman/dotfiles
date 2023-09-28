package commandhandler

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

type CommandManager interface {
	MustGetCommandOutput(programName, errorMsg string, args ...string) string
	MustRunCommand(programName, errorMsg string, args ...string)
	MustChangeDirectoryTo(path string)
	GetCurrentDirectory() (string, error)
}

type CommandHandler struct {
	logger logger.Logger
}

func NewCommandHandler(logger logger.Logger) *CommandHandler {
	return &CommandHandler{
		logger: logger,
	}
}

func (ch *CommandHandler) MustGetCommandOutput(programName, errorMsg string, args ...string) string {
	cmd := exec.Command(programName, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		ch.logger.WriteError(fmt.Sprintf(`%s: %s`, errorMsg, err))
	}

	return string(output)
}

func (ch *CommandHandler) MustRunCommand(programName, errorMsg string, args ...string) {
	cmd := exec.Command(programName, args...)
	err := cmd.Run()
	if err != nil {
		ch.logger.WriteError(fmt.Sprintf(`%s: %s`, errorMsg, err))
	}
}

func (ch *CommandHandler) MustChangeDirectoryTo(path string) {
	err := os.Chdir(path)

	if err != nil {
		ch.logger.WriteError(fmt.Sprintf(`failed to change directory to "%s": %s`, path, err))
	}
}

func (ch *CommandHandler) GetCurrentDirectory() (string, error) {
	return os.Getwd()
}
