package sevenseasentertainment

import (
	"fmt"
	"slices"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/slug"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/crawler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

type VolumeInfo struct {
	Name        string
	ReleaseDate *time.Time
}

func GetVolumeInfo(seriesName string, slugOverride *string, verbose bool) []VolumeInfo {
	var seriesSlug string
	if slugOverride != nil {
		seriesSlug = *slugOverride
	} else {
		seriesSlug = slug.GetSeriesSlugFromName(seriesName)
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

	var volumeInfo = []VolumeInfo{}
	var index = 1
	for _, contentHtml := range volumeContent {
		var tempVolumeInfo, err = ParseVolumeInfo(seriesName, contentHtml, index)
		if err != nil {
			logger.WriteError(err.Error())
		}

		if tempVolumeInfo != nil {
			volumeInfo = append(volumeInfo, *tempVolumeInfo)
			index++
		}
	}

	slices.Reverse(volumeInfo)

	return volumeInfo
}
