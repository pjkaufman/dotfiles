package sevenseasentertainment

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

var volumeNameRegex = regexp.MustCompile(`<a[^>]*>([^<]+)</a>`)
var earlyDigitalAccessRegex = regexp.MustCompile(`<b>Early Digital:</b> (\d{4}/\d{2}/\d{2})`)
var releaseDateRegex = regexp.MustCompile(`<b>Release Date</b>: (\d{4}/\d{2}/\d{2})`)
var seriesInvalidSlugCharacters = regexp.MustCompile(`[\(\),:\-?!]`)

func GetVolumeInfo(seriesName string, slugOverride *string, verbose bool) []VolumeInfo {
	var seriesSlug string
	if slugOverride != nil {
		seriesSlug = *slugOverride
	} else {
		seriesSlug = getSeriesSlugFromName(seriesName)
	}

	c := crawler.CreateNewCollyCrawler(verbose)

	var err error

	var volumeContent = []string{}
	c.OnHTML(".series-volume", func(e *colly.HTMLElement) {
		contentHtml, err := e.DOM.Html()
		if err != nil {
			logger.WriteError(fmt.Sprintf("failed to get content body: %s", err))
		}

		volumeContent = append(volumeContent, contentHtml)
	})

	var url = googleCacheURL + baseURL + seriesPath + seriesSlug + "/"
	err = c.Visit(url)
	if err != nil {
		logger.WriteError(fmt.Sprintf("failed call to google cache for \"%s\": %s", url, err))
	}

	var volumeInfo = make([]VolumeInfo, len(volumeContent))
	for i, contentHtml := range volumeContent {
		var tempVolumeInfo, err = ParseVolumeInfo(seriesName, contentHtml, i+1)
		if err != nil {
			logger.WriteError(err.Error())
		}

		if tempVolumeInfo != nil {
			volumeInfo[i] = *tempVolumeInfo
		}
	}

	slices.Reverse(volumeInfo)

	return volumeInfo
}

func getSeriesSlugFromName(seriesName string) string {
	var slug = seriesInvalidSlugCharacters.ReplaceAllString(seriesName, "")

	slug = strings.ReplaceAll(strings.ToLower(slug), " ", "-")
	slug = strings.ReplaceAll(strings.ToLower(slug), "'", "-")

	return slug
}
