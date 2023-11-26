package cmd

import (
	"errors"
	"strings"

	"github.com/MakeNowJust/heredoc"
	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/config"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	NameArgEmpty      = "name must have a non-whitespace value"
	TypeArgEmpty      = "type must have a non-whitespace value"
	PublisherArgEmpty = "publisher must have a non-whitespace value"
)

var (
	seriesName      string
	seriesType      string
	seriesPublisher string
	slugOverride    string
)

// AddCmd represents the add book info command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds the provided series info to the list of series to keep track of",
	Example: heredoc.Doc(`To write the output of converting the cover file to a specific file:
	song-converter create-cover -f cover-file.md -o output-file.html
	
	To write the output of converting the cover file to std out:
	song-converter create-cover -f cover-file.md
	`),
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateAddSeriesFlags(seriesName, seriesPublisher, seriesType)
		if err != nil {
			logger.WriteError(err.Error())
		}

		seriesInfo := config.GetConfig()
		if seriesInfo.HasSeries(seriesName) {
			logger.WriteInfo("The series already exists in the list.")

			return
		}

		var override *string
		if strings.TrimSpace(slugOverride) != "" {
			override = &slugOverride
		}

		newSeries := config.SeriesInfo{
			Name:         seriesName,
			Publisher:    config.PublisherType(seriesPublisher),
			Type:         config.SeriesType(seriesType),
			SlugOverride: override,
		}

		seriesInfo.Series = append(seriesInfo.Series, newSeries)

		config.WriteConfig(seriesInfo)
	},
}

func init() {
	rootCmd.AddCommand(AddCmd)

	AddCmd.Flags().StringVarP(&seriesName, "name", "n", "", "the name of the series")
	AddCmd.Flags().StringVarP(&seriesPublisher, "publisher", "p", "", "publisher")
	AddCmd.Flags().StringVarP(&seriesType, "type", "t", "", "the series type")
	AddCmd.Flags().StringVarP(&slugOverride, "slug", "s", "", "the slug for the series to use instead of the one based on the series name")

	AddCmd.MarkFlagRequired("name")
}

func ValidateAddSeriesFlags(seriesName, seriesPublisher, seriesType string) error {
	if strings.TrimSpace(seriesName) == "" {
		return errors.New(NameArgEmpty)
	}

	if strings.TrimSpace(seriesPublisher) == "" {
		return errors.New(TypeArgEmpty)
	}

	if strings.TrimSpace(seriesType) == "" {
		return errors.New(PublisherArgEmpty)
	}

	return nil
}
