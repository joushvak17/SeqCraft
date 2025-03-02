package cli

import (
	"github.com/joushvak17/Bioinformatics-CLI-Tool/internal/parse" // Import the parse package
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Bioinformatics-CLI-Tool",
	Short: "Bioinformatics CLI Tool written in Go",
	Long:  "CLI tool for sequence analysis, alignment, and structure analysis, all accessible through an easy-to-use command line interface.",
}

// Execute runs the root command and returns any error encountered.
// It should be called to start the CLI application.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add the parse command to the root command
	rootCmd.AddCommand(parse.NewParseCmd())
}