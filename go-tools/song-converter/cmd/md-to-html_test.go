//go:build unit

package cmd_test

import (
	"testing"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/cmd"
	"github.com/stretchr/testify/assert"
)

type MdToHtmlTestCase struct {
	InputFilePath   string
	ExistingFiles   map[string]struct{}
	ExistingFolders map[string]struct{}
	PathToFolders   map[string][]string
	CmdErr          error
	ExpectedError   string
	ExpectPanic     bool
}

// errors that get handled as errors are represented as panics
var MdToHtmlTestCases = map[string]MdToHtmlTestCase{
	"make sure that an empty file path causes a validation error": {
		InputFilePath: "",
		ExpectedError: cmd.FilePathArgEmpty,
		ExpectPanic:   true,
	},
	"make sure that an non-markdown file path causes a validation error": {
		InputFilePath: "file.txt",
		ExpectedError: cmd.FilePathNotMarkdownFile,
		ExpectPanic:   true,
	},
	"make sure that the file path not existing causes a validation error": {
		InputFilePath: "file.md",
		ExpectedError: `file-path: "file.md" must exist`,
		ExpectPanic:   true,
	},
	// "make sure that when there is an error the error is the expected error message (i.e. the panic value)": {
	// 	CmdErr:        errors.New("exit status 1"),
	// 	ExpectedError: "failed to update the submodule for the current repo: exit status 1",
	// 	ExpectPanic:   true,
	// },
}

func TestMdToHtml(t *testing.T) {
	for name, args := range MdToHtmlTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleMdToHtmlPanic(t, args)

			var log = logger.NewMockLoggerHandler()
			var fileHandler = filehandler.NewMockFileHandler(log, nil, nil, nil, nil)
			cmd.MdToHtml(log, fileHandler, args.InputFilePath)
		})
	}
}

func handleMdToHtmlPanic(t *testing.T, args MdToHtmlTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
