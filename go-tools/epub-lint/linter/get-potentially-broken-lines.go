package linter

import (
	"regexp"
	"strings"
)

var unendedParagraphRegex = regexp.MustCompile(`((^|\n)[ \t]*<p[^>]*>)([^\n]*[a-zA-z,\d]["']?)( ?)(</p>\n)`)

func GetPotentiallyBrokenLines(fileContent string) map[string]string {
	var subMatches = unendedParagraphRegex.FindAllStringSubmatch(fileContent, -1)
	var originalToSuggested = make(map[string]string, len(subMatches))
	if len(subMatches) == 0 {
		return originalToSuggested
	}

	var parsedLines = map[string]struct{}{}
	for _, groups := range subMatches {
		var currentLine = groups[0]
		if _, ok := parsedLines[groups[0]]; ok {
			continue
		}

		var originalString = groups[0]
		var suggestedString = groups[1] + groups[3] + " "
		var nextLine = currentLine
		for lineIsPotentiallyBroken := true; lineIsPotentiallyBroken; {
			nextLine = getNextLine(fileContent, nextLine)
			originalString += nextLine

			var nextLineGroups = unendedParagraphRegex.FindStringSubmatch(nextLine)
			lineIsPotentiallyBroken = len(nextLineGroups) > 0
			if lineIsPotentiallyBroken {
				suggestedString += nextLineGroups[3] + " "
			} else {
				var endOfOpeningTag = strings.Index(nextLine, ">")

				if endOfOpeningTag == -1 {
					suggestedString += nextLine
				} else {
					suggestedString += nextLine[endOfOpeningTag+1:]
				}
			}
		}

		// we included an ending newline character for the next lines that we pulled bock
		// we do not need them when it comes to the ending of the original and suggested strings
		originalString = strings.TrimRight(originalString, "\n")
		suggestedString = strings.TrimRight(suggestedString, "\n")

		originalToSuggested[originalString] = suggestedString
	}

	return originalToSuggested
}

func getNextLine(fileContent, currentLine string) string {
	var endOfLineIndex = strings.Index(fileContent, currentLine)
	if endOfLineIndex == -1 {
		return ""
	}

	endOfLineIndex += len(currentLine)

	var substring = fileContent[endOfLineIndex:]
	var indexOfEndOfLine = strings.Index(substring, "\n")

	if indexOfEndOfLine == -1 {
		return substring
	}

	return substring[0 : indexOfEndOfLine+1]
}
