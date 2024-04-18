package wikipedia

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

func GetNextTableAndItsEndPosition(sectionHtml string) (string, int) {
	var wikiStartLocation = wikiTableRegex.FindStringIndex(sectionHtml)
	if len(wikiStartLocation) == 0 {
		return "", -1
	}

	var maxAttempts = strings.Count(sectionHtml, tableEnd)
	var wikipediaTableStart = wikiStartLocation[0]
	var tableHtml = sectionHtml[wikipediaTableStart:]
	var potentialTableHtml = tableHtml
	var wikipediaTableEnd = wikipediaTableStart
	var attemptNum = 1
	for {
		var possibleWikiTableEnd = strings.Index(potentialTableHtml, tableEnd)

		if possibleWikiTableEnd == -1 {
			return sectionHtml[wikipediaTableStart:], len(sectionHtml)
		}

		wikipediaTableEnd += possibleWikiTableEnd + len(tableEnd)
		tableHtml = sectionHtml[wikipediaTableStart:wikipediaTableEnd]

		if strings.Count(tableHtml, tableEnd) == strings.Count(tableHtml, tableStart) {
			break
		}

		potentialTableHtml = potentialTableHtml[possibleWikiTableEnd+len(tableEnd):]
		attemptNum++

		if attemptNum > maxAttempts {
			logger.WriteError(fmt.Sprintf("something went wrong trying to parse out the table from %s, as there were only %d instances of table endings and we are trying to find the %d table ending", sectionHtml, maxAttempts, attemptNum))
		}
	}

	return tableHtml, wikipediaTableEnd
}
