package parse

import (
	"fmt"
	"github.com/joushvak17/Bioinformatics-CLI-Tool/pkg/parse"
	"github.com/spf13/cobra"
)

// NewParseCmd creates and returns the `parse` command.
func NewParseCmd() *cobra.Command {
	parseCmd := &cobra.Command{
		Use:   "parse <file>",
		Short: "Parse and display the contents of a FASTA file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			records, err := parse.Parse(args[0])
			if err != nil {
				fmt.Printf("Error parsing FASTA file: %v\n", err)
				return
			}
			for _, record := range records {
				fmt.Printf(">%s %s\n%s\n", record.ID, record.Description, record.Sequence)
			}
		},
	}

	return parseCmd
}