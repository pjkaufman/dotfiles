//go:build unit

package cmd_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/cmd"
	"github.com/stretchr/testify/assert"
)

type ValidateReplaceStringsFlagsTestCase struct {
	InputFilePaths               []string
	InputExtraReplaceStringsPath string
	ExpectedError                string
}

var ValidateReplaceStringsFlagsTestCases = map[string]ValidateReplaceStringsFlagsTestCase{
	"make sure that an empty file paths causes a validation error": {
		InputFilePaths:               nil,
		InputExtraReplaceStringsPath: "file.md",
		ExpectedError:                cmd.FilePathsArgEmpty,
	},
	"make sure that an empty file paths which is an array with a single empty value causes a validation error": {
		InputFilePaths:               []string{"   	"},
		InputExtraReplaceStringsPath: "file.md",
		ExpectedError:                cmd.FilePathsArgEmpty,
	},
	"make sure that an empty extra string replace path causes a validation error": {
		InputFilePaths:               []string{"value"},
		InputExtraReplaceStringsPath: "",
		ExpectedError:                cmd.ExtraStringReplaceArgEmpty,
	},
	"make sure that a non-md extra string replace path causes a validation error": {
		InputFilePaths:               []string{"value"},
		InputExtraReplaceStringsPath: "file.txt",
		ExpectedError:                cmd.ExtraStringReplaceArgNonMd,
	},
	"make sure that an extra string replace path as an md file and a value for file paths passes validation": {
		InputFilePaths:               []string{"value"},
		InputExtraReplaceStringsPath: "file.md",
		ExpectedError:                "",
	},
}

func TestValidateReplaceStringsFlags(t *testing.T) {
	for name, args := range ValidateReplaceStringsFlagsTestCases {
		t.Run(name, func(t *testing.T) {
			err := cmd.ValidateReplaceStringsFlags(args.InputFilePaths, args.InputExtraReplaceStringsPath)

			if err != nil {
				assert.Equal(t, args.ExpectedError, err.Error())
			} else {
				assert.Equal(t, args.ExpectedError, "")
			}
		})
	}
}
