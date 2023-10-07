//go:build unit

package linter_test

import (
	"fmt"
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/stretchr/testify/assert"
)

type GetPotentialSectionBreaksTestCase struct {
	InputText           string
	InputContextBreak   string
	ExpectedSuggestions map[string]string
}

var GetPotentialSectionBreaksTestCases = map[string]GetPotentialSectionBreaksTestCase{
	"make sure that a file with no section breaks gives no suggestions": {
		InputText: `<p>Here is some content.</p>
<p>Here is some more content</p>`,
		InputContextBreak:   contextBreak,
		ExpectedSuggestions: map[string]string{},
	},
	"make sure that a file a couple section breaks gives suggestions": {
		InputText: fmt.Sprintf(`<p>Here is some content.</p>
<p>%[1]s</p>
<p><a id="pg10"></a>%[1]s</p>
<p>Here is some more content</p>`, contextBreak),
		InputContextBreak: contextBreak,
		ExpectedSuggestions: map[string]string{
			fmt.Sprintf("\n<p>%s</p>", contextBreak):                    "\n" + linter.SectionBreakEl,
			fmt.Sprintf("\n<p><a id=\"pg10\"></a>%s</p>", contextBreak): "\n<p><a id=\"pg10\"></a>" + linter.SectionBreakEl + "</p>",
		},
	},
}

func TestGetPotentialSectionBreaks(t *testing.T) {
	for name, args := range GetPotentialSectionBreaksTestCases {
		t.Run(name, func(t *testing.T) {
			actual := linter.GetPotentialSectionBreaks(args.InputText, args.InputContextBreak)

			assert.Equal(t, args.ExpectedSuggestions, actual)
		})
	}
}
