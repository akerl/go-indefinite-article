package cmd

import (
	"fmt"

	"github.com/akerl/go-indefinite-article/indefinite"

	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse WORD",
	Short: "parse the article for WORD",
	RunE:  parseRunner,
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

func parseRunner(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no word given. See --help for syntax")
	}
	word := args[0]
	result := indefinite.AddArticle(word)
	fmt.Printf("%s", result)
	return nil
}
