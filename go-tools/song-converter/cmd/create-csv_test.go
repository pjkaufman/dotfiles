//go:build unit

package cmd_test

import (
	"testing"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/cmd"
	"github.com/stretchr/testify/assert"
)

type CreateCsvTestCase struct {
	InputStagingDir     string
	ExistingFiles       map[string]struct{}
	ExistingFolders     map[string]struct{}
	ExistingFileContent map[string]string
	ExpectedError       string
	ExpectPanic         bool
	ExpectedHtml        string
}

// errors that get handled as errors are represented as panics
var CreateCsvTestCases = map[string]CreateCsvTestCase{
	"make sure that an empty working dir causes a validation error": {
		InputStagingDir: "",
		ExpectedError:   cmd.StagingDirArgEmpty,
		ExpectPanic:     true,
	},
	"make sure that working dir not existing causes a validation error": {
		InputStagingDir: "does not exist",
		ExpectedError:   `working-dir: "does not exist" must exist`,
		ExpectPanic:     true,
	},
}

func TestCreateCsv(t *testing.T) {
	for name, args := range CreateCsvTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleCreateCsvPanic(t, args)

			var log = logger.NewMockLoggerHandler()
			var fileHandler = filehandler.NewMockFileHandler(log, args.ExistingFiles, args.ExistingFolders, nil, args.ExistingFileContent)
			cmd.CreateCsv(log, fileHandler, args.InputStagingDir, "")
		})
	}
}

func handleCreateCsvPanic(t *testing.T, args CreateCsvTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
