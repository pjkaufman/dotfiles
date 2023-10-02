package linter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

var validPageListIdsRegex = regexp.MustCompile(fmt.Sprintf(`id[ \t]*=["']((%s)\d+)["']`, strings.Join(validPageListAbbreviations, "|")))

func GetPageIdsForFile(l logger.Logger, text, file string, pageIds []PageIdInfo) []PageIdInfo {
	var validPageIds = validPageListIdsRegex.FindAllStringIndex(text, -1)
	if len(validPageIds) == 0 {
		return pageIds
	}

	for _, locs := range validPageIds {
		pageIds = append(pageIds, parseIdToIdInfo(l, text[locs[0]:locs[1]], file))
	}

	return pageIds
}

func parseIdToIdInfo(l logger.Logger, id, file string) PageIdInfo {
	var startOfId = strings.Index(id, "=")
	var val = id[startOfId+2 : len(id)-1]
	var num = validPageListAbbrevsRegex.ReplaceAllString(val, "")
	intVar, err := strconv.Atoi(num)
	if err != nil {
		l.WriteWarn(fmt.Sprintf(`Possible bad id "%s" tried to parse "%s": %v`, id, num, err))
	}

	return PageIdInfo{
		Id:     val,
		Number: intVar,
		File:   file,
	}
}
