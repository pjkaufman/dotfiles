//go:build unit

package commandhandler

import (
	"fmt"

	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

type CommandHandlerMock struct {
	logger    logger.Logger
	CmdErr    error
	CmdOutput string
}

func NewMockCommandHandler(logger logger.Logger) *CommandHandlerMock {
	return &CommandHandlerMock{
		logger: logger,
	}
}

func (ch *CommandHandlerMock) MustGetCommandOutput(programName, errorMsg string, args ...string) string {
	if ch.CmdErr != nil {
		ch.logger.WriteError(fmt.Sprintf(`%s: %s`, errorMsg, ch.CmdErr))
	}

	return ch.CmdOutput
}

func (ch *CommandHandlerMock) MustRunCommand(programName, errorMsg string, args ...string) {
	if ch.CmdErr != nil {
		ch.logger.WriteError(fmt.Sprintf(`%s: %s`, errorMsg, ch.CmdErr))
	}
}

func (ch *CommandHandlerMock) MustChangeDirectoryTo(path string) {
	if ch.CmdErr != nil {
		ch.logger.WriteError(fmt.Sprintf(`failed to change directory to "%s": %s`, path, ch.CmdErr))
	}
}

func (ch *CommandHandlerMock) GetCurrentDirectory() (string, error) {
	return ch.CmdOutput, ch.CmdErr
}
