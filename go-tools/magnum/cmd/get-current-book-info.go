package cmd

import (
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/config"
	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/yenpress"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

// const (
// 	CoverPathArgEmpty  = "cover-file must have a non-whitespace value"
// 	CoverPathNotMdFile = "cover-file must be an md file"
// )

var verbose bool

// var coverOutputFile string
// var coverInputFilePath string

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
		// var series = "The Asterisk War"
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
	// createCoverCmd.Flags().StringVarP(&coverOutputFile, "output", "o", "", "the html file to write the output to")
	// createCoverCmd.MarkFlagRequired("cover-file")
}

func getSeriesVolumeInfo(seriesInfo config.SeriesInfo) config.SeriesInfo {
	logger.WriteInfo(fmt.Sprintf("Checking for volume info for \"%s\"", seriesInfo.Name))
	volumes, numVolumes := yenpress.GetVolumes(seriesInfo.Name, verbose)

	if len(volumes) == 0 {
		logger.WriteInfo("The yen press light novels do not exist for this series.")

		return seriesInfo
	}

	if numVolumes == seriesInfo.TotalVolumes {
		logger.WriteWarn("No change in volumes from last check.")
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

	for i, unreleasedVol := range unreleasedVolumes {
		logger.WriteInfo(fmt.Sprintf("\"%s\" releases on %s", unreleasedVol, releaseDateInfo[i]))
	}

	// set the volume info we have gathered before finishing up in the function
	seriesInfo.TotalVolumes = numVolumes
	seriesInfo.LatestVolume = volumes[0].Name
	seriesInfo.UnreleasedVolumes = unreleasedVolumes

	return seriesInfo
}

// func ValidateCreateCoverFlags(songsCoverFilePath string) error {
// 	if strings.TrimSpace(songsCoverFilePath) == "" {
// 		return errors.New(CoverPathArgEmpty)
// 	}

// 	if !strings.HasSuffix(songsCoverFilePath, ".md") {
// 		return errors.New(CoverPathNotMdFile)
// 	}

// 	return nil
// }
