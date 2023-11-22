package wikipedia

import (
	"fmt"
	"strings"
	"time"

	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

func ParseWikipediaTableToVolumeInfo(namePrefix, tableHtml string) []VolumeInfo {
	var rows = volumeRowHeaderRegex.FindAllStringSubmatch(tableHtml, -1)
	if len(rows) == 0 {
		logger.WriteError("failed to find table row info for: " + namePrefix)
	}

	var volumeInfo = []VolumeInfo{}
	var rowHtml = tableHtml
	var startOfRow, endOfRow int
	var releaseDateString string
	var hasValidAmountOfColumns bool
	for _, rowSubmatches := range rows {
		startOfRow = strings.Index(rowHtml, rowSubmatches[0])
		rowHtml = rowHtml[startOfRow:]
		endOfRow = strings.Index(rowHtml, wikiTableRowEnd)

		releaseDateString, hasValidAmountOfColumns = getEnglishReleaseDateFromRow(rowHtml[:endOfRow])
		if !hasValidAmountOfColumns {
			logger.WriteWarn(fmt.Sprintf("skipped rows for \"%s\" since it did not have the expected amount of rows", namePrefix))
			return volumeInfo
		}
		var date *time.Time
		if releaseDateString != "" {
			tempDate, err := time.Parse(releaseDateFormat, releaseDateString)
			if err != nil {
				logger.WriteError(fmt.Sprintf("failed to parse \"%s\" to a date time value: %v", releaseDateString, err))
			}

			date = &tempDate
		}

		volumeInfo = append(volumeInfo, VolumeInfo{
			Name:        fmt.Sprintf("%s Vol. %s", namePrefix, strings.TrimSpace(rowSubmatches[1])),
			ReleaseDate: date,
		})

		rowHtml = rowHtml[endOfRow:]
	}

	return volumeInfo
}

func getEnglishReleaseDateFromRow(rowHtml string) (string, bool) {
	var actualColumns = strings.Count(rowHtml, `<td`)
	expectedDateColumn, ok := columnAmountToExpectedColumn[actualColumns]
	if !ok {
		return "", false
	}

	var releaseDateColumn = rowHtml
	for i := 0; i < expectedDateColumn; i++ {
		releaseDateColumn = releaseDateColumn[strings.Index(releaseDateColumn, `<td`)+4:]
	}

	var endOfRow = strings.Index(releaseDateColumn, `</td`)
	if endOfRow != -1 {
		releaseDateColumn = releaseDateColumn[:endOfRow]
	}

	var digitalVersionIndex = strings.Index(strings.ToLower(releaseDateColumn), "(digital")
	if digitalVersionIndex != -1 {
		releaseDateColumn = releaseDateColumn[:digitalVersionIndex]
	}

	if strings.HasPrefix(releaseDateColumn, "<") {
		releaseDateColumn = releaseDateColumn[strings.Index(releaseDateColumn, ">")+1:]
	}

	var firstOpeningHtmlIndicator = strings.Index(releaseDateColumn, "<")
	if firstOpeningHtmlIndicator != -1 {
		releaseDateColumn = releaseDateColumn[:firstOpeningHtmlIndicator]
	}

	releaseDateColumn = strings.TrimSpace(releaseDateColumn)
	if releaseDateColumn == "—" || releaseDateColumn == "TBA" || strings.Contains(strings.ToLower(releaseDateColumn), "(physical") {
		return "", true
	}

	return releaseDateColumn, true
}
