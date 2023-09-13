//go:build unit

package linter_test

import (
	"fmt"
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
)

type ExtraStringReplaceTestCase struct {
	InputText             string
	InputFindsAndReplaces map[string]string
	InputHits             map[string]int
	ExpectedText          string
	ExpectedHits          map[string]int
}

var ExtraStringReplaceTestCases = map[string]ExtraStringReplaceTestCase{
	"make sure that when a replacement is made with an empty map, the number of hits in the map is updated accordingly": {
		InputText: `Here is some text that gets broken into
		multiple lines with a couple of words to be replaced`,
		InputFindsAndReplaces: map[string]string{
			"Here is":        "This was",
			"to be replaced": "that were replaced",
		},
		InputHits: map[string]int{},
		ExpectedHits: map[string]int{
			"Here is":        1,
			"to be replaced": 1,
		},
		ExpectedText: `This was some text that gets broken into
		multiple lines with a couple of words that were replaced`,
	},
	"make sure that when multiple instances of a value to replace in a string are present that all of them get replaced": {
		InputText: `I talk way too much as if I were not going to get another chance to talk to myself. I wonder why that is.`,
		InputFindsAndReplaces: map[string]string{
			"I": "You",
		},
		InputHits: map[string]int{},
		ExpectedHits: map[string]int{
			"I": 3,
		},
		ExpectedText: `You talk way too much as if You were not going to get another chance to talk to myself. You wonder why that is.`,
	},
	"make sure that not finding a value in a file when it does not already exist just sets the value for that search value to 0": {
		InputText: `Text not found`,
		InputFindsAndReplaces: map[string]string{
			"I": "You",
		},
		InputHits: map[string]int{},
		ExpectedHits: map[string]int{
			"I": 0,
		},
		ExpectedText: `Text not found`,
	},
	"make sure that not finding a value in a file when it does not already exist does not affect the resulting hit count": {
		InputText: `Text not found`,
		InputFindsAndReplaces: map[string]string{
			"I": "You",
		},
		InputHits: map[string]int{
			"I": 5,
		},
		ExpectedHits: map[string]int{
			"I": 5,
		},
		ExpectedText: `Text not found`,
	},
	"make sure that when a replacement is made and the value already exists in the hit count it gets incremented": {
		InputText: `This is not what I expected. This could get dangerous. This is not what I signed up for!`,
		InputFindsAndReplaces: map[string]string{
			"This": "That",
		},
		InputHits: map[string]int{
			"This": 2,
		},
		ExpectedHits: map[string]int{
			"This": 5,
		},
		ExpectedText: `That is not what I expected. That could get dangerous. That is not what I signed up for!`,
	},
}

func TestExtraStringReplace(t *testing.T) {
	for name, args := range ExtraStringReplaceTestCases {
		t.Run(name, func(t *testing.T) {
			actual := linter.ExtraStringReplace(args.InputText, args.InputFindsAndReplaces, args.InputHits)

			if actual != args.ExpectedText {
				t.Errorf("output text doesn't match: expected \"%s\", got \"%s\"", args.ExpectedText, actual)
			}

			if !stringToIntMapsAreEqual(args.ExpectedHits, args.InputHits) {
				t.Errorf("output map doesn't match: expected %v, got %v", args.ExpectedHits, args.InputHits)
			}
		})
	}
}

func stringToIntMapsAreEqual(expected, actual map[string]int) bool {
	if len(expected) != len(actual) {
		return false
	}

	for key, value := range expected {
		if value2, found := actual[key]; !found || value != value2 {
			if !found {
				fmt.Printf("expected value \"%s\" but did not find it", key)
			} else {
				fmt.Printf("expected value \"%d\" but got \"%d\"", value, value2)
			}

			return false
		}
	}

	return true
}
