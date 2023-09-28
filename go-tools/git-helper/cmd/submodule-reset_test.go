//go:build unit

package cmd_test

import (
	"errors"
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/git-helper/cmd"
	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/stretchr/testify/assert"
)

type ResetSubmoduleTestCase struct {
	CmdErr        error
	ExpectedError string
	ExpectPanic   bool
}

var ResetSubmoduleTestCases = map[string]ResetSubmoduleTestCase{
	"make sure that when there is no error there is no error message that gets written (i.e. a panic)": {
		ExpectPanic: false,
	},
	"make sure that when there is an error the error is the expected error message (i.e. the panic value)": {
		CmdErr:        errors.New("exit status 1"),
		ExpectedError: "failed to update the submodule for the current repo: exit status 1",
		ExpectPanic:   true,
	},
}

func TestResetSubmodule(t *testing.T) {
	for name, args := range ResetSubmoduleTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleResetSubmodulePanic(t, args)

			var cmdHandler = commandhandler.NewMockCommandHandler(logger.NewMockLoggerHandler())
			cmd.ResetSubmodule(cmdHandler)
		})
	}
}

func handleResetSubmodulePanic(t *testing.T, args ResetSubmoduleTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
