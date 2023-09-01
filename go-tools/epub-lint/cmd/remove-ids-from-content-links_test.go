//go:build unit

package cmd_test

type RemoveIdsFromContentLinksTestCase struct {
	InputText       string
	ExpectedText    string
	ExpectedFileMap map[string]int
}

// var RemoveIdsFromContentLinksTestCases = map[string]RemoveIdsFromContentLinksTestCase{
// 	"make sure that html comments are left alone": {
// 		Input:    "<!--this is a comment. comments are not displayed in the browser-->",
// 		Expected: "<!--this is a comment. comments are not displayed in the browser-->",
// 	},
// 	"make sure that two en dashes are replaced with an em dash": {
// 		Input:    "-- test --",
// 		Expected: "— test —",
// 	},
// 	"make sure that three periods with a 0 or 1 spaces between them get cut down to proper ellipsis": {
// 		Input: `
// 		  ...
// 		  . . .
// 		  . ..
// 		  .. .
// 		  .  . .
// 		`,
// 		Expected: `
// 		  …
// 		  …
// 		  …
// 		  …
// 		  .  . .
// 		`,
// 	},
// 	"make sure that a lowercase 'by the by' results in a lowercase 'by the way'": {
// 		Input:    "by the by",
// 		Expected: "by the way",
// 	},
// 	"make sure that an uppercase 'By the by' results in an uppercase 'By the way'": {
// 		Input:    "By the by",
// 		Expected: "By the way",
// 	},
// 	"make sure that an uppercase 'Sneaked' results in an uppercase 'Snuck'": {
// 		Input:    "Sneaked",
// 		Expected: "Snuck",
// 	},
// 	"make sure that a lowercase 'snuck' results in a lowercase 'snuck'": {
// 		Input:    "On his way he sneaked out the door",
// 		Expected: "On his way he snuck out the door",
// 	},
// 	"make sure that single tilde is converted to an exclamation mark": {
// 		Input:    "~wow isn't this a joy~",
// 		Expected: "!wow isn't this a joy!",
// 	},
// 	"make sure that multiple tildes in a row are not converted to an exclamation mark": {
// 		Input:    "~~ is completely ~~~ left alone",
// 		Expected: "~~ is completely ~~~ left alone",
// 	},
// 	"make sure that a lowercase 'a bolt out of the blue' is correctly converted to 'out of the blue": {
// 		Input:    "a bolt out of the blue",
// 		Expected: "out of the blue",
// 	},
// 	"make sure that an uppercase 'A bolt out of the blue' is correctly converted to 'Out of the blue": {
// 		Input:    "A bolt out of the blue, a loud bang went off waking up people who had been sleeping soundly up until that point.",
// 		Expected: "Out of the blue, a loud bang went off waking up people who had been sleeping soundly up until that point.",
// 	},
// 	"make sure that a lowercase 'little wonder' is correctly converted to 'no wonder": {
// 		Input:    "little wonder your attempt failed",
// 		Expected: "no wonder your attempt failed",
// 	},
// 	"make sure that an uppercase 'Little wonder' is correctly converted to 'No wonder": {
// 		Input:    "Little wonder, you were outmatched from the start",
// 		Expected: "No wonder, you were outmatched from the start",
// 	},
// }

// func TestRemoveIdsFromContentLinks(t *testing.T) {
// 	for name, args := range RemoveIdsFromContentLinksTestCases {
// 		t.Run(name, func(t *testing.T) {
// 			actual := cmd.RemoveIdsFromContentLinks(args.Input)

// 			if actual != args.Expected {
// 				t.Errorf("output text doesn't match: expected %v, got %v", args.Expected, actual)
// 			}
// 		})
// 	}

// }
