package yenpress

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/crawler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

type VolumeInfo struct {
	Name         string
	RelativeLink string
}

func GetVolumes(seriesName string, verbose bool) ([]*VolumeInfo, int) {
	c := crawler.CreateNewCrawler(verbose)

	var volumes = []*VolumeInfo{}

	c.OnHTML("#volumes-list > div > div > div > a", func(e *colly.HTMLElement) {
		var link = e.Attr("href")
		if strings.TrimSpace(link) != "" {
			volumes = append(volumes, &VolumeInfo{
				RelativeLink: link,
			})
		}
	})

	var index = 0
	c.OnHTML("#volumes-list > div > div > div > a > p > span", func(e *colly.HTMLElement) {
		if strings.TrimSpace(e.Text) != "" {
			volumes[index].Name = e.Text
			index++
		}
	})

	var numVolumes int
	c.OnHTML("body > div > div:nth-child(5) > div > section.content-heading.fade-in-container > div > h1 > sup", func(e *colly.HTMLElement) {
		if strings.TrimSpace(e.Text) != "" {
			val, err := strconv.Atoi(e.Text)
			if err == nil {
				numVolumes = val
			}
		}
	})

	var seriesSlug = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(seriesName), " ", "-"), ")", ""), "(", "")
	var seriesURL = baseURL + seriesPath + seriesSlug
	err := c.Visit(seriesURL)
	if err != nil {
		logger.WriteError(fmt.Sprintf("failed call to yen press: %s", err))
		return nil, 0
	}

	return volumes, numVolumes
}
