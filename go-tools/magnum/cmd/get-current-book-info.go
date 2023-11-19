package cmd

import (
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/config"
	jnovelclub "github.com/pjkaufman/dotfiles/go-tools/magnum/internal/jnovel-club"
	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/yenpress"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var verbose bool

// getCurrentInfoCmd represents the createCover command
var getCurrentInfoCmd = &cobra.Command{
	Use:   "get-info",
	Short: "Takes in the cover file path and creates the html cover file",
	Example: heredoc.Doc(`To write the output of converting the cover file to a specific file:
	song-converter create-cover -f cover-file.md -o output-file.html
	
	To write the output of converting the cover file to std out:
	song-converter create-cover -f cover-file.md
	`),
	Run: func(cmd *cobra.Command, args []string) {
		seriesInfo := config.GetConfig()

		for i, series := range seriesInfo.Series {
			seriesInfo.Series[i] = getSeriesVolumeInfo(series)
		}

		config.WriteConfig(seriesInfo)
	},
}

func init() {
	rootCmd.AddCommand(getCurrentInfoCmd)

	getCurrentInfoCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show more info about what is going on")
}

func getSeriesVolumeInfo(seriesInfo config.SeriesInfo) config.SeriesInfo {
	logger.WriteInfo(fmt.Sprintf("Checking for volume info for \"%s\"", seriesInfo.Name))

	switch seriesInfo.Publisher {
	case config.YenPress:
		return yenPressGetSeriesVolumeInfo(seriesInfo)
	case config.JNovelClub:
		return jNovelClubGetSeriesVolumeInfo(seriesInfo)
	default:
		return seriesInfo
	}
}

func yenPressGetSeriesVolumeInfo(seriesInfo config.SeriesInfo) config.SeriesInfo {
	volumes, numVolumes := yenpress.GetVolumes(seriesInfo.Name, seriesInfo.SlugOverride, verbose)

	if len(volumes) == 0 {
		logger.WriteInfo("The yen press light novels do not exist for this series.")

		return seriesInfo
	}

	if numVolumes == seriesInfo.TotalVolumes {
		logger.WriteWarn("No change in list of volumes from last check.")

		for _, unreleasedVol := range seriesInfo.UnreleasedVolumes {
			logger.WriteInfo(fmt.Sprintf("\"%s\" releases on %s", unreleasedVol.Name, unreleasedVol.ReleaseDate))
		}

		return seriesInfo
	}

	var today = time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	var unreleasedVolumes = []string{}
	var releaseDateInfo = []string{}
	for _, info := range volumes {
		releaseDate := yenpress.GetReleaseDateInfo(info, verbose)

		if releaseDate != nil {
			if releaseDate.Before(today) {
				break
			} else {
				releaseDateInfo = append(releaseDateInfo, releaseDate.Format("January 2, 2006"))
				unreleasedVolumes = append(unreleasedVolumes, info.Name)
			}
		}
	}

	return printReleaseInfoAndUpdateSeriesInfo(seriesInfo, unreleasedVolumes, releaseDateInfo, numVolumes, volumes[0].Name)
}

func jNovelClubGetSeriesVolumeInfo(seriesInfo config.SeriesInfo) config.SeriesInfo {
	volumeInfo := jnovelclub.GetVolumeInfo(seriesInfo.Name, seriesInfo.SlugOverride, verbose)

	if len(volumeInfo) == 0 {
		logger.WriteInfo("The jnovel club light novels do not exist for this series.")

		return seriesInfo
	}

	if len(volumeInfo) == seriesInfo.TotalVolumes {
		logger.WriteWarn("No change in list of volumes from last check.")

		for _, unreleasedVol := range seriesInfo.UnreleasedVolumes {
			logger.WriteInfo(fmt.Sprintf("\"%s\" releases on %s", unreleasedVol.Name, unreleasedVol.ReleaseDate))
		}

		return seriesInfo
	}

	var today = time.Now()
	today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	var unreleasedVolumes = []string{}
	var releaseDateInfo = []string{}
	for _, info := range volumeInfo {
		if info.ReleaseDate.Before(today) {
			break
		} else {
			releaseDateInfo = append(releaseDateInfo, info.ReleaseDate.Format("January 2, 2006"))
			unreleasedVolumes = append(unreleasedVolumes, info.Name)
		}

	}

	return printReleaseInfoAndUpdateSeriesInfo(seriesInfo, unreleasedVolumes, releaseDateInfo, len(volumeInfo), volumeInfo[0].Name)
}

func printReleaseInfoAndUpdateSeriesInfo(seriesInfo config.SeriesInfo, unreleasedVolumes, releaseDateInfo []string, totalVolumes int, latestVolumeName string) config.SeriesInfo {
	var releaseInfo = []config.ReleaseInfo{}
	for i, unreleasedVol := range unreleasedVolumes {
		releaseInfo = append(releaseInfo, config.ReleaseInfo{
			Name:        unreleasedVol,
			ReleaseDate: releaseDateInfo[i],
		})

		logger.WriteInfo(fmt.Sprintf("\"%s\" releases on %s", unreleasedVol, releaseDateInfo[i]))
	}

	seriesInfo.TotalVolumes = totalVolumes
	seriesInfo.LatestVolume = latestVolumeName
	seriesInfo.UnreleasedVolumes = releaseInfo

	return seriesInfo
}
