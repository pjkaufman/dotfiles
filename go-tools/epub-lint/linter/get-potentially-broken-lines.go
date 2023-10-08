package linter

import (
	"regexp"
	"strings"
)

var unendedParagraphRegex = regexp.MustCompile(`((^|\n)[ \t]*<p[^>]*>)([^\n]*[a-zA-z,\d]["']?)( ?)(</p>\n)`)
var paragraphsWithDoubleQuotes = regexp.MustCompile(`((^|\n)[ \t]*<p[^>]*>)([^\n]*)(")([^\n]*)(</p>\n)`)

func GetPotentiallyBrokenLines(fileContent string) map[string]string {
	var originalToSuggested = make(map[string]string)
	var parsedLines = map[string]struct{}{}

	parseUnendedParagraphs(fileContent, parsedLines, originalToSuggested)
	parseUnendedDoubleQuotes(fileContent, parsedLines, originalToSuggested)

	return originalToSuggested
}

func parseUnendedParagraphs(fileContent string, parsedLines map[string]struct{}, originalToSuggested map[string]string) {
	var subMatches = unendedParagraphRegex.FindAllStringSubmatch(fileContent, -1)
	if len(subMatches) == 0 {
		return
	}

	for _, groups := range subMatches {
		var currentLine = groups[0]
		if _, ok := parsedLines[groups[0]]; ok {
			continue
		}

		parsedLines[currentLine] = struct{}{}

		var originalString = groups[0]
		var suggestedString = groups[1] + groups[3] + " "
		var nextLine = currentLine
		for lineIsPotentiallyBroken := true; lineIsPotentiallyBroken; {
			nextLine = getNextLine(fileContent, nextLine)
			parsedLines[nextLine] = struct{}{}
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
}

func parseUnendedDoubleQuotes(fileContent string, parsedLines map[string]struct{}, originalToSuggested map[string]string) {
	var subMatches = paragraphsWithDoubleQuotes.FindAllStringSubmatch(fileContent, -1)
	if len(subMatches) == 0 {
		return
	}

	for _, groups := range subMatches {
		var currentLine = groups[0]
		var doubleQuoteCount = strings.Count(currentLine, "\"")
		if doubleQuoteCount%2 == 0 {
			continue
		}

		// May need to handle parsed lines to make it so that it does not conflict between the two options that get parsed
		// but for now this should work just fine
		if _, ok := parsedLines[groups[0]]; ok {
			continue
		}

		parsedLines[currentLine] = struct{}{}

		var originalString = groups[0]
		var suggestedString = groups[1] + groups[3] + groups[4] + groups[5]
		if !strings.HasSuffix(suggestedString, " ") {
			suggestedString += " "
		}

		var nextLine = currentLine
		for lineIsPotentiallyBroken := true; lineIsPotentiallyBroken; {
			nextLine = getNextLine(fileContent, nextLine)
			parsedLines[nextLine] = struct{}{}
			originalString += nextLine
			doubleQuoteCount += strings.Count(nextLine, "\"")

			lineIsPotentiallyBroken = doubleQuoteCount%2 != 0 && nextLine != ""
			var endOfOpeningTag = strings.Index(nextLine, ">")

			if endOfOpeningTag == -1 {
				suggestedString += nextLine
			} else {
				suggestedString += nextLine[endOfOpeningTag+1:]
			}
		}

		// we included an ending newline character for the next lines that we pulled bock
		// we do not need them when it comes to the ending of the original and suggested strings
		originalString = strings.TrimRight(originalString, "\n")
		suggestedString = strings.TrimRight(suggestedString, "\n")
		suggestedString = strings.ReplaceAll(suggestedString, "  ", " ")

		originalToSuggested[originalString] = suggestedString
	}
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
