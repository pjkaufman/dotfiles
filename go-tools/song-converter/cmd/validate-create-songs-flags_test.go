//go:build unit

package cmd_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/song-converter/cmd"
	"github.com/stretchr/testify/assert"
)

type ValidateCreateSongsFlagsTestCase struct {
	InputStagingDir string
	InputStylesPath string
	ExpectedError   string
}

var ValidateCreateSongsFlagsTestCases = map[string]ValidateCreateSongsFlagsTestCase{
	"make sure that an empty working dir causes a validation error": {
		InputStagingDir: "",
		ExpectedError:   cmd.StagingDirArgEmpty,
	},
	"make sure that an empty styles path causes a validation error": {
		InputStylesPath: "",
		InputStagingDir: "folder",
		ExpectedError:   cmd.StylesPathArgEmpty,
	},
	"make sure that an non-html styles path causes a validation error": {
		InputStylesPath: "styles.txt",
		InputStagingDir: "folder",
		ExpectedError:   cmd.StylesPathNotHtmlFile,
	},
	"make sure that the styles path and staging dir not being whitespace passes validation": {
		InputStylesPath: "styles.html",
		InputStagingDir: "folder",
		ExpectedError:   "",
	},
}

func TestValidateCreateSongsFlags(t *testing.T) {
	for name, args := range ValidateCreateSongsFlagsTestCases {
		t.Run(name, func(t *testing.T) {
			err := cmd.ValidateCreateSongsFlags(args.InputStagingDir, args.InputStylesPath)
			if err != nil {
				assert.Equal(t, args.ExpectedError, err.Error())
			} else {
				assert.Equal(t, args.ExpectedError, "")
			}
		})
	}
}
