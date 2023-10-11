//go:build unit

package cmd_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/song-converter/cmd"
	"github.com/stretchr/testify/assert"
)

type ValidateCreateCoverFlagsTestCase struct {
	InputCoverPath string
	ExpectedError  string
}

// errors that get handled as errors are represented as panics
var ValidateCreateCoverFlagsTestCases = map[string]ValidateCreateCoverFlagsTestCase{
	"make sure that an empty cover path causes a validation error": {
		InputCoverPath: "",
		ExpectedError:  cmd.CoverPathArgEmpty,
	},
	"make sure that an non-md styles path causes a validation error": {
		InputCoverPath: "cover.txt",
		ExpectedError:  cmd.CoverPathNotMdFile,
	},
}

func TestValidateCreateCoverFlags(t *testing.T) {
	for name, args := range ValidateCreateCoverFlagsTestCases {
		t.Run(name, func(t *testing.T) {
			err := cmd.ValidateCreateCoverFlags(args.InputCoverPath)
			if err != nil {
				assert.Equal(t, args.ExpectedError, err.Error())
			} else {
				assert.Equal(t, args.ExpectedError, "")
			}
		})
	}
}
