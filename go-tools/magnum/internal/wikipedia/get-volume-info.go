package wikipedia

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/crawler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

type VolumeInfo struct {
	Name        string
	ReleaseDate *time.Time
}

var wikiTableRegex = regexp.MustCompile(`<table[^>]*class="wikitable"[^>]*>`)
var volumeRowHeaderRegex = regexp.MustCompile(`<th[^>]*scope="row"[^>]*id=[^>]*>([^<]*)</th>`)

const (
	wikiTableEnd    = `</table>`
	wikiTableRowEnd = `</tr>`
)

func GetVolumeInfo(userAgent, title string, slugOverride *string, verbose bool) []VolumeInfo {
	var seriesSlug string
	if slugOverride != nil {
		seriesSlug = *slugOverride
	} else {
		seriesSlug = convertTitleToSlug(title)
	}

	sections := getSectionInfo(userAgent, seriesSlug)
	var lnSection SectionInfo
	var sectionAfterLn SectionInfo
	var subSectionTiles []string
	for _, section := range sections.Parse.Sections {
		if lnSection.Anchor != "" {
			if section.Level <= lnSection.Level {
				sectionAfterLn = section
				break
			} else {
				var heading = section.Heading
				var htmlElEndIndicatorIndex = strings.Index(heading, ">")
				if htmlElEndIndicatorIndex != -1 {
					heading = heading[htmlElEndIndicatorIndex+1:]
					heading = heading[:strings.Index(heading, "<")]
				}

				subSectionTiles = append(subSectionTiles, heading)
			}

			continue
		}

		if strings.HasPrefix(strings.ToLower(section.Heading), "light novel") {

			lnSection = section
		}
	}

	if lnSection.Heading == "" {
		logger.WriteError("failed to get light novel section")
	}

	c := crawler.CreateNewCollyCrawler(verbose)

	var err error

	var contentHtml string
	c.OnHTML("#content > div.vector-page-toolbar", func(e *colly.HTMLElement) {
		var content = e.DOM.Parent()
		contentHtml, err = content.Html()

		if err != nil {
			logger.WriteError(fmt.Sprintf("failed to get content body: %s", err))
		}
	})

	var lnHeadingHtml string
	var startIndexOfLnSection int
	c.OnHTML("#"+lnSection.Anchor, func(e *colly.HTMLElement) {
		var parents = e.DOM.Parent()
		lnHeadingHtml, err = parents.Html()
		if err != nil {
			logger.WriteError(fmt.Sprintf("failed to get content body: %s", err))
		}

		startIndexOfLnSection = strings.Index(contentHtml, lnHeadingHtml)
		if startIndexOfLnSection == -1 {
			logger.WriteError(fmt.Sprintf("failed to find light novel section: %s", err))
		}
	})

	var lnSectionAfterLnHtml string
	var endIndexOfLnSection int = -1
	if sectionAfterLn.Heading != "" {
		c.OnHTML("#"+sectionAfterLn.Anchor, func(e *colly.HTMLElement) {
			var parents = e.DOM.Parent()
			lnSectionAfterLnHtml, err = parents.Html()
			if err != nil {
				logger.WriteError(fmt.Sprintf("failed to get section after light novel section: %s", err))
			}

			endIndexOfLnSection = strings.Index(contentHtml, lnSectionAfterLnHtml)
			if endIndexOfLnSection == -1 {
				logger.WriteError(fmt.Sprintf("failed to find section after light novel section: %s", err))
			}
		})
	}

	var url = baseURL + wikiPath + seriesSlug
	err = c.Visit(url)
	if err != nil {
		logger.WriteError(fmt.Sprintf("failed call to wikipedia for \"%s\": %s", url, err))
	}

	var lnSectionHtml string
	if endIndexOfLnSection != -1 {
		lnSectionHtml = contentHtml[startIndexOfLnSection:endIndexOfLnSection]
	} else {
		lnSectionHtml = contentHtml[startIndexOfLnSection:]
	}

	if len(subSectionTiles) == 0 {
		subSectionTiles = []string{title}
	}

	var numTables = strings.Count(lnSectionHtml, "wikitable")
	if numTables == 0 {
		logger.WriteError(fmt.Sprintf("could not find tables for light novel section: %s", err))
	} else if len(subSectionTiles)+1 == numTables {
		subSectionTiles = append([]string{title}, subSectionTiles...)
	} else if len(subSectionTiles) != numTables {
		logger.WriteError(fmt.Sprintf("number of tables does not match number of table title prefixes for \"%s\": %d vs. %d", title, len(subSectionTiles), numTables))
	}

	var volumeInfo = []VolumeInfo{}
	for _, subSectionTitle := range subSectionTiles {
		start, stop := getTableStartAndStop(lnSectionHtml)
		volumeInfo = append(volumeInfo, parseWikipediaTableToVolumeInfo(subSectionTitle, lnSectionHtml[start:stop])...)

		lnSectionHtml = lnSectionHtml[stop:]
	}

	slices.Reverse(volumeInfo)

	return volumeInfo
}

func parseWikipediaTableToVolumeInfo(namePrefix, tableHtml string) []VolumeInfo {
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
			Name:        fmt.Sprintf("%s Vol. %s", namePrefix, rowSubmatches[1]),
			ReleaseDate: date,
		})

		rowHtml = rowHtml[endOfRow:]
	}

	return volumeInfo
}

func getEnglishReleaseDateFromRow(rowHtml string) (string, bool) {
	var actualColumns = strings.Count(rowHtml, `<td`)
	if actualColumns != 4 {
		return "", false
	}

	var releaseDateColumn = rowHtml
	for i := 0; i < 3; i++ {
		releaseDateColumn = releaseDateColumn[strings.Index(releaseDateColumn, `<td`)+4:]
	}

	var endOfRow = strings.Index(releaseDateColumn, `</td`)
	if endOfRow != -1 {
		releaseDateColumn = releaseDateColumn[:endOfRow]
	}

	var digitalVersionIndex = strings.Index(strings.ToLower(releaseDateColumn), "(digital")
	if digitalVersionIndex != -1 {
		releaseDateColumn = releaseDateColumn[:digitalVersionIndex]
	} else {
		var firstOpeningHtmlIndicator = strings.Index(releaseDateColumn, "<")
		if firstOpeningHtmlIndicator != -1 {
			releaseDateColumn = releaseDateColumn[:firstOpeningHtmlIndicator]
		}
	}

	releaseDateColumn = strings.TrimSpace(releaseDateColumn)
	if releaseDateColumn == "—" || strings.Contains(strings.ToLower(releaseDateColumn), "(physical") {
		return "", true
	}

	return releaseDateColumn, true
}

func getTableStartAndStop(sectionHtml string) (int, int) {
	var wikiStartLocation = wikiTableRegex.FindStringIndex(sectionHtml)
	if len(wikiStartLocation) == 0 {
		return -1, -1
	}

	var wikipediaTableStart = wikiStartLocation[0]
	var wikipediaTableEnd = strings.Index(sectionHtml, wikiTableEnd)

	if wikipediaTableEnd == -1 {
		return wikipediaTableStart, len(sectionHtml)
	}

	return wikipediaTableStart, wikipediaTableEnd + len(wikiTableEnd)
}

func convertTitleToSlug(title string) string {
	return strings.ReplaceAll(title, " ", "_")
}
