package config

import "strings"

type SeriesInfo struct {
	Name              string        `json:"name"`
	TotalVolumes      int           `json:"total_volumes"`
	LatestVolume      string        `json:"latest_volume"`
	UnreleasedVolumes []string      `json:"unreleased_volumes"`
	Type              SeriesType    `json:"type"`
	Publisher         PublisherType `json:"publisher"`
}

type Config struct {
	Series []SeriesInfo `json:"series"`
}

func (c *Config) HasSeries(name string) bool {
	for _, series := range c.Series {
		if strings.EqualFold(name, series.Name) {
			return true
		}
	}

	return false
}
