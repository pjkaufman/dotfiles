package linter

import "strings"

func ExtraStringReplace(text string, extraFindAndReplaces map[string]string, numHits map[string]int) string {
	var newText = text
	for search, replace := range extraFindAndReplaces {
		if hits, ok := numHits[search]; ok {
			numHits[search] = hits + strings.Count(newText, search)
		} else {
			numHits[search] = strings.Count(newText, search)
		}

		newText = strings.ReplaceAll(newText, search, replace)
	}

	return newText
}
