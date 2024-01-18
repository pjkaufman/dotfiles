package cmd

import (
	"fmt"
)

var (
	verbose bool
)

const (
	defaultReleaseDate = "TBA"
	releaseDateFormat  = "January 2, 2006"
	userAgent          = "Magnum/1.0"
)

func getUnreleasedVolumeDisplayText(unreleasedVol, releaseDate string) string {
	if releaseDate == defaultReleaseDate {
		return fmt.Sprintf("\"%s\" release has not been announced yet", unreleasedVol)
	}

	return fmt.Sprintf("\"%s\" releases on %s", unreleasedVol, releaseDate)
}
