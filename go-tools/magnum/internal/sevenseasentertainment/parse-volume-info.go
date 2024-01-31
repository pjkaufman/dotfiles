package sevenseasentertainment

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var volumeNameRegex = regexp.MustCompile(`<a[^>]*>([^<]+)</a>`)
var earlyDigitalAccessRegex = regexp.MustCompile(`<b>Early Digital:</b> (\d{4}/\d{2}/\d{2})`)
var releaseDateRegex = regexp.MustCompile(`<b>Release Date</b>: (\d{4}/\d{2}/\d{2})`)

func ParseVolumeInfo(series, contentHtml string, volume int) (*VolumeInfo, error) {
	// get name from the anchor in the h3
	var firstHeading = volumeNameRegex.FindStringSubmatch(contentHtml)
	if len(firstHeading) < 2 {
		return nil, fmt.Errorf(`failed to get the name of volume %d for series "%s"`, volume, series)
	}

	var heading = firstHeading[1]
	if strings.Contains(strings.ToLower(heading), "(audiobook)") {
		return nil, nil
	}

	// get early digital release if present
	var earlyDigitalAccessDateInfo = earlyDigitalAccessRegex.FindStringSubmatch(contentHtml)
	var releaseDateString string
	if len(earlyDigitalAccessDateInfo) > 1 {
		releaseDateString = earlyDigitalAccessDateInfo[1]
	}

	// if not present get release date
	if releaseDateString == "" {
		var releaseDateInfo = releaseDateRegex.FindStringSubmatch(contentHtml)
		if len(releaseDateInfo) > 1 {
			releaseDateString = releaseDateInfo[1]
		}
	}

	var releaseDate *time.Time
	if releaseDateString != "" {
		tempDate, err := time.Parse(releaseDateFormat, releaseDateString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse \"%s\" to a date time value: %v", releaseDateString, err)
		}

		releaseDate = &tempDate
	}

	return &VolumeInfo{
		Name:        heading,
		ReleaseDate: releaseDate,
	}, nil
}
