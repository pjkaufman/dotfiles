package linter

import (
	"regexp"
)

const pageBrakeEl = `<hr class="blankSpace" />`

var emptyParagraphsOrDivs = regexp.MustCompile(`(\n<(p|div)[^\n>]*>)[ \t]*(</(p|div)>)`)

func GetPotentialPageBreaks(fileContent string) map[string]string {
	var subMatches = emptyParagraphsOrDivs.FindAllStringSubmatch(fileContent, -1)
	var originalToSuggested = make(map[string]string, len(subMatches))
	if len(subMatches) == 0 {
		return originalToSuggested
	}

	for _, groups := range subMatches {
		originalToSuggested[groups[0]] = "\n" + pageBrakeEl
	}

	return originalToSuggested
}
