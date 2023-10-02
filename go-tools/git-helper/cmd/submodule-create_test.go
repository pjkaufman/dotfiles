//go:build unit

package cmd_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/git-helper/cmd"
	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/stretchr/testify/assert"
)

type CreateSubmoduleTestCase struct {
	InputTicket         string
	InputBranchName     string
	InputRepoParentPath string
	ExistingFiles       map[string]struct{}
	ExistingFolders     map[string]struct{}
	PathToFolders       map[string][]string
	CmdErr              error
	ExpectedError       string
	ExpectPanic         bool
}

// errors that get handled as errors are represented as panics
var CreateSubmoduleTestCases = map[string]CreateSubmoduleTestCase{
	"make sure that an empty ticket name causes a validation error": {
		InputTicket:         "",
		InputBranchName:     "name",
		InputRepoParentPath: "/users/username/home/",
		ExpectedError:       cmd.TicketArgEmpty,
		ExpectPanic:         true,
	},
	"make sure that an empty branch name causes a validation error": {
		InputTicket:         "ticket",
		InputBranchName:     "",
		InputRepoParentPath: "/users/username/home/",
		ExpectedError:       cmd.BranchNameArgEmpty,
		ExpectPanic:         true,
	},
	"make sure that an empty repo parent directory causes a validation error": {
		InputTicket:         "ticket",
		InputBranchName:     "name",
		InputRepoParentPath: "/users/username/home/",
		ExpectedError:       cmd.RepoParentPathArgEmpty,
		ExpectPanic:         true,
	},
	// "make sure that when there is an error the error is the expected error message (i.e. the panic value)": {
	// 	CmdErr:        errors.New("exit status 1"),
	// 	ExpectedError: "failed to update the submodule for the current repo: exit status 1",
	// 	ExpectPanic:   true,
	// },
}

func TestCreateSubmodule(t *testing.T) {
	for name, args := range CreateSubmoduleTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleCreateSubmodulePanic(t, args)

			var log = logger.NewMockLoggerHandler()
			var cmdHandler = commandhandler.NewMockCommandHandler(log)
			var fileHandler = filehandler.NewMockFileHandler(log, nil, nil, nil, nil)
			cmd.CreateSubmoduleBranches(log, cmdHandler, fileHandler, args.InputTicket, args.InputBranchName, args.InputRepoParentPath)
		})
	}
}

func handleCreateSubmodulePanic(t *testing.T, args CreateSubmoduleTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
