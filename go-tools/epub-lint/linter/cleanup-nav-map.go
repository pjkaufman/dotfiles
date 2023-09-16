package linter

import (
	"fmt"
	"regexp"
	"strings"
)

var contentSrcAttributeRegex = regexp.MustCompile("(src=[\"'])([^\"'#]*)(#[^\"']+)?([\"'])")
var openingNavMapRegex = regexp.MustCompile("<navMap[^>\n]*>")
var openingNavPointRegex = regexp.MustCompile("<navPoint[^>\n]*>")
var playPointRegex = regexp.MustCompile(`(playOrder=["'])([^'"]*)(["'])`)
var ErrNoNavMap = fmt.Errorf("no nav map found")
var ErrNoEndOfNavMap = fmt.Errorf("nav map is incorrectly formatted since it has no closing nav map element")

const (
	closingNavMapEl = "</navMap>"
)

func CleanupNavMap(text string) (string, error) {
	var startOfNavMap = openingNavMapRegex.FindStringIndex(text)
	if len(startOfNavMap) == 0 {
		return text, ErrNoNavMap
	}

	var endOfNavMap = strings.Index(text, closingNavMapEl)
	if endOfNavMap < 0 {
		return text, ErrNoEndOfNavMap
	}

	var existingFileLinks = make(map[string]int)
	var originalNavMap = text[startOfNavMap[0]:endOfNavMap]
	var newNavMap = ""
	var openingNavPoints = openingNavPointRegex.FindAllStringIndex(originalNavMap, -1)
	var startOfLastEl = len(originalNavMap)
	for i := len(openingNavPoints) - 1; i >= 0; i-- {
		var startingIndex = openingNavPoints[i][0]
		navPoint, err := removeIdFromNavPointLink(originalNavMap[startingIndex:startOfLastEl], existingFileLinks)
		if err != nil {
			return text, err
		}

		navPoint = updatePlayOrder(navPoint, i+1)

		startOfLastEl = startingIndex
		newNavMap = navPoint + newNavMap
	}
	if len(openingNavPoints) > 0 {
		newNavMap = originalNavMap[0:startOfLastEl] + newNavMap
	}

	return strings.Replace(text, originalNavMap, newNavMap, 1), nil
}

func removeIdFromNavPointLink(navPoint string, existingFiles map[string]int) (string, error) {
	var groups = contentSrcAttributeRegex.FindStringSubmatch(navPoint)
	if len(groups) != 5 {
		return navPoint, fmt.Errorf(fmt.Sprintf(`possible problem with content tag src attribute: %v`, groups))
	}

	var file = groups[2]
	if _, exists := existingFiles[file]; exists {
		return navPoint, fmt.Errorf(`there is more than one reference to "%s". Consider restructuring the epub to fix up the nav map`, file)
	}
	existingFiles[file] = 0

	return strings.Replace(navPoint, groups[0], groups[1]+file+groups[4], 1), nil
}

func updatePlayOrder(navPoint string, playOrder int) string {
	return playPointRegex.ReplaceAllString(navPoint, fmt.Sprintf("${1}%d${3}", playOrder))
}
