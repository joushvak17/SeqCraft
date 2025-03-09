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

// Execute runs the root command and returns any error encountered.
// It should be called to start the CLI application.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add the parse command to the root command
	rootCmd.AddCommand(parse.NewParseCmd())
}

// TODO: The following will list out the things that need to be done in the future

// 1. Improve Usability
// - Easier to read output formats (e.g. tabular output) 
// 	 [DONE: Added color to output for the parse command]
// - Error handling and reporting
// - Input validation and error messages
// - Add more detailed help messages for each command
// - Progress bars for long-running tasks
// - Support for interactive mode with prompts and menus

// 2. Logging
// - Log levels: Use different log levels, like INFO, WARNING, ERROR, DEBUG
// - Log Files: Write logs to a file for debugging and analysis

// 3. Optimization
// - Parallel Processing: Use concurrency to speed up processing
// - Memory Optimization: Reduce memory usage for large datasets

// 4. Advanced Features
// - Sequence Alignment: Pairwise sequence alignment and support for MSA
// - Sequence Translation: Translate DNA/RNA sequences to protein sequences
// - Motif Search: Search for specific motifs in sequences
// - Secondary Structure Prediction: Predict secondary structure of protein sequences
// - Phylogenetic Analysis: Construct phylogenetic trees from sequence data

// Add Support for Additional File Formats
// Add a Web Interface
