package cmd

import (
	"math/rand"
	"os"
	"time"

	"github.com/pjkaufman/dotfiles/go-tools/cat-ascii/internal/ascii"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cat-acii",
	Short: "A cat ascii art generator that displays a random cat ascii art on each invocation",
	Run: func(cmd *cobra.Command, args []string) {
		fileContent, err := ascii.GetAllAsciiFileContent()
		if err != nil {
			logger.WriteError(err.Error())
		}

		generator := rand.New(rand.NewSource(time.Now().UnixNano()))
		n := int(generator.Int63()) % len(fileContent)

		logger.WriteInfo(fileContent[n])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
