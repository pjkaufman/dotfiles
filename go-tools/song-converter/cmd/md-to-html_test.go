//go:build unit

package cmd_test

import (
	"fmt"
	"testing"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/cmd"
	"github.com/stretchr/testify/assert"
)

type MdToHtmlTestCase struct {
	InputFilePath       string
	ExistingFiles       map[string]struct{}
	ExistingFolders     map[string]struct{}
	ExistingFileContent map[string]string
	PathToFolders       map[string][]string
	CmdErr              error
	ExpectedError       string
	ExpectPanic         bool
	ExpectedHtml        string
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
	"a valid file should properly get turned into html": {
		InputFilePath: "file.md",
		ExistingFiles: map[string]struct{}{
			"file.md": {},
		},
		ExistingFileContent: map[string]string{
			"file.md": `---
melody:
key: Key C
authors: Chris Knauf
in-church: N
verse:
location: (MS16) (B20)
type: song
tags:
---

# He Is

\~ 1 \~ He is fairer than the lily of the valley, He is brighter than the morning star.  
He is purer than the snow, fresher than the breeze, Lovelier by far than all of these.

\~ 2 \~ But He calms all the storms, and conquers the raging seas.  
He is the high and lofty One who inhabits eternity,  
Creator of the universe, and He\'s clothed with majesty.  
He is and ever more shall be.  
(Repeat first stanza)
			`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedHtml: `<div class="keep-together">
<div><div class="metadata"><div><div class="author">Chris Knauf</div></div><div><div class="key">Key C</div></div><div><div class="location">(MS16) (B20)</div></div></div></div><br/><br/>
<h1 id="he-is">He Is</h1>

<p>~ 1 ~ He is fairer than the lily of the valley, He is brighter than the morning star.</p>
<p>He is purer than the snow, fresher than the breeze, Lovelier by far than all of these.</p>

<p>~ 2 ~ But He calms all the storms, and conquers the raging seas.</p>
<p>He is the high and lofty One who inhabits eternity,</p>
<p>Creator of the universe, and He&rsquo;s clothed with majesty.</p>
<p>He is and ever more shall be.</p>
<p>(Repeat first stanza)</p>
</div>
<br/>`,
	},
	// "make sure that when there is an error the error is the expected error message (i.e. the panic value)": {
	// 	CmdErr:        errors.New("exit status 1"),
	// 	ExpectedError: "failed to update the submodule for the current repo: exit status 1",
	// 	ExpectPanic:   true,
	// },
}

/**

 */

func TestMdToHtml(t *testing.T) {
	for name, args := range MdToHtmlTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleMdToHtmlPanic(t, args)

			var log = logger.NewMockLoggerHandler()
			var fileHandler = filehandler.NewMockFileHandler(log, args.ExistingFiles, nil, nil, args.ExistingFileContent)
			actual := cmd.MdToHtml(log, fileHandler, args.InputFilePath)

			fmt.Println(actual)
			assert.Equal(t, args.ExpectedHtml, actual)
		})
	}
}

func handleMdToHtmlPanic(t *testing.T, args MdToHtmlTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
