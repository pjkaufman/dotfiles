//go:build unit

package cmd_test

import (
	"testing"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/cmd"
	"github.com/stretchr/testify/assert"
)

type CreateCoverTestCase struct {
	InputCoverPath      string
	ExistingFiles       map[string]struct{}
	ExistingFolders     map[string]struct{}
	ExistingFileContent map[string]string
	ExpectedError       string
	ExpectPanic         bool
	ExpectedHtml        string
}

// errors that get handled as errors are represented as panics
var CreateCoverTestCases = map[string]CreateCoverTestCase{
	"make sure that an empty cover path causes a validation error": {
		InputCoverPath: "",
		ExpectedError:  cmd.CoverPathArgEmpty,
		ExpectPanic:    true,
	},
	"make sure that an non-md styles path causes a validation error": {
		InputCoverPath: "cover.txt",
		ExpectedError:  cmd.CoverPathNotMdFile,
		ExpectPanic:    true,
	},
	"make sure that the cover path not existing causes a validation error": {
		InputCoverPath: "cover.md",
		ExpectedError:  `cover-file: "cover.md" must exist`,
		ExpectPanic:    true,
	},
}

func TestCreateCover(t *testing.T) {
	for name, args := range CreateCoverTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleCreateCoverPanic(t, args)

			var log = logger.NewMockLoggerHandler()
			var fileHandler = filehandler.NewMockFileHandler(log, args.ExistingFiles, args.ExistingFolders, nil, args.ExistingFileContent)
			cmd.CreateCover(log, fileHandler, args.InputCoverPath, "")
		})
	}
}

func handleCreateCoverPanic(t *testing.T, args CreateCoverTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
