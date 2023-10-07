package linter

import (
	"fmt"
	"regexp"
	"strings"
)

const contextBreakEl = `<hr class="character" />`

func GetPotentialContextBreaks(fileContent, sectionBreakIndicator string) map[string]string {
	var contextBreakRegex = regexp.MustCompile(fmt.Sprintf(`(\n<p[^\n>]*>([^\n])*)%s(([^\n]*)</p>)`, sectionBreakIndicator))

	var subMatches = contextBreakRegex.FindAllStringSubmatch(fileContent, -1)
	var originalToSuggested = make(map[string]string, len(subMatches))
	if len(subMatches) == 0 {
		return originalToSuggested
	}

	for _, groups := range subMatches {
		var replaceValue = contextBreakEl
		if strings.TrimSpace(groups[2]) != "" || strings.TrimSpace(groups[4]) != "" {
			replaceValue = groups[1] + contextBreakEl + groups[3]
		}

		originalToSuggested[groups[0]] = replaceValue
	}

	return originalToSuggested
}
