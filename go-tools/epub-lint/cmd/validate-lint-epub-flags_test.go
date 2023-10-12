//go:build unit

package cmd_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/cmd"
	"github.com/stretchr/testify/assert"
)

type ValidateLintEpubFlagsTestCase struct {
	InputOpfFile  string
	InputLang     string
	ExpectedError string
}

var ValidateLintEpubFlagsTestCases = map[string]ValidateLintEpubFlagsTestCase{
	"make sure that an empty file path causes a validation error": {
		InputOpfFile:  "",
		InputLang:     "en",
		ExpectedError: cmd.FilePathArgEmpty,
	},
	"make sure that a file path that is not an opf file causes a validation error": {
		InputOpfFile:  "package.txt",
		InputLang:     "en",
		ExpectedError: cmd.FilePathArgNonOpf,
	},
	"make sure that an empty lang causes a validation error": {
		InputOpfFile:  "package.opf",
		InputLang:     "",
		ExpectedError: cmd.LangArgEmpty,
	},
	"make sure that an opf file and a lang value passes validation": {
		InputOpfFile:  "package.opf",
		InputLang:     "en",
		ExpectedError: "",
	},
}

func TestValidateLintEpubFlags(t *testing.T) {
	for name, args := range ValidateLintEpubFlagsTestCases {
		t.Run(name, func(t *testing.T) {
			err := cmd.ValidateLintEpubFlags(args.InputOpfFile, args.InputLang)

			if err != nil {
				assert.Equal(t, args.ExpectedError, err.Error())
			} else {
				assert.Equal(t, args.ExpectedError, "")
			}
		})
	}
}
