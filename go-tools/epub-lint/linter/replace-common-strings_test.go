//go:build unit

package linter_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/stretchr/testify/assert"
)

type CommonStringReplaceTestCase struct {
	Input    string
	Expected string
}

var commonStringReplaceTestCases = map[string]CommonStringReplaceTestCase{
	"make sure that html comments are left alone": {
		Input:    "<!--this is a comment. comments are not displayed in the browser-->",
		Expected: "<!--this is a comment. comments are not displayed in the browser-->",
	},
	"make sure that two en dashes are replaced with an em dash": {
		Input:    "-- test --",
		Expected: "— test —",
	},
	"make sure that three periods with a 0 or 1 spaces between them get cut down to proper ellipsis": {
		Input: `
		  ...
		  . . .
		  . .. 
		  .. .
		  .  . .
		`,
		Expected: `
		  …
		  …
		  … 
		  …
		  …
		`,
	},
	"make sure that a lowercase 'by the by' results in a lowercase 'by the way'": {
		Input:    "by the by",
		Expected: "by the way",
	},
	"make sure that an uppercase 'By the by' results in an uppercase 'By the way'": {
		Input:    "By the by",
		Expected: "By the way",
	},
	"make sure that an uppercase 'Sneaked' results in an uppercase 'Snuck'": {
		Input:    "Sneaked",
		Expected: "Snuck",
	},
	"make sure that a lowercase 'snuck' results in a lowercase 'snuck'": {
		Input:    "On his way he sneaked out the door",
		Expected: "On his way he snuck out the door",
	},
	"make sure that single tilde is converted to an exclamation mark": {
		Input:    "~wow isn't this a joy~",
		Expected: "!wow isn't this a joy!",
	},
	"make sure that multiple tildes in a row are not converted to an exclamation mark": {
		Input:    "~~ is completely ~~~ left alone",
		Expected: "~~ is completely ~~~ left alone",
	},
	"make sure that a lowercase 'a bolt out of the blue' is correctly converted to 'out of the blue": {
		Input:    "a bolt out of the blue",
		Expected: "out of the blue",
	},
	"make sure that an uppercase 'A bolt out of the blue' is correctly converted to 'Out of the blue": {
		Input:    "A bolt out of the blue, a loud bang went off waking up people who had been sleeping soundly up until that point.",
		Expected: "Out of the blue, a loud bang went off waking up people who had been sleeping soundly up until that point.",
	},
	"make sure that a lowercase 'little wonder' is correctly converted to 'no wonder": {
		Input:    "little wonder your attempt failed",
		Expected: "no wonder your attempt failed",
	},
	"make sure that an uppercase 'Little wonder' is correctly converted to 'No wonder": {
		Input:    "Little wonder, you were outmatched from the start",
		Expected: "No wonder, you were outmatched from the start",
	},
	"make sure that words with 2 or more spaces between them have the multiple spaces cut down to 1": {
		Input:    "This  is an    interestingly spaced   sentence.  See the multiple    blanks?",
		Expected: "This is an interestingly spaced sentence. See the multiple blanks?",
	},
	"make sure that smart double quotes are replaced with straight quotes": {
		Input: `“Hey. How are you?”
		“I am doing great!”`,
		Expected: `"Hey. How are you?"
		"I am doing great!"`,
	},
	"make sure that smart single quotes are replaced with straight quotes": {
		Input: `‘Hey. How are you?’
		‘I am doing great!’`,
		Expected: `'Hey. How are you?'
		'I am doing great!'`,
	},
	// "make sure that simple missing oxford comma situations get handled properly": {
	// 	Input:    `Here is a situation where I run, skip and jump for a long time.`,
	// 	Expected: `Here is a situation where I run, skip, and jump for a long time.`,
	// },
	// "make sure that simple missing oxford comma situations where it is just the list on the line are handled properly": {
	// 	Input: `Here is a situation where I
	// 	run, walk, sleep, skip or jump
	// 	for a long time.`,
	// 	Expected: `Here is a situation where I
	// 	run, walk, sleep, skip, or jump
	// 	for a long time.`,
	// },
	// "make sure that complex scenarios missing the oxford comma are ignored": {
	// 	Input:    `Here is a situation where red and white, blue and green and white is the list`,
	// 	Expected: `Here is a situation where red and white, blue and green and white is the list`,
	// },
	// "make sure that complex scenarios missing the oxford comma with different conjunctions are ignored": {
	// 	Input:    `Here is a situation where red and white, blue and green or white is the list`,
	// 	Expected: `Here is a situation where red and white, blue and green or white is the list`,
	// },
}

func TestCommonStringReplace(t *testing.T) {
	for name, args := range commonStringReplaceTestCases {
		t.Run(name, func(t *testing.T) {
			actual := linter.CommonStringReplace(args.Input)

			assert.Equal(t, args.Expected, actual)
		})
	}

}
