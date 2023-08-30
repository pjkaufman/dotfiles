//go:build unit

package cmd_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/cmd"
)

type ParseTextReplacementsTestCase struct {
	Input    string
	Expected map[string]string
}

// TestCase(
//
//	name="make sure that an empty table results in an empty dictionary",
//
// input="""| Text to replace | Text replacement |
// | ---- | ---- |
//
//	""",
//	expected={},
//
// ),
// TestCase(
//
//	name="make sure that a non-empty table results in the appropriate amount of entries being placed in a dictionary",
//
// input="""| Text to replace | Text replacement |
// | ---- | ---- |
// | replace | with me |
// | "I am quoted" | 'I am single quoted' |
//
//	""",
//	expected={
//		'replace': 'with me',
//		'\"I am quoted\"': '\'I am single quoted\'',
//	},
//
// ),
var ParseTextReplacementsTestCases = map[string]ParseTextReplacementsTestCase{
	"make sure that an empty table results in an empty dictionary": {
		Input: `| Text to replace | Text replacement |
		| ---- | ---- |`,
		Expected: map[string]string{},
	},
	"make sure that a non-empty table results in the appropriate amount of entries being placed in a dictionary": {
		Input: `| Text to replace | Text replacement |
		| ---- | ---- |
		| replace | with me |
		| "I am quoted" | 'I am single quoted' |`,
		Expected: map[string]string{
			"replace":         "with me",
			"\"I am quoted\"": "'I am single quoted'",
		},
	},
}

func TestCParseTextReplacements(t *testing.T) {
	for name, args := range ParseTextReplacementsTestCases {
		t.Run(name, func(t *testing.T) {
			actual := cmd.ParseTextReplacements(args.Input)

			if !mapsAreEqual(args.Expected, actual) {
				t.Errorf("output map doesn't match: expected %v, got %v", args.Expected, actual)
			}
		})
	}
}

func mapsAreEqual(expected, actual map[string]string) bool {
	if len(expected) != len(actual) {
		return false
	}

	for key, value := range expected {
		if value2, found := actual[key]; !found || value != value2 {
			return false
		}
	}

	return true
}
