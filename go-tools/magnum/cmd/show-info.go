package cmd

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/config"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

// ShowInfoCmd represents the add book info command
var ShowInfoCmd = &cobra.Command{
	Use:   "show-info",
	Short: "Shows each series that has upcoming releases along with when the releases are in the order they are going to be released",
	Example: heredoc.Doc(`To show upcoming releases in order of when they are releasing:
	magnum show-info
	`),
	Run: func(cmd *cobra.Command, args []string) {
		seriesInfo := config.GetConfig()

		if len(seriesInfo.Series) == 0 {
			logger.WriteInfo("No series have been added to the list to keep track of.")

			return
		}

		var unreleasedVolumes []config.ReleaseInfo
		for _, series := range seriesInfo.Series {
			if len(series.UnreleasedVolumes) == 0 {
				continue
			}

			for i, unreleasedVolume := range series.UnreleasedVolumes {
				if strings.HasPrefix(unreleasedVolume.Name, "Vol") {
					series.UnreleasedVolumes[i].Name = series.Name + ": " + unreleasedVolume.Name
				}
			}

			unreleasedVolumes = append(unreleasedVolumes, series.UnreleasedVolumes...)
		}

		if len(unreleasedVolumes) == 0 {
			logger.WriteInfo("No release are upcoming")
			return
		}

		logger.WriteInfo("Upcoming releases:")
		logger.WriteInfo("")
		sort.Slice(unreleasedVolumes, func(i, j int) bool {
			if unreleasedVolumes[i].ReleaseDate == defaultReleaseDate {
				return false
			}

			date1, err := time.Parse(releaseDateFormat, unreleasedVolumes[i].ReleaseDate)
			if err != nil {
				logger.WriteError(fmt.Sprintf("failed to parse release date \"%s\" for \"%s\": %s", unreleasedVolumes[i].Name, unreleasedVolumes[i].ReleaseDate, err))
			}

			date2, err := time.Parse(releaseDateFormat, unreleasedVolumes[j].ReleaseDate)
			if err != nil {
				logger.WriteError(fmt.Sprintf("failed to parse release date \"%s\" for \"%s\": %s", unreleasedVolumes[j].Name, unreleasedVolumes[j].ReleaseDate, err))
			}

			return date1.Before(date2)
		})

		for _, unreleasedVolume := range unreleasedVolumes {
			logger.WriteInfo(getUnreleasedVolumeDisplayText(unreleasedVolume.Name, unreleasedVolume.ReleaseDate))
		}
	},
}

func init() {
	rootCmd.AddCommand(ShowInfoCmd)
}
