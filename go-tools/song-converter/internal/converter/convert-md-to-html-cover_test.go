//go:build unit

package converter_test

import (
	"testing"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/internal/converter"
	"github.com/stretchr/testify/assert"
)

type ConvertMdToHtmlCoverTestCase struct {
	InputFilePath       string
	ExistingFiles       map[string]struct{}
	ExistingFileContent map[string]string
	ExpectedError       string
	ExpectPanic         bool
	ExpectedHtml        string
}

// errors that get handled as errors are represented as panics
var ConvertMdToHtmlCoverTestCases = map[string]ConvertMdToHtmlCoverTestCase{
	"make sure that the file path not existing causes an error": {
		InputFilePath: "file.md",
		ExpectedError: `could not read in file contents for "file.md": path not found`,
		ExpectPanic:   true,
	},
	"a valid file should properly get turned into html": {
		InputFilePath: "file.md",
		ExistingFiles: map[string]struct{}{
			"file.md": {},
		},
		ExistingFileContent: map[string]string{
			"file.md": `# Church Songs - E Version

<br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/>

### Key:

#### Books

R=Red Book (Songs We Love)  
MS=More Songs section of Songs We Love  
B=Blue Book

#### Authors

CA= Cyndi Aarrestad  
EHW= Ewald Wanagas  
FTP= Frank Paterson  
GBS= Gail Shepherd  
ZW= Zelma Wanagas

<br/> <br/>

*\*When searching electronic format, punctuation & spelling will cause
non returns*

*\*Punctuation alters the alphabetical order*`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedHtml: `<div style="text-align: center">
<h1 id="church-songs-e-version">Church Songs - E Version</h1>
<p><br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/></p>
<h3 id="key">Key:</h3>
<h4 id="books">Books</h4>
<p>R=Red Book (Songs We Love)<br>
MS=More Songs section of Songs We Love<br>
B=Blue Book</p>
<h4 id="authors">Authors</h4>
<p>CA= Cyndi Aarrestad<br>
EHW= Ewald Wanagas<br>
FTP= Frank Paterson<br>
GBS= Gail Shepherd<br>
ZW= Zelma Wanagas</p>
<p><br/> <br/></p>
<p><em>*When searching electronic format, punctuation &amp; spelling will cause
non returns</em></p>
<p><em>*Punctuation alters the alphabetical order</em></p>
</div>
`,
	},
}

func TestConvertMdToHtmlCover(t *testing.T) {
	for name, args := range ConvertMdToHtmlCoverTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleConvertMdToHtmlCoverPanic(t, args)

			var log = logger.NewMockLoggerHandler()
			var fileHandler = filehandler.NewMockFileHandler(log, args.ExistingFiles, nil, nil, args.ExistingFileContent)
			actual := converter.ConvertMdToHtmlCover(log, fileHandler, args.InputFilePath)

			assert.Equal(t, args.ExpectedHtml, actual)
		})
	}
}

func handleConvertMdToHtmlCoverPanic(t *testing.T, args ConvertMdToHtmlCoverTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
