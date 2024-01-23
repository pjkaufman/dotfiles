package cmd

import (
	"errors"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/config"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	NameArgEmpty = "name must have a non-whitespace value"
)

var (
	seriesName      string
	seriesType      string
	seriesPublisher string
	slugOverride    string
	seriesStatus    string
)

// AddCmd represents the add book info command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds the provided series info to the list of series to keep track of",
	// Example: heredoc.Doc(`To write the output of converting the cover file to a specific file:
	// song-converter create-cover -f cover-file.md -o output-file.html

	// To write the output of converting the cover file to std out:
	// song-converter create-cover -f cover-file.md
	// `),
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateAddSeriesFlags(seriesName)
		if err != nil {
			logger.WriteError(err.Error())
		}

		seriesInfo := config.GetConfig()
		if seriesInfo.HasSeries(seriesName) {
			logger.WriteInfo("The series already exists in the list.")

			return
		}

		var publisher = config.PublisherType(seriesPublisher)
		if strings.TrimSpace(seriesPublisher) == "" || !config.IsPublisherType(seriesPublisher) {
			publisher = selectPublisher()
		}

		var typeOfSeries = config.SeriesType(seriesType)
		if strings.TrimSpace(seriesType) == "" || !config.IsSeriesType(seriesType) {
			typeOfSeries = selectSeriesType()
		}

		var status = config.BookStatus(seriesStatus)
		if strings.TrimSpace(seriesStatus) == "" || !config.IsStatus(seriesStatus) {
			status = selectBookStatus()
		}

		var override *string
		if strings.TrimSpace(slugOverride) != "" {
			override = &slugOverride
		}

		newSeries := config.SeriesInfo{
			Name:         seriesName,
			Publisher:    publisher,
			Type:         typeOfSeries,
			SlugOverride: override,
			Status:       status,
		}

		seriesInfo.Series = append(seriesInfo.Series, newSeries)

		config.WriteConfig(seriesInfo)
	},
}

func init() {
	rootCmd.AddCommand(AddCmd)

	AddCmd.Flags().StringVarP(&seriesName, "name", "n", "", "the name of the series")
	AddCmd.Flags().StringVarP(&seriesPublisher, "publisher", "p", "", "the publisher of the series")
	AddCmd.Flags().StringVarP(&seriesType, "type", "t", "", "the series type")
	AddCmd.Flags().StringVarP(&slugOverride, "slug", "s", "", "the slug for the series to use instead of the one based on the series name")
	AddCmd.Flags().StringVarP(&slugOverride, "status", "r", string(config.Ongoing), "the status of the series (defaults to Ongoing)")

	AddCmd.MarkFlagRequired("name")
}

func ValidateAddSeriesFlags(seriesName string) error {
	if strings.TrimSpace(seriesName) == "" {
		return errors.New(NameArgEmpty)
	}

	return nil
}
