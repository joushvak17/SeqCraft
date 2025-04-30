package cli

import (
	"github.com/joushvak17/SeqCraft/internal/parse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seqcraft",
	Short: "SeqCraft - Bioinformatics CLI Tool written in Go",
	Long:  "SeqCraft - Bioinformatics CLI tool written in Go that provides efficient features for FASTA file analysis, including sequence parsing, alignment, and structure prediction.",
}

// Execute runs the root command of the SeqCraft CLI tool.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(parse.NewParseCmd())
}
