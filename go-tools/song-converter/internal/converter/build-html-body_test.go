//go:build unit

package converter_test

import (
	"fmt"
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/song-converter/internal/converter"
	"github.com/stretchr/testify/assert"
)

type BuildHtmlBodyTestCase struct {
	InputStylesHtml string
	InputMdInfo     []converter.MdFileInfo
	ExpectedHtml    string
}

// TODO: add one that has styles and regular content
var BuildHtmlBodyTestCases = map[string]BuildHtmlBodyTestCase{
	"no files provided along with no styles should just result in a blank line": {
		InputStylesHtml: "",
		ExpectedHtml:    "\n",
	},
	"multiple files with empty styles should be a new line character followed by each song with a new line character after it": {
		InputStylesHtml: "",
		InputMdInfo: []converter.MdFileInfo{
			{
				FilePath:     "Above It All (There Stands Jesus).md",
				FileContents: AboveItAllFileMd,
			},
			{
				FilePath:     "Be Thou Exalted.md",
				FileContents: BeThouExaltedFileMd,
			},
			{
				FilePath:     "Behold The Heavens.md",
				FileContents: BeholdTheHeavensFileMd,
			},
			{
				FilePath:     "He Is.md",
				FileContents: HeIsFileMd,
			},
		},
		ExpectedHtml: fmt.Sprintf("\n%s\n%s\n%s\n%s\n", AboveItAllFileHtml, BeThouExaltedFileHtml, BeholdTheHeavensFileHtml, HeIsFileHtml),
	},
}

func TestBuildHtmlBody(t *testing.T) {
	for name, args := range BuildHtmlBodyTestCases {
		t.Run(name, func(t *testing.T) {

			actual, err := converter.BuildHtmlBody(args.InputStylesHtml, args.InputMdInfo)
			if err != nil {
				assert.Fail(t, "there should be no errors when parsing the YAML for the html UTs")
			}

			assert.Equal(t, args.ExpectedHtml, actual)
		})
	}
}
