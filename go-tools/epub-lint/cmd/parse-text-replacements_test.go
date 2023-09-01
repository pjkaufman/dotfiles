//go:build unit

package cmd_test

import (
	"fmt"
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/cmd"
)

type ParseTextReplacementsTestCase struct {
	Input    string
	Expected map[string]string
}

var ParseTextReplacementsTestCases = map[string]ParseTextReplacementsTestCase{
	"make sure that an empty table results in an empty map": {
		Input: `| Text to replace | Text replacement |
		| ---- | ---- |`,
		Expected: map[string]string{},
	},
	"make sure that a non-empty table results in the appropriate amount of entries being placed in a map": {
		Input: `| Text to replace | Text replacement |
		| ---- | ---- |
		| replace | with me |
		| "I am quoted" | 'I am single quoted' |`,
		Expected: map[string]string{
			"replace":         "with me",
			"\"I am quoted\"": "'I am single quoted'",
		},
	},
	"make sure that values get trimmed before getting added to the map": {
		Input: `| Text to replace | Text replacement |
		| ---- | ---- |
		| replace | with me |
		| "I am quoted" | 'I am single quoted' |
		|       I have lots of whitespace around me      | I have   wonky internal spacing |`,
		Expected: map[string]string{
			"replace":                             "with me",
			"\"I am quoted\"":                     "'I am single quoted'",
			"I have lots of whitespace around me": "I have   wonky internal spacing",
		},
	},
	"make sure that lines without a pipe/table row get ignored": {
		Input: `| Text to replace | Text replacement |
		| ---- | ---- |
		| replace | with me |
		| "I am quoted" | 'I am single quoted' |
		|       I have lots of whitespace around me      | I have   wonky internal spacing |
		Some text here
		Another line here`,
		Expected: map[string]string{
			"replace":                             "with me",
			"\"I am quoted\"":                     "'I am single quoted'",
			"I have lots of whitespace around me": "I have   wonky internal spacing",
		},
	},
}

func TestParseTextReplacements(t *testing.T) {
	for name, args := range ParseTextReplacementsTestCases {
		t.Run(name, func(t *testing.T) {
			actual := cmd.ParseTextReplacements(args.Input)

			if !stringMapsAreEqual(args.Expected, actual) {
				t.Errorf("output map doesn't match: expected %v, got %v", args.Expected, actual)
			}
		})
	}
}

func stringMapsAreEqual(expected, actual map[string]string) bool {
	if len(expected) != len(actual) {
		return false
	}

	for key, value := range expected {
		if value2, found := actual[key]; !found || value != value2 {
			if !found {
				fmt.Printf("expected value \"%s\" but did not find it", key)
			} else {
				fmt.Printf("expected value \"%s\" but got \"%s\"", value, value2)
			}

			return false
		}
	}

	return true
}
