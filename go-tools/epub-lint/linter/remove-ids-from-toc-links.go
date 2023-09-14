package linter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
)

// TODO: handle lang="en" xml:lang="en" in html tag in non-opf files

var contentElRegex = regexp.MustCompile("(<content.* src=[\"'])([^\"']*)([\"'][^>\n]*>)")
var contentSrcAttributeRegex = regexp.MustCompile("(src=[\"'])([^\"'#]*)(#[^\"']+)?([\"'])")
var listAnchorElRegex = regexp.MustCompile("(<li>[ \t\n\r]*<a.* href=[\"'])([^\"']*)([\"'][^>\n]*>)")
var listAnchorHrefAttributeRegex = regexp.MustCompile("(href=[\"'])([^\"'#]*)(#[^\"']+)?([\"'])")
var ExistingFileLinks map[string]int

func RemoveIdsFromTocLinks(text, filePath string) (string, string) {
	var newText = text
	if filePath == "nav" {
		newText = RemoveIdsFromContentLinks(text)
	} else if filePath == "ncx" {
		newText = RemoveIdsFromListAnchorLinks(text)
	}

	return newText, ""
}

func RemoveIdsFromContentLinks(text string) string {
	ExistingFileLinks = make(map[string]int)

	return contentElRegex.ReplaceAllStringFunc(text, removeIdsFromSrcLinks)
}

func removeIdsFromSrcLinks(part string) string {
	var groups = contentSrcAttributeRegex.FindStringSubmatch(part)
	if len(groups) != 5 {
		utils.WriteWarn(fmt.Sprintf(`possible problem with content tag: \"%s\", %v`, part, groups))

		return part
	}

	var file = groups[2]
	if count, exists := ExistingFileLinks[file]; exists {
		ExistingFileLinks[file] = count + 1
	} else {
		ExistingFileLinks[file] = 0
	}

	return strings.Replace(part, groups[0], groups[1]+file+groups[4], 1)
}

func RemoveIdsFromListAnchorLinks(text string) string {
	ExistingFileLinks = make(map[string]int)

	return listAnchorElRegex.ReplaceAllStringFunc(text, removeIdsFromHrefLinks)
}

func removeIdsFromHrefLinks(part string) string {
	var groups = listAnchorHrefAttributeRegex.FindStringSubmatch(part)
	if len(groups) != 5 {
		utils.WriteWarn(fmt.Sprintf(`possible problem with list anchor tag tag: \"%s\", %v`, part, groups))

		return part
	}

	var file = groups[2]
	if count, exists := ExistingFileLinks[file]; exists {
		ExistingFileLinks[file] = count + 1
	} else {
		ExistingFileLinks[file] = 0
	}

	return strings.Replace(part, groups[0], groups[1]+file+groups[4], 1)
}
