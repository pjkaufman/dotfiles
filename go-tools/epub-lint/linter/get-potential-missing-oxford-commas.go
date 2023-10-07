package linter

import (
	"regexp"
)

var oxfordCommaRegex = regexp.MustCompile(`(\n<p[^\n>]*>[^\n]*)(\w+)((,\s*\w+)+)(\s+)(and|or)(\s+\w+)([^\n]*</p>)`)

func GetPotentialMissingOxfordCommas(fileContent string) map[string]string {
	var subMatches = oxfordCommaRegex.FindAllStringSubmatch(fileContent, -1)
	var originalToSuggested = make(map[string]string, len(subMatches))
	if len(subMatches) == 0 {
		return originalToSuggested
	}

	for _, groups := range subMatches {
		originalToSuggested[groups[0]] = groups[1] + groups[2] + groups[3] + "," + groups[5] + groups[6] + groups[7] + groups[8]
	}

	return originalToSuggested
}
