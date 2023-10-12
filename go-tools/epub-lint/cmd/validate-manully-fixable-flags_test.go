//go:build unit

package cmd_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/cmd"
	"github.com/stretchr/testify/assert"
)

type ValidateManuallyFixableFlagsTestCase struct {
	InputFilePaths       []string
	InputCssPaths        []string
	InputRunAll          bool
	InputRunBrokenLines  bool
	InputRunSectionBreak bool
	InputRunPageBreak    bool
	InputRunOxfordCommas bool
	InputRunAlthoughBut  bool
	ExpectedError        string
}

var ValidateManuallyFixableFlagsTestCases = map[string]ValidateManuallyFixableFlagsTestCase{
	"make sure that all bool flags being false causes a validation error": {
		ExpectedError: cmd.OneRunBoolArgMustBeEnabled,
	},
	"make sure that an empty file paths causes a validation error": {
		InputFilePaths: nil,
		InputRunAll:    true,
		ExpectedError:  cmd.FilePathsArgEmpty,
	},
	"make sure that an empty file paths which is an array with a single empty value causes a validation error": {
		InputFilePaths: []string{"   	"},
		InputRunAll:    true,
		ExpectedError:  cmd.FilePathsArgEmpty,
	},
	"make sure that empty css paths when run all changes is true causes a validation error": {
		InputFilePaths: []string{"value"},
		InputCssPaths:  nil,
		InputRunAll:    true,
		ExpectedError:  cmd.CssPathsEmptyWhenArgIsNeeded,
	},
	"make sure that empty css paths which is an array with a single empty value when run all changes is true causes a validation error": {
		InputFilePaths: []string{"value"},
		InputCssPaths:  []string{"   	"},
		InputRunAll:    true,
		ExpectedError:  cmd.CssPathsEmptyWhenArgIsNeeded,
	},
	"make sure that empty css paths when running a section brake change causes a validation error": {
		InputFilePaths:       []string{"value"},
		InputCssPaths:        nil,
		InputRunSectionBreak: true,
		ExpectedError:        cmd.CssPathsEmptyWhenArgIsNeeded,
	},
	"make sure that empty css paths when running a page brake change causes a validation error": {
		InputFilePaths:    []string{"value"},
		InputCssPaths:     nil,
		InputRunPageBreak: true,
		ExpectedError:     cmd.CssPathsEmptyWhenArgIsNeeded,
	},
	"make sure that empty css paths when run all, run page break, and run section break changes are false passes validation": {
		InputFilePaths:      []string{"value"},
		InputCssPaths:       nil,
		InputRunBrokenLines: true,
		ExpectedError:       "",
	},
	"make sure that no boolean flags being on causes a validation error": {
		InputFilePaths: []string{"value"},
		InputCssPaths:  nil,
		ExpectedError:  cmd.OneRunBoolArgMustBeEnabled,
	},
}

func TestValidateManuallyFixableFlags(t *testing.T) {
	for name, args := range ValidateManuallyFixableFlagsTestCases {
		t.Run(name, func(t *testing.T) {
			err := cmd.ValidateManuallyFixableFlags(args.InputFilePaths, args.InputCssPaths, args.InputRunAll, args.InputRunBrokenLines, args.InputRunSectionBreak, args.InputRunPageBreak, args.InputRunOxfordCommas, args.InputRunAlthoughBut)

			if err != nil {
				assert.Equal(t, args.ExpectedError, err.Error())
			} else {
				assert.Equal(t, args.ExpectedError, "")
			}
		})
	}
}
