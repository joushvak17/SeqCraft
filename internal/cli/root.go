package cli

import (
	"github.com/fatih/color"
	"github.com/joushvak17/SeqCraft/internal/parse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   color.GreenString("SeqCraft"),
	Short: color.GreenString("SeqCraft") + " - Bioinformatics CLI Tool written in Go",
	Long:  color.GreenString("SeqCraft") + " - CLI tool for sequence analysis, alignment, and structure analysis, all accessible through an easy-to-use command line interface.",
}

// Execute runs the root command of the SeqCraft CLI tool.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(parse.NewParseCmd())
}
