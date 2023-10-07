//go:build unit

package cmd_test

import (
	"testing"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/cmd"
	"github.com/stretchr/testify/assert"
)

type CreateSongsTestCase struct {
	InputStylesPath     string
	InputStagingDir     string
	ExistingFiles       map[string]struct{}
	ExistingFolders     map[string]struct{}
	ExistingFileContent map[string]string
	ExpectedError       string
	ExpectPanic         bool
	ExpectedHtml        string
}

// errors that get handled as errors are represented as panics
var CreateSongsTestCases = map[string]CreateSongsTestCase{
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
	"make sure that an empty styles path causes a validation error": {
		InputStylesPath: "",
		InputStagingDir: "folder",
		ExistingFolders: map[string]struct{}{
			"folder": {},
		},
		ExpectedError: cmd.StylesPathArgEmpty,
		ExpectPanic:   true,
	},
	"make sure that an non-html styles path causes a validation error": {
		InputStylesPath: "styles.txt",
		InputStagingDir: "folder",
		ExistingFolders: map[string]struct{}{
			"folder": {},
		},
		ExpectedError: cmd.StylesPathNotHtmlFile,
		ExpectPanic:   true,
	},
	"make sure that the styles path not existing causes a validation error": {
		InputStylesPath: "styles.html",
		InputStagingDir: "folder",
		ExistingFolders: map[string]struct{}{
			"folder": {},
		},
		ExpectedError: `styles-file: "styles.html" must exist`,
		ExpectPanic:   true,
	},
}

func TestCreateSongs(t *testing.T) {
	for name, args := range CreateSongsTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleCreateSongsPanic(t, args)

			var log = logger.NewMockLoggerHandler()
			var fileHandler = filehandler.NewMockFileHandler(log, args.ExistingFiles, args.ExistingFolders, nil, args.ExistingFileContent)
			cmd.CreateSongs(log, fileHandler, args.InputStagingDir, args.InputStylesPath, "")
		})
	}
}

func handleCreateSongsPanic(t *testing.T, args CreateSongsTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
