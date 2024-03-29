package config

import "strings"

type ReleaseInfo struct {
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
}

type SeriesInfo struct {
	Name              string        `json:"name"`
	TotalVolumes      int           `json:"total_volumes"`
	LatestVolume      string        `json:"latest_volume"`
	UnreleasedVolumes []ReleaseInfo `json:"unreleased_volumes"`
	SlugOverride      *string       `json:"slug_override"`
	Type              SeriesType    `json:"type"`
	Publisher         PublisherType `json:"publisher"`
	Status            SeriesStatus  `json:"status"`
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
