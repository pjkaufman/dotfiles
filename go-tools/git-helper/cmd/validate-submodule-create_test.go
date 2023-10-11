//go:build unit

package cmd_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/git-helper/cmd"
	"github.com/stretchr/testify/assert"
)

type ValidateCreateSubmoduleTestCase struct {
	InputTicket         string
	InputBranchName     string
	InputRepoParentPath string
	ExpectedError       string
}

var ValidateCreateSubmoduleTestCases = map[string]ValidateCreateSubmoduleTestCase{
	"make sure that an empty ticket name causes a validation error": {
		InputTicket:         "",
		InputBranchName:     "name",
		InputRepoParentPath: "/users/username/home/",
		ExpectedError:       cmd.TicketArgEmpty,
	},
	"make sure that an empty branch name causes a validation error": {
		InputTicket:         "ticket",
		InputBranchName:     "",
		InputRepoParentPath: "/users/username/home/",
		ExpectedError:       cmd.BranchNameArgEmpty,
	},
	"make sure that an empty repo parent directory causes a validation error": {
		InputTicket:         "ticket",
		InputBranchName:     "name",
		InputRepoParentPath: "",
		ExpectedError:       cmd.RepoParentPathArgEmpty,
	},
	"make sure that ticket, branch name, and repo parent directory having values passes validation": {
		InputTicket:         "ticket",
		InputBranchName:     "name",
		InputRepoParentPath: "/users/username/home/",
		ExpectedError:       "",
	},
}

func TestValidateCreateSubmodule(t *testing.T) {
	for name, args := range ValidateCreateSubmoduleTestCases {
		t.Run(name, func(t *testing.T) {
			err := cmd.ValidateSubmoduleCreate(args.InputTicket, args.InputBranchName, args.InputRepoParentPath)

			if err != nil {
				assert.Equal(t, args.ExpectedError, err.Error())
			} else {
				assert.Equal(t, args.ExpectedError, "")
			}
		})
	}
}
