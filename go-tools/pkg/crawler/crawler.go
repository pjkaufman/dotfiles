package crawler

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

func CreateNewCrawler(verbose bool) *colly.Collector {
	c := colly.NewCollector()

	if verbose {
		c.OnRequest(func(r *colly.Request) {
			logger.WriteInfo(fmt.Sprintf("Visiting: %v", r.URL))
		})

		c.OnResponse(func(r *colly.Response) {
			logger.WriteInfo(fmt.Sprintf("Page visited: %v", r.Request.URL))
		})

		c.OnScraped(func(r *colly.Response) {
			logger.WriteInfo(fmt.Sprintf("Finished visiting: %v", r.Request.URL))
		})
	}

	c.OnError(func(_ *colly.Response, err error) {
		logger.WriteError(fmt.Sprintf("Something went wrong making an http call: %s", err))
	})

	return c
}
