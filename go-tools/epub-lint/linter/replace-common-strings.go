package linter

import (
	"fmt"
	"regexp"
)

type ReplaceString struct {
	Search   *regexp.Regexp
	Replace  string
	Rational string
}

type ReplaceStringFunc struct {
	Search   *regexp.Regexp
	Replace  func(string) string
	Rational string
}

var regexEscapedPeriod = regexp.QuoteMeta(".")
var commonReplaceStrings = []ReplaceString{
	{
		// [^\w\s] means any non-whitespace or alphanumeric values or an underscore
		Search:   regexp.MustCompile(`(\b|[^\w\s])( ){2,}(\b|[^\w\s])`),
		Replace:  "${1} ${3}",
		Rational: "Replace multiple spaces in a row between words with a single space since this can cause issues with replace strings",
	},
	{
		Search:   regexp.MustCompile(fmt.Sprintf("(%[1]s ?){2}%[1]s", regexEscapedPeriod)),
		Replace:  "…",
		Rational: "Proper ellipses should be used where possible as it keeps things clean and consistent",
	},
	{
		Search:   regexp.MustCompile("(^|[^!])--([^>]|$)"),
		Replace:  "$1—$2",
		Rational: "Em dashes should be used where possible as it keeps things clean and consistent",
	},
	{
		Search:   regexp.MustCompile("(^|[^~])~([^~]|$)"),
		Replace:  "$1!$2",
		Rational: "Tildes should be replaced with an exclamation mark when they are by themselves as they seem interchangeable, though it could be another form of punctuation for drawing out the sound of the last letter used",
	},
	{
		Search:   regexp.MustCompile("(B|b)y the by"),
		Replace:  "${1}y the way",
		Rational: "'By the by' seems to be an improper translation of 'By the way', so we should auto-correct it to its proper English idiom",
	},
	{
		Search:   regexp.MustCompile("(S|s)neaked"),
		Replace:  "${1}nuck",
		Rational: "Use snuck instead of sneaked as it is the more commonly used version of the word nowadays",
	},
}

var commonReplaceStringFuncs = []ReplaceStringFunc{
	{
		Search: regexp.MustCompile("(A|a) bolt (o)(ut of the blue)"),
		Replace: func(part string) string {
			var firstLetter = "O"
			if part[0] == 'a' {
				firstLetter = "o"
			}

			return firstLetter + "ut of the blue"
		},
		Rational: "'a bolt out of the blue' seems to be an improper translation of 'out of the blue', so we should auto-correct it to its proper English idiom",
	},
	{
		Search: regexp.MustCompile("(L|l)ittle( wonder)"),
		Replace: func(part string) string {
			var firstLetter = "N"
			if part[0] == 'l' {
				firstLetter = "n"
			}

			return firstLetter + "o wonder"
		},
		Rational: "'little wonder' seems to be an improper translation of 'no wonder', so we should auto-correct it to its proper English idiom",
	},
}

func CommonStringReplace(text string) string {
	var newText = text

	for _, replaceString := range commonReplaceStrings {
		newText = replaceString.Search.ReplaceAllString(newText, replaceString.Replace)
	}

	for _, replaceStringFunc := range commonReplaceStringFuncs {
		newText = replaceStringFunc.Search.ReplaceAllStringFunc(newText, replaceStringFunc.Replace)
	}

	return newText
}
