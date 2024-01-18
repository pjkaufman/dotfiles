package cmd

import (
	"fmt"

	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/config"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

// ListCmd represents the add book info command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the names of each of the series that is currently being tracked",
	// Example: heredoc.Doc(`To write the output of converting the cover file to a specific file:
	// song-converter create-cover -f cover-file.md -o output-file.html

	// To write the output of converting the cover file to std out:
	// song-converter create-cover -f cover-file.md
	// `),
	Run: func(cmd *cobra.Command, args []string) {
		seriesInfo := config.GetConfig()

		if len(seriesInfo.Series) == 0 {
			logger.WriteInfo("No series have been added to the list to keep track of.")

			return
		}

		for _, series := range seriesInfo.Series {
			logger.WriteInfo(series.Name)
			if verbose {
				logger.WriteInfo("Publisher: " + string(series.Publisher))
				logger.WriteInfo("Type: " + string(series.Type))
				logger.WriteInfo(fmt.Sprintf("Total Volumes: %d", series.TotalVolumes))

				var slugOverride = "N/A"
				if series.SlugOverride != nil {
					slugOverride = *series.SlugOverride
				}
				logger.WriteInfo("Slug Override: " + slugOverride)

				logger.WriteInfo("")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ListCmd)

	ListCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show more the publisher and other info about the series")
}
