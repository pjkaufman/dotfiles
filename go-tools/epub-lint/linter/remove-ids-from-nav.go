package linter

import (
	"fmt"
	"regexp"
	"strings"
)

var listAnchorHrefAttributeRegex = regexp.MustCompile("(href=[\"'])([^\"'#]*)(#[^\"']+)?([\"'])")
var openingNavRegex = regexp.MustCompile("<nav[^>\n]*>")
var openingListItemRegex = regexp.MustCompile("<li[^>\n]*>")
var ErrNoEndOfNav = fmt.Errorf("nav is incorrectly formatted since it has no closing nav element")

const (
	closingNavEl = "</nav>"
)

// RemoveIdsFromNav will remove ids from the nav element in an xhtml file if it is present
// It is meant to clean up the TOC for epub files that use an xhtml file as the TOC
// Note: it assumes the TOC will show up as list items in the nav element
func RemoveIdsFromNav(text string) (string, error) {
	var startOfNav = openingNavRegex.FindStringIndex(text)
	if len(startOfNav) == 0 {
		return text, nil
	}

	var endOfNav = strings.Index(text, closingNavEl)
	if endOfNav < 0 {
		return text, ErrNoEndOfNav
	}

	var existingFileLinks = make(map[string]int)
	var originalNav = text[startOfNav[0]:endOfNav]
	var newNav = ""
	var openingNavPoints = openingListItemRegex.FindAllStringIndex(originalNav, -1)
	var startOfLastEl = len(originalNav)
	for i := len(openingNavPoints) - 1; i >= 0; i-- {
		var startingIndex = openingNavPoints[i][0]
		navPoint, err := removeIdFromListItemLink(originalNav[startingIndex:startOfLastEl], existingFileLinks)
		if err != nil {
			return text, err
		}

		startOfLastEl = startingIndex
		newNav = navPoint + newNav
	}
	if len(openingNavPoints) > 0 {
		newNav = originalNav[0:startOfLastEl] + newNav
	}

	return strings.Replace(text, originalNav, newNav, 1), nil
}

func removeIdFromListItemLink(listItem string, existingFiles map[string]int) (string, error) {
	var groups = listAnchorHrefAttributeRegex.FindStringSubmatch(listItem)
	if len(groups) != 5 {
		return listItem, fmt.Errorf(fmt.Sprintf(`possible problem with list anchor tag href: %v`, groups))
	}

	var file = groups[2]
	if _, exists := existingFiles[file]; exists {
		return listItem, fmt.Errorf(`there is more than one reference to "%s". Consider restructuring the epub to fix up the nav`, file)
	}
	existingFiles[file] = 0

	return strings.Replace(listItem, groups[0], groups[1]+file+groups[4], 1), nil
}
