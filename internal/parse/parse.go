package parse

import (
	"fmt"
	"github.com/joushvak17/Bioinformatics-CLI-Tool/pkg/parse"
	"github.com/spf13/cobra"
)

// NewParseCmd creates and returns the `parse` command.
func NewParseCmd() *cobra.Command {
	var (
		// Define flags for the command
		sequenceLength bool

		// TODO: Add additional flags for analyzing the sequences
		// Primarily the GC content and the reverse complement
		// gcContent      bool
		// reverseComp    bool
	)

	parseCmd := &cobra.Command{
		Use:   "parse <file>",
		Short: "Parse and analyze a FASTA file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]
			records, err := parse.Parse(filename)
			if err != nil {
				fmt.Printf("Error parsing FASTA file: %v\n", err)
				return
			}

			for _, record := range records {
				fmt.Printf(">%s %s\n%s\n", record.ID, record.Description, record.Sequence)

				if sequenceLength {
					length := len(record.Sequence)
					fmt.Printf("Sequence Length: %d\n", length)
				}

				// TODO: Add additional analysis for the sequences
				// if gcContent {
				// 	// Calculate GC content
				// 	gc := calculateGC(record.Sequence)
				// 	fmt.Printf("GC Content: %.2f%%\n", gc)
				// }
				// if reverseComp {
				// 	// Calculate reverse complement
				// 	reverse := reverseComplement(record.Sequence)
				// 	fmt.Printf("Reverse Complement: %s\n", reverse)
				// }

				fmt.Println() // Add an empty line between records
			}
		},
	}

	// Add flags to the command
	parseCmd.Flags().BoolVarP(&sequenceLength, "length", "l", false, "Calculate sequence length")

	// TODO: Add additional flags for analyzing the sequences
	// Primarily the GC content and the reverse complement
	// parseCmd.Flags().BoolVarP(&reverseComp, "reverse", "r", false, "Calculate reverse complement")
	// parseCmd.Flags().BoolVarP(&sequenceLength, "length", "l", false, "Calculate sequence length")

	return parseCmd
}
