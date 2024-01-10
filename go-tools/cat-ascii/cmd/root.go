package cmd

import (
	"fmt"
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
		generator := rand.New(rand.NewSource(time.Now().UnixNano()))
		n := int(generator.Int63()) % len(ascii.CatAsciiFileNames)

		var filePath = fmt.Sprintf(`%s.txt`, ascii.CatAsciiFileNames[n])
		data, err := ascii.AsciiEmbeds.ReadFile(filePath)
		if err != nil {
			logger.WriteError(fmt.Sprintf(`failed to get cat ascii art for embedded path "%s": %s`, filePath, err))
		}

		logger.WriteInfo(string(data))
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
