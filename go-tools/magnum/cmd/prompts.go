package cmd

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/config"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

func selectBookName(series []config.SeriesInfo, includeCompleted bool) string {
	var seriesNames = []string{}
	for _, series := range series {
		if series.Status != config.Completed || includeCompleted {
			seriesNames = append(seriesNames, series.Name)
		}
	}

	prompt := promptui.Select{
		Label: "Select Book Name",
		Items: seriesNames,
		Searcher: func(input string, index int) bool {
			seriesName := seriesNames[index]
			name := strings.Replace(strings.ToLower(seriesName), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)

			return strings.Contains(name, input)
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		logger.WriteError(fmt.Sprintf("Book name prompt failed %v", err))
	}

	return result
}

func selectBookStatus() config.Status {
	var statuses = []config.Status{
		config.Ongoing,
		config.Hiatus,
		config.Completed,
	}
	var seriesStatuses = []string{
		string(config.Ongoing) + " - Ongoing",
		string(config.Hiatus) + " - Hiatus",
		string(config.Completed) + " - Completed",
	}

	prompt := promptui.Select{
		Label: "Select Book Status",
		Items: seriesStatuses,
	}

	i, _, err := prompt.Run()
	if err != nil {
		logger.WriteError(fmt.Sprintf("Book status prompt failed %v", err))
	}

	return statuses[i]
}
