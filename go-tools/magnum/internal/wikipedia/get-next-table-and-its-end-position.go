package wikipedia

import "strings"

func GetNextTableAndItsEndPosition(sectionHtml string) (string, int) {
	var wikiStartLocation = wikiTableRegex.FindStringIndex(sectionHtml)
	if len(wikiStartLocation) == 0 {
		return "", -1
	}

	var wikipediaTableStart = wikiStartLocation[0]
	var tableHtml = sectionHtml[wikipediaTableStart:]
	var potentialTableHtml = tableHtml
	var wikipediaTableEnd = wikipediaTableStart
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

		potentialTableHtml = potentialTableHtml[possibleWikiTableEnd:]
	}

	return tableHtml, wikipediaTableEnd
}
